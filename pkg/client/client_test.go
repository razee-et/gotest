package messagebroker

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"

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
