# Interact with Shipmate from your Go code

## Installation

You can install the module by running the following command in your terminal:

```bash
go get github.com/shipmate-io/go-shipmate
```

## Usage

### Message queue

You can interact with your Shipmate message queues as follows.

Publish a message:

```go
import (
    "fmt"
    shipmate "github.com/shipmate-io/go-shipmate"
    "os"
)

messageQueueName := os.Getenv("SHIPMATE_MESSAGE_QUEUE_NAME")
messageQueue, err := shipmate.NewMessageQueue(messageQueueName)

if err != nil {
    fmt.Println("Unable to instantiate message queue")
    return
}

message := shipmate.Message{
    Type: "user.created",
    Payload: []byte(`{ "first_name": "John", "last_name": "Doe" }`),
}

messageQueue.PublishMessage(&message)

messageQueue.Close()
```

Handle a message:

```go
package main

import (
    "github.com/gin-gonic/gin"
    shipmate "github.com/shipmate-io/go-shipmate"
    "io"
    "net/http"
)

func main() {
    r := gin.Default()
    
    r.POST("/shipmate/handle-message", func(c *gin.Context) {
        requestPayload, _ := io.ReadAll(c.Request.Body)

        message, err := shipmate.ParseMessage(requestPayload)

        if err != nil {
            c.String(http.StatusUnprocessableEntity, "Unable to parse message")
            return
        }
        
        // TODO: use message

        c.String(http.StatusOK, "Message handled")
    })

    r.Run()
}
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.