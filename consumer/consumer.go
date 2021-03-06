package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "gokafka_topic", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	// message, _ := conn.ReadMessage(1e6)
	// fmt.Println(string(message.Value))//First message read
	batch := conn.ReadBatch(1e3, 1e9) //1e3 1000bytes
	bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
