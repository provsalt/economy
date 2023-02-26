package provider

import "io"

// Provider is an interface that defines the methods that a provider must implement.
type Provider interface {
	// Balance returns the balance of the player with the given UUID.
	Balance(UUID string) (uint64, error)

	// Set sets the balance of the player with the given UUID to the given value.
	Set(UUID string, value uint64) error

	io.Closer
}
