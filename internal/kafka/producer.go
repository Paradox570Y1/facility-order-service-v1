package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/dto"
)

func PublishOrderCreated(ctx context.Context, brokerAddr string, orderID string, facilityCode string) error {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(brokerAddr),
		Topic: "order_created",
	}
	defer writer.Close()

	message := dto.OrderCreatedMessage{
		OrderID:      orderID,
		FacilityCode: facilityCode,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error encoding message: %v", err)
		return err
	}

	err = writer.WriteMessages(ctx, kafka.Message{
		Value: messageBytes,
	})

	if err != nil {
		log.Printf("Error publishing to Kafka: %v", err)
		return err
	}

	log.Printf("Published order created: order_id=%s, facility_code=%s", orderID, facilityCode)
	return nil
}
