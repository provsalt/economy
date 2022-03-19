package economy

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/provsalt/economy/handler"
	"github.com/provsalt/economy/provider"
)

// Economy is a struct that contains the economy provider and the event handler.
type Economy struct {
	p provider.Provider
	h handler.EconomyHandler
}

// New creates a new economy instance with a provider.
func New(p provider.Provider) *Economy {
	return &Economy{
		p,
		handler.NopEconomyHandler{},
	}
}

// Handle adds a new handler to handle economy changes.
func (e *Economy) Handle(h handler.EconomyHandler) {
	e.h = h
}

// Balance ...
func (e *Economy) Balance(XUID string) (uint64, error) {
	return e.p.Balance(XUID)
}

// Set ...
func (e *Economy) Set(XUID string, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, XUID, handler.ChangeTypeSet, amount)
	var err error
	ctx.Continue(func() {
		err = e.p.Set(XUID, amount)
	})
	return err
}

// Increase ...
func (e *Economy) Increase(XUID string, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, XUID, handler.ChangeTypeIncrease, amount)
	var err error
	ctx.Continue(func() {
		err = e.p.Increase(XUID, amount)
	})
	return err
}

// Decrease ...
func (e *Economy) Decrease(XUID string, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, XUID, handler.ChangeTypeDecrease, amount)
	var err error
	ctx.Continue(func() {
		err = e.p.Decrease(XUID, amount)
	})
	return err
}

// Close ...
func (e *Economy) Close() error {
	return e.p.Close()
}
