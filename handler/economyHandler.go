package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/google/uuid"
)

// ChangeType is the type of change that has occurred.
const (
	ChangeTypeSet = iota
	ChangeTypeIncrease
	ChangeTypeDecrease
)

// EconomyHandler is a handler that handles economy changes.
type EconomyHandler interface {
	// HandleChange handles an economy change. The change is of the type ChangeType.
	HandleChange(ctx *event.Context, UUID uuid.UUID, _type int, amount uint64)
}

// NopEconomyHandler is a handler that does not handle any economy changes when called.
type NopEconomyHandler struct{}

// HandleChange ...
func (n NopEconomyHandler) HandleChange(*event.Context, uuid.UUID, int, uint64) {}
