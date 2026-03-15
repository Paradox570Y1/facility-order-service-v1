package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/dto"
)

func ConsumerWorker(ctx context.Context, brokerAddr string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerAddr},
		Topic:       "order_created",
		GroupID:     "order-logger",
		StartOffset: kafka.LastOffset,
	})
	defer reader.Close()

	log.Println("Consumer started, waiting for messages...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer stopped")
			return
		default:
			message, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}

			var orderMsg dto.OrderCreatedMessage
			err = json.Unmarshal(message.Value, &orderMsg)
			if err != nil {
				log.Printf("Error parsing message: %v", err)
				continue
			}

			log.Printf("Order Created Event Received:")
			log.Printf("   - Order ID: %s", orderMsg.OrderID)
			log.Printf("   - Facility Code: %s", orderMsg.FacilityCode)
			log.Printf("   ─────────────────────────────────")
		}
	}
}
