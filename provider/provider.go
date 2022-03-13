package provider

// Provider is an interface that defines the methods that a provider must implement.
type Provider interface {
	Balance(XUID string) (uint64, error)

	Set(XUID string, value uint64) error

	Decrease(XUID string, value uint64) error

	Increase(XUID string, value uint64) error

	Close() error
}
