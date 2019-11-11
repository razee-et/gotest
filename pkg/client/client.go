package messagebroker

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	"github.com/go-stomp/stomp"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const connectionTimeOut = 60 * 1000
const readChannelCapacity = 50
const writeChannelCapacity = 50

const delimiter = "; "
const applicationType = "application/x-protobuf"
const bufType = "acquia-protobuf-name: "
const ProtobufNamespace = "acquia.messages."

// Client contains necessary values for interacting with the message broker.
type client struct {
	nc       *tls.Conn
	sc       *stomp.Conn
	uri      string
	username string
	password string
	subs     []subscription
}

type Client interface {
	Subscribe(name string) error
	Unsubscribe(name string)
	Receive(name string) ([]byte, error)
	Send(name string, message string) error
	Disconnect() error
}

// subscription contains the values for interacting with a message broker subscription.
type subscription struct {
	name string
	sub  *stomp.Subscription
}

// Client creates and returns a client for interacting with the message broker.
func Connect(uri string, username string, password string) (*client, error) {
	netConn, err := tls.Dial("tcp", uri, &tls.Config{})

	if err != nil {
		return nil, err
	}

	var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
		stomp.ConnOpt.Login(username, password),
		// Turn off timeouts on either reading or writing.
		stomp.ConnOpt.HeartBeat(0, 0),
		stomp.ConnOpt.ReadChannelCapacity(readChannelCapacity),
		stomp.ConnOpt.WriteChannelCapacity(writeChannelCapacity),
	}
	stompConn, err := stomp.Connect(netConn, options...)
	if err != nil {
		netConn.Close()
		return nil, err
	}

	subscriptions := make([]subscription, 1)
	return &client{nc: netConn, sc: stompConn, uri: uri, username: username, password: password, subs: subscriptions}, nil
}

// Disconnect closes all connections.
func (c *client) Disconnect() error {
	stompErr := c.sc.Disconnect()
	netErr := c.nc.Close()
	if stompErr != nil {
		netErr = errors.Wrapf(netErr, "stomp disconnect error: %v", stompErr.Error())
	}
	return netErr
}

// Reconnect establishes connections again.
func (c *client) Reconnect() error {
	c.unsubscribeAll()
	newClient, err := Connect(c.uri, c.username, c.password)
	c.nc = newClient.nc
	c.sc = newClient.sc
	return err
}

// getSubscription looks for the stomp subscription in the client's subscriptions slice.
func (c *client) getSubscription(name string) *stomp.Subscription {
	for _, data := range c.subs {
		if data.name == name {
			return data.sub
		}
	}
	return nil
}

// Subscribe adds a subscription to the client for the given topic.
func (c *client) Subscribe(name string) error {
	sub := c.getSubscription(name)
	if sub != nil {
		// If already subscribed, do nothing.
		return nil
	}
	sub, err := c.sc.Subscribe(
		name,
		stomp.AckAuto,
		stomp.SubscribeOpt.Header("activemq.retroactive", "true"),
		stomp.SubscribeOpt.Header("activemq.dispatchAsync", "true"))
	if err != nil {
		if err == stomp.ErrAlreadyClosed || err == stomp.ErrClosedUnexpectedly {
			err = c.Reconnect()
			if err == nil {
				sub, err = c.sc.Subscribe(name, stomp.AckClient)
			}
		}
		if err != nil {
			errors.Wrapf(err, "Cannot subscribe to %v", name)
			return err
		}
	}
	c.subs = append(c.subs, subscription{name: name, sub: sub})
	return nil
}

// Unsubscribe removes the subscription from the client for the given topic.
func (c *client) Unsubscribe(name string) {
	for i, data := range c.subs {
		if data.name == name {
			c.subs = append(c.subs[:i], c.subs[i+1:]...)
			return
		}
	}
}

func (c *client) unsubscribeAll() {
	c.subs = make([]subscription, 0)
}

// Receive looks for a message on the given topic.
func (c *client) Receive(name string) (proto.Message, error) {
	var result proto.Message

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in Receive: %v\n", r)
			debug.PrintStack()
		}
	}()

	subscription := c.getSubscription(name)
	if subscription == nil {
		err := c.Subscribe(name)
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to receive on %v because unable to subscribe", name)
		}
		subscription = c.getSubscription(name)
		if subscription == nil {
			return nil, errors.Wrapf(err, "Unable to receive on %v because unable to subscribe", name)
		}

	}
	// The channel will not block forever, but instead time out and disconnect if there is no message.
	msg := <-subscription.C
	if msg == nil {
		err := c.Reconnect()
		return nil, err
	}
	if err := c.sc.Ack(msg); err != nil {
		return nil, err
	}
	if msg.Header != nil {
		messageName, err := extractMessageName(msg.Header.Get("content-type"))
		if err != nil {
			return nil, err
		}
		contentType := proto.MessageType(messageName)
		if contentType == nil {
			return nil, fmt.Errorf("unrecognized message type (%s)", messageName)
		}
		acquiaMessage := reflect.New(contentType.Elem()).Interface().(proto.Message)
		err = proto.Unmarshal(msg.Body, acquiaMessage)
		result = acquiaMessage.(proto.Message)
	}
	return result, nil
}

func extractMessageName(content string) (string, error) {
	sections := strings.Split(content, delimiter)
	if len(sections) > 1 {
		index := len(bufType)
		messageName := sections[1][index:]
		if messageName != "" {
			if strings.Contains(messageName, ProtobufNamespace) {
				return messageName, nil
			}
		}
	}
	return "", fmt.Errorf("unexpected content-type (%s)", content)
}

func generateContentType(message proto.Message) (string, error) {
	messageName := proto.MessageName(message) // Get the fully-qualified name of "message"'s message type
	if !strings.Contains(messageName, ProtobufNamespace) {
		return "", fmt.Errorf("unrecognized message type (%s)", ProtobufNamespace)
	}
	contentType := fmt.Sprintf("%s%s%s%s", applicationType, delimiter, bufType, messageName)
	return contentType, nil
}

// Send adds a message to the given subscription topic on the message broker.
func (c *client) Send(name string, message proto.Message) error {
	now := time.Now().UTC().UnixNano() / int64(time.Millisecond)
	expiration := fmt.Sprintf("%d", now+connectionTimeOut)

	contentType, err := generateContentType(message)
	if err != nil {
		return err
	}
	bytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	err = c.sc.Send(
		name,
		contentType,
		bytes,
		stomp.SendOpt.Receipt,
		stomp.SendOpt.Header("persistent", "true"),
		stomp.SendOpt.Header("expires", expiration))
	if err != nil {
		if err == stomp.ErrAlreadyClosed || err == stomp.ErrClosedUnexpectedly {
			err = c.Reconnect()
			if err == nil {
				err = c.sc.Send(
					name,
					contentType,
					bytes,
					stomp.SendOpt.Receipt,
					stomp.SendOpt.Header("persistent", "true"),
					stomp.SendOpt.Header("expires", expiration))
			}
		}
	}
	return err
}
