package oidc

import (
	"fmt"
	"strings"
)

type ProviderRegistry struct {
	providers map[string]OIDCProvider
}

func NewProviderRegistry(providers ...OIDCProvider) *ProviderRegistry {
	registry := &ProviderRegistry{providers: make(map[string]OIDCProvider)}
	for _, provider := range providers {
		registry.Register(provider)
	}
	return registry
}

func (r *ProviderRegistry) Register(provider OIDCProvider) {
	if provider == nil {
		return
	}
	r.providers[strings.ToLower(provider.Name())] = provider
}

func (r *ProviderRegistry) Get(name string) (OIDCProvider, error) {
	provider, ok := r.providers[strings.ToLower(strings.TrimSpace(name))]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrProviderNotFound, name)
	}
	return provider, nil
}
