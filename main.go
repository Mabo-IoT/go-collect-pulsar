package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar/pulsar-client-go/pulsar"
	"github.com/mabo-iot/go-collect-pulsar/util"
	"log"
	"runtime"
	"time"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                     "pulsar://localhost:6650",
		OperationTimeoutSeconds: 5,
		MessageListenerThreads:  runtime.NumCPU(),
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	for {
		SendMessage(client)
	}
	// RecvMessage(client)

}

func SendMessage(client pulsar.Client) {
	ctx := context.Background()
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "topic-1",
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar producer: %v", err)
	}

	defer producer.Close()

	data := util.Generate()

	msg := pulsar.ProducerMessage{
		Payload: data,
	}

	if err := producer.Send(ctx, msg); err != nil {
		log.Fatalf("Producer could not send message: %v", err)
	}
	fmt.Println("send message to topic")
	time.Sleep(1 * time.Second)
}

// func RecvMessage(client pulsar.Client) {
// 	msgChannel := make(chan pulsar.ConsumerMessage, 1)

// 	consumerOpts := pulsar.ConsumerOptions{
// 		Topic:            "topic",
// 		SubscriptionName: "my-subscription-1",
// 		Type:             pulsar.Exclusive,
// 		MessageChannel:   msgChannel,
// 	}

// 	consumer, err := client.Subscribe(consumerOpts)

// 	if err != nil {
// 		log.Fatalf("Could not establish subscription: %v", err)
// 	}

// 	defer consumer.Close()

// 	for cm := range msgChannel {
// 		msg := cm.Message

// 		fmt.Printf("Message ID: %s", msg.ID())
// 		fmt.Printf("Message value: %s", string(msg.Payload()))

// 		consumer.Ack(msg)
// 	}

// }
