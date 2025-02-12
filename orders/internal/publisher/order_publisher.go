package publisher

import (
	"encoding/json"
	"log"
	"ticketing/orders/internal/domain"
	"ticketing/orders/internal/model"

	"github.com/nats-io/nats.go"
)

type OrderPublisher interface {
	Created(order *domain.Order) error
	Cancelled(order *domain.Order) error
}

type OrderPublisherImpl struct {
	NatsConn *nats.Conn
}

func NewOrderPublisher(natsConn *nats.Conn) OrderPublisher {
	return &OrderPublisherImpl{
		NatsConn: natsConn,
	}
}

func (p *OrderPublisherImpl) Created(order *domain.Order) error {
	message := model.OrderCreatedEvent{
		ID:        order.ID,
		Status:    order.Status,
		UserID:    order.UserID,
		Ticket:    order.Ticket,
		ExpiresAt: order.ExpiresAt,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.NatsConn.Publish(domain.OrderCreated, data)
	if err != nil {
		return err
	}

	log.Printf("Published event on subject: %s", domain.OrderCreated)

	return nil
}

func (p *OrderPublisherImpl) Cancelled(order *domain.Order) error {
	message := model.OrderCancelledEvent{
		ID:       order.ID,
		TicketID: order.TicketID,
		UserID:   order.UserID,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.NatsConn.Publish(domain.OrderCancelled, data)
	if err != nil {
		return err
	}

	log.Printf("Published event on subject: %s", domain.OrderCancelled)
	return nil
}
