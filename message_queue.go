package go_shipmate

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"encoding/json"
)

type requestPayloadType struct {
	Message struct {
		Attributes map[string]string `json:"attributes"`
		Data       string            `json:"data"`
		MessageId  string            `json:"messageId"`
	} `json:"message"`
}

type MessageQueue struct {
	googleClient *pubsub.Client
	name         string
}

func NewMessageQueue(name string) (*MessageQueue, error) {
	ctx := context.Background()

	shipmateConfig := ShipmateConfig{}

	shipmateEnvironmentId, err := shipmateConfig.GetEnvironmentId()

	if err != nil {
		return nil, err
	}

	googleClient, err := pubsub.NewClient(ctx, shipmateEnvironmentId)

	if err != nil {
		return nil, err
	}

	messageQueue := &MessageQueue{
		googleClient: googleClient,
		name:         name,
	}

	return messageQueue, nil
}

func ParseMessage(requestPayload []byte) (*Message, error) {
	var parsedRequestPayload requestPayloadType

	err := json.Unmarshal(requestPayload, &parsedRequestPayload)

	if err != nil {
		return nil, err
	}

	payload, _ := base64.StdEncoding.DecodeString(parsedRequestPayload.Message.Data)

	message := Message{
		Type:    parsedRequestPayload.Message.Attributes["type"],
		Payload: payload,
		Id:      parsedRequestPayload.Message.MessageId,
	}

	return &message, nil
}

func (m *MessageQueue) PublishMessage(message *Message) {
	ctx := context.Background()

	m.googleClient.Topic(m.name).Publish(ctx, &pubsub.Message{
		Data: message.Payload,
		Attributes: map[string]string{
			"type": message.Type,
		},
	})
}

func (m *MessageQueue) Close() {
	defer m.googleClient.Close()
}
