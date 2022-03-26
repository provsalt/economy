package provider

import "io"

// Provider is an interface that defines the methods that a provider must implement.
type Provider interface {
	io.Closer

	Balance(UUID string) (uint64, error)

	Set(UUID string, value uint64) error

	Decrease(UUID string, value uint64) error

	Increase(UUID string, value uint64) error

	Close() error
}
