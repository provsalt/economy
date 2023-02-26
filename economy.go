package economy

import (
	"context"
	"errors"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	ristretto2 "github.com/eko/gocache/store/ristretto/v4"
	"github.com/google/uuid"
	"github.com/provsalt/economy/handler"
	"github.com/provsalt/economy/provider"
)

// Economy is a struct that contains the economy provider and the event handler.
type Economy struct {
	p provider.Provider
	h handler.EconomyHandler
	c cache.Cache[uint64]
}

// ErrEventCancelled is an error that is returned when the event is cancelled.
var ErrEventCancelled = errors.New("event cancelled")

// New creates a new economy instance with a provider.
func New(p provider.Provider, h handler.EconomyHandler) *Economy {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1 << 16,
		MaxCost:     1 << 8,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := ristretto2.NewRistretto(ristrettoCache)

	cacheManager := cache.New[uint64](ristrettoStore)
	return &Economy{
		p,
		h,
		*cacheManager,
	}
}

// Handle adds a new handler to handle economy changes.
func (e *Economy) Handle(h handler.EconomyHandler) {
	e.h = h
}

// Balance returns the balance of a player.
func (e *Economy) Balance(UUID uuid.UUID) (uint64, error) {
	ctx := context.Background()
	d, err := e.c.Get(ctx, UUID)
	if err != nil {
		bal, err := e.p.Balance(UUID.String())
		if err != nil {
			return 0, err
		}
		err = e.c.Set(ctx, UUID, bal)
		if err != nil {
			return 0, err
		}
		return bal, nil
	}
	return d, nil
}

// Set sets the balance of a player to a given value.
// You should be using this function to initialize a player's balance or Balance function will throw an error.
func (e *Economy) Set(UUID uuid.UUID, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, UUID, handler.ChangeTypeSet, amount)
	if ctx.Cancelled() {
		return ErrEventCancelled
	}
	err := e.c.Set(context.Background(), UUID, amount)
	if err != nil {
		return err
	}
	return e.p.Set(UUID.String(), amount)
}

// Increase is a wrapper for Set that increases the balance of a player.
func (e *Economy) Increase(UUID uuid.UUID, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, UUID, handler.ChangeTypeIncrease, amount)
	if ctx.Cancelled() {
		return ErrEventCancelled
	}
	bal, err := e.Balance(UUID)
	if err != nil {
		return err
	}
	return e.Set(UUID, bal+amount)
}

// Decrease is a wrapper for Set that decreases the balance of a player.
func (e *Economy) Decrease(UUID uuid.UUID, amount uint64) error {
	ctx := event.C()
	e.h.HandleChange(ctx, UUID, handler.ChangeTypeDecrease, amount)
	if ctx.Cancelled() {
		return ErrEventCancelled
	}
	bal, err := e.Balance(UUID)
	if err != nil {
		return err
	}
	return e.Set(UUID, bal-amount)
}

// Close closes the economy provider.
func (e *Economy) Close() error {
	return e.p.Close()
}
