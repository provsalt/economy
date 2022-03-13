package handler

import "github.com/df-mc/dragonfly/dragonfly/event"

// ChangeType is the type of change that has occurred.
const (
	ChangeTypeSet = iota
	ChangeTypeIncrease
	ChangeTypeDecrease
)

// EconomyHandler is a handler that handles economy changes.
type EconomyHandler interface {
	// HandleChange handles an economy change. The change is of the type ChangeType.
	HandleChange(ctx *event.Context, XUID string, _type int, amount uint64)
}

// NopEconomyHandler is a handler that does not handle any economy changes when called.
type NopEconomyHandler struct{}

// HandleChange ...
func (n NopEconomyHandler) HandleChange(*event.Context, string, int, uint64) {}
