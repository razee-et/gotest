// +build integration

package messagebroker

import (
	"testing"
	"time"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/require"

	"github.com/acquia/message-broker-sdk-go/pkg/acquia/messages/hosting"
	"github.com/acquia/message-broker-sdk-go/pkg/acquia/messages/shared"
)

func TestSendReceiveMessageFromClient(t *testing.T) {
	testTopic := "integrationtest"

	con, err := Connect("localhost:61613", "user", "password", "consumer001")
	if err != nil {
		t.Error("connection failure to broker", err)
	}

	err = con.Subscribe(testTopic)
	require.Nil(t, err)

	h := shared.Header{CorrelationUUID: uuid.New(), ResponseTopic: testTopic}
	validMessage := hosting.CreateApplication{Header: &h, ApplicationUUID: uuid.New(), ApplicationName: "integration-test"}

	err = con.Send(testTopic, &validMessage)
	require.Nil(t, err)
	time.Sleep(1 * time.Second)

	message, err := con.Receive(testTopic)
	require.Nil(t, err)
	require.NotNil(t, message)

	con.Unsubscribe(testTopic)
}
