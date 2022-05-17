package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
)

func main() {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: -1,
		MaxAttempts:  3,
	}

	newss := &[]models.News{
		{
			NewsID:      1,
			GenreID:     1,
			Description: "Hello There",
		},
		{NewsID: 1,
			GenreID:     1,
			Description: "Hello There",
		},
	}

	newsBytes, err := json.Marshal(&newss)
	if err != nil {
		log.Fatal(err)
	}

	msg := kafka.Message{
		Value: newsBytes,
	}
	err = w.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
