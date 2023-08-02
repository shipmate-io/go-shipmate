package go_shipmate

import (
	"encoding/json"
	"testing"
)

func TestMessage(t *testing.T) {
	messagePayload, _ := json.Marshal(userCreatedMessage{FirstName: "John", LastName: "Doe"})

	message := Message{
		Type:    "user.created",
		Payload: messagePayload,
	}

	if message.Type != "user.created" {
		t.Errorf("the Message type should be user.created; got %s", message.Type)
	}

	var restoredMessagePayload userCreatedMessage

	err := json.Unmarshal(message.Payload, &restoredMessagePayload)

	if err != nil {
		t.Errorf("Unable to deserialize Message payload")
	}

	if restoredMessagePayload.FirstName != "John" {
		t.Errorf("expected the first name of the user to be John; got %s", restoredMessagePayload.FirstName)
	}

	if restoredMessagePayload.LastName != "Doe" {
		t.Errorf("expected the last name of the user to be Doe; got %s", restoredMessagePayload.LastName)
	}
}
