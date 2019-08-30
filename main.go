package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar/pulsar-client-go/pulsar"
	"log"
)

// func main() {
// 	client, err := pulsar.NewClient(pulsar.ClientOptions{
// 		URL:                     "pulsar://localhost:6650",
// 		OperationTimeoutSeconds: 5,
// 		MessageListenerThreads:  runtime.NumCPU(),
// 	})

// 	if err != nil {
// 		log.Fatalf("Could not instantiate Pulsar client: %v", err)
// 	}

// 	SendMessage(client)

// }

// func SendMessage(client pulsar.Client) {
// 	ctx := context.Background()
// 	producer, err := client.CreateProducer(pulsar.ProducerOptions{
// 		Topic: "my-topic",
// 	})

// 	if err != nil {
// 		log.Fatalf("Could not instantiate Pulsar producer: %v", err)
// 	}

// 	defer producer.Close()

// 	msg := pulsar.ProducerMessage{
// 		Payload: []byte("Hello, Pulsar"),
// 	}

// 	if err := producer.Send(ctx, msg); err != nil {
// 		log.Fatalf("Producer could not send message: %v", err)
// 	}
// }

func main() {
	// Instantiate a Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})

	if err != nil {
		log.Fatal(err)
	}

	// Use the client to instantiate a producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Send 10 messages synchronously and 10 messages asynchronously
	for i := 0; i < 10; i++ {
		// Create a message
		msg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("message-%d", i)),
		}

		// Attempt to send the message
		if err := producer.Send(ctx, msg); err != nil {
			log.Fatal(err)
		}

		// Create a different message to send asynchronously
		asyncMsg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("async-message-%d", i)),
		}

		// Attempt to send the message asynchronously and handle the response
		producer.SendAsync(ctx, asyncMsg, func(msg pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("the %s successfully published", string(msg.Payload))
		})
	}
}
