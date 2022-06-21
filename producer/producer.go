package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "gokafka_topic", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
	for{
		conn.WriteMessages(kafka.Message{Value: []byte("Test Kafka Console")})
	}
	
}
