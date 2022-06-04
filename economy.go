package economy

import (
	"errors"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/google/uuid"
	"github.com/provsalt/economy/handler"
	"github.com/provsalt/economy/provider"
)

// Economy is a struct that contains the economy provider and the event handler.
type Economy struct {
	p provider.Provider
	h handler.EconomyHandler
}

var ErrEventCancelled = errors.New("Event cancelled")

// New creates a new economy instance with a provider.
func New(p provider.Provider, h handler.EconomyHandler) *Economy {
	return &Economy{
		p,
		h,
	}
}

// Handle adds a new handler to handle economy changes.
func (e *Economy) Handle(h handler.EconomyHandler) {
	e.h = h
}

// Balance ...
func (e *Economy) Balance(UUID uuid.UUID) (uint64, error) {
	return e.p.Balance(UUID.String())
}

// Set ...
func (e *Economy) Set(UUID uuid.UUID, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, UUID, handler.ChangeTypeSet, amount)
	if ctx.Cancelled() {
		return ErrEventCancelled
	}
	return e.p.Set(UUID.String(), amount)
}

// Increase ...
func (e *Economy) Increase(UUID uuid.UUID, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, UUID, handler.ChangeTypeIncrease, amount)
	if ctx.Cancelled() {
		return ErrEventCancelled
	}
	bal, err := e.p.Balance(UUID.String())
	if err != nil {
		return err
	}
	return e.p.Set(UUID.String(), bal+amount)
}

// Decrease ...
func (e *Economy) Decrease(UUID uuid.UUID, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, UUID, handler.ChangeTypeDecrease, amount)
	if ctx.Cancelled() {
		return ErrEventCancelled
	}
	bal, err := e.p.Balance(UUID.String())
	if err != nil {
		return err
	}
	return e.p.Set(UUID.String(), bal-amount)
}

// Close ...
func (e *Economy) Close() error {
	return e.p.Close()
}
