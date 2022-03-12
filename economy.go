package economy

import "github.com/provsalt/economy/provider"

type Economy struct {
	provider.Provider
}

func New(p provider.Provider) *Economy {
	return &Economy{
		p,
	}
}
