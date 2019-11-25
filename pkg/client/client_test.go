// +build unit

package messagebroker

import (
	"crypto/tls"
	"fmt"
	"github.com/go-stomp/stomp"
	"github.com/go-stomp/stomp/frame"
	"github.com/go-stomp/stomp/testutil"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	. "gopkg.in/check.v1"

	"github.com/acquia/message-broker-sdk-go/pkg/acquia/messages/hosting"
)

func TestExtractMessageName(t *testing.T) {
	expectedMessageName := "acquia.messages.hosting.CreateApplication"
	validContentType := fmt.Sprintf("application/x-protobuf; acquia-protobuf-name: %s", expectedMessageName)

	messageName, err := extractMessageName(validContentType)
	require.Nil(t, err)
	require.Equal(t, expectedMessageName, messageName)

	invalidContentType := "application/x-protobuf; acquia-protobuf-name: nope.not.today"
	messageName, err = extractMessageName(invalidContentType)
	require.NotNil(t, err)
	require.Empty(t, messageName)
	require.Contains(t, err.Error(), "unexpected")
}

func TestGenerateContentType(t *testing.T) {
	message := &hosting.CreateApplication{}
	expectedContentType := fmt.Sprintf("application/x-protobuf; acquia-protobuf-name: %s", proto.MessageName(message))
	contentType, err := generateContentType(message)
	require.Nil(t, err)
	require.Equal(t, expectedContentType, contentType)

	var empty proto.Message
	contentType, err = generateContentType(empty)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "unrecognized")
	require.Empty(t, contentType)
}

func TestGenerateReadTopicName(t *testing.T) {
	client := generateTestClient()
	formattedName, _ := client.generateReadTopicName("topic")
	require.Equal(t, formattedName, "/queue/Consumer.consumerA.acquia.messages.topic")
}

func TestGenerateWriteTopicName(t *testing.T) {
	client := generateTestClient()
	formattedName, _ := client.generateWriteTopicName("topic")
	require.Equal(t, formattedName, "/topic/acquia.messages.topic")
}

func TestValidateTopicParameter(t *testing.T) {
	error := validateTopicParameter("name")
	require.NoError(t, error)

	error = validateTopicParameter("name ?")
	require.Equal(t, error.Error(), "topic name or consumer can not be blank or have a '?' character")

	error = validateTopicParameter("")
	require.Equal(t, error.Error(), "topic name or consumer can not be blank or have a '?' character")

	error = validateTopicParameter(" ")
	require.Equal(t, error.Error(), "topic name or consumer can not be blank or have a '?' character")
}

func generateTestClient() *client {
	netConn, _ := tls.Dial("tcp", "127.0.0.1:123", &tls.Config{})
	stompFakeConn1, stompFakeConn2 := testutil.NewFakeConn(&C{})
	subscriptions := make([]subscription, 1)

	go func() {
		reader := frame.NewReader(stompFakeConn2)
		writer := frame.NewWriter(stompFakeConn2)
		reader.Read()
		connectedFrame := frame.New("CONNECTED")
		writer.Write(connectedFrame)
	}()

	stompConn, _ := stomp.Connect(stompFakeConn1)

	return &client{nc: netConn, sc: stompConn, uri: "127.0.0.1", username: "username", password: "password", subs: subscriptions, consumer: "consumerA"}
}
