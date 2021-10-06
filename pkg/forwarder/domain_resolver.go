// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package forwarder

import "github.com/DataDog/datadog-agent/pkg/forwarder/transaction"

type DomainResolver interface {
	// Maybe allow a API key override Resolve(route) (string, string) ?
	// Or even return (host, api key)
	Resolve(endpoint transaction.Endpoint) string
	GetApiKeys() []string
	GetBaseDomain() string
	GetAlternateDomains() []string
	SetBaseDomain(domain string)
}

// SingleDomainResolver will always return the same host
type SingleDomainResolver struct {
	domain  string
	apiKeys []string
}

func NewSingleDomainResolver(domain string, apiKeys []string) DomainResolver {
	return &SingleDomainResolver{
		domain,
		apiKeys,
	}
}

func (r *SingleDomainResolver) Resolve(endpoint transaction.Endpoint) string {
	return r.domain
}

func (r *SingleDomainResolver) GetBaseDomain() string {
	return r.domain
}

func (r *SingleDomainResolver) GetApiKeys() []string {
	return r.apiKeys
}

func (r *SingleDomainResolver) SetBaseDomain(domain string) {
	r.domain = domain
}

func (r *SingleDomainResolver) GetAlternateDomains() []string {
	return []string{}
}

// MultiDomainResolver holds a default value and can provide alternate domain for some route
type MultiDomainResolver struct {
	baseDomain string
	apiKeys    []string
	// Route => overriden hostname map
	overrides map[string]string
}

func NewMultiDomainFormarder(baseDomain string, apiKeys []string, overrides map[string]string) DomainResolver {
	return &MultiDomainResolver{
		baseDomain,
		apiKeys,
		overrides,
	}
}

func (r *MultiDomainResolver) GetApiKeys() []string {
	return r.apiKeys
}

func (s *MultiDomainResolver) Resolve(endpoint transaction.Endpoint) string {
	if d, ok := s.overrides[endpoint.Name]; ok {
		return d
	}
	return s.baseDomain
}

func (r *MultiDomainResolver) GetBaseDomain() string {
	return r.baseDomain
}

func (r *MultiDomainResolver) SetBaseDomain(domain string) {
	r.baseDomain = domain
}

func (r *MultiDomainResolver) GetAlternateDomains() []string {
	domains := make([]string, 0, len(r.overrides))
	for _, domain := range r.overrides {
		domains = append(domains, domain)
	}
	return domains
}
