package events

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Event represents a generic event in the system
type Event interface{}

// Handler is a function that handles an event
type Handler func(ctx context.Context, event Event) error

// Bus is a publish/subscribe event bus
type Bus struct {
	subscribers map[string][]Handler
	mu          sync.RWMutex
	logger      *zap.Logger
}

// New creates a new event bus
func New(logger *zap.Logger) *Bus {
	return &Bus{
		subscribers: make(map[string][]Handler),
		logger:      logger,
	}
}

// Subscribe registers a handler for an event type
func (b *Bus) Subscribe(eventType string, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[eventType] = append(b.subscribers[eventType], handler)
	b.logger.Debug("Subscribed to event", zap.String("type", eventType))
}

// Publish sends an event to all subscribers
func (b *Bus) Publish(ctx context.Context, eventType string, event Event) error {
	b.mu.RLock()
	handlers, exists := b.subscribers[eventType]
	b.mu.RUnlock()

	if !exists {
		b.logger.Debug("No subscribers for event", zap.String("type", eventType))
		return nil
	}

	// Call all handlers
	var errs []error
	for _, handler := range handlers {
		if err := handler(ctx, event); err != nil {
			b.logger.Error(
				"Handler failed",
				zap.String("event_type", eventType),
				zap.Error(err),
			)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("event handlers failed: %v", errs)
	}

	return nil
}
