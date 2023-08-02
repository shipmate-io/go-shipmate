package go_shipmate

import (
	"encoding/json"
	"testing"
)

func TestParseMessage(t *testing.T) {
	requestPayload := []byte(`{
		"message": {
    		"attributes": {
				"type": "user.created"
			},
			"data": "ew0gImZpcnN0X25hbWUiOiAiSm9obiIsDSAibGFzdF9uYW1lIjogIkRvZSINfQo=",
			"messageId": "123456"
		}
	}`)

	message, err := ParseMessage(requestPayload)

	if err != nil {
		t.Errorf("Unable to parse message")
	}

	if message.Type != "user.created" {
		t.Errorf("the message type should be user.created; got %s", message.Type)
	}

	var messagePayload userCreatedMessage
	json.Unmarshal(message.Payload, &messagePayload)

	if messagePayload.FirstName != "John" {
		t.Errorf("expected the first name of the user to be John; got %s", messagePayload.FirstName)
	}

	if messagePayload.LastName != "Doe" {
		t.Errorf("expected the last name of the user to be Doe; got %s", messagePayload.LastName)
	}

	if message.Id != "123456" {
		t.Errorf("expected the message id to be 123456; got %s", message.Id)
	}
}
