// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package forwarder

import "github.com/DataDog/datadog-agent/pkg/forwarder/transaction"

type DestinationType int

const (
	Datadog DestinationType = iota
	Vector
)

type DomainResolver interface {
	Resolve(endpoint transaction.Endpoint) (string, DestinationType)
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

func NewSingleDomainResolver(domain string, apiKeys []string) *SingleDomainResolver {
	return &SingleDomainResolver{
		domain,
		apiKeys,
	}
}

func NewSingleDomainResolvers(keysPerDomain map[string][]string) map[string]DomainResolver {
	resolvers := make(map[string]DomainResolver)
	for domain, keys := range keysPerDomain {
		resolvers[domain] = NewSingleDomainResolver(domain, keys)
	}
	return resolvers
}

func (r *SingleDomainResolver) Resolve(endpoint transaction.Endpoint) (string, DestinationType) {
	return r.domain, Datadog
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

type Destination struct {
	domain string
	dType  DestinationType
}

// MultiDomainResolver holds a default value and can provide alternate domain for some route
type MultiDomainResolver struct {
	baseDomain string
	apiKeys    []string
	// endpoint name => overriden hostname map
	overrides map[string]Destination
}

func NewMultiDomainResolver(baseDomain string, apiKeys []string) *MultiDomainResolver {
	return &MultiDomainResolver{
		baseDomain,
		apiKeys,
		make(map[string]Destination),
	}
}

func (r *MultiDomainResolver) GetApiKeys() []string {
	return r.apiKeys
}

func (s *MultiDomainResolver) Resolve(endpoint transaction.Endpoint) (string, DestinationType) {
	if d, ok := s.overrides[endpoint.Name]; ok {
		return d.domain, d.dType
	}
	return s.baseDomain, Datadog
}

func (r *MultiDomainResolver) GetBaseDomain() string {
	return r.baseDomain
}

func (r *MultiDomainResolver) SetBaseDomain(domain string) {
	r.baseDomain = domain
}

func (r *MultiDomainResolver) GetAlternateDomains() []string {
	domains := make([]string, 0, len(r.overrides))
	for _, dest := range r.overrides {
		domains = append(domains, dest.domain)
	}
	return domains
}

func NewDomainResolverWithMetricToVector(mainEndpoint string, apiKeys []string, vectorEndpoint string) *MultiDomainResolver {
	dest := Destination{
		vectorEndpoint,
		Vector,
	}
	overrides := map[string]Destination{
		v1SeriesEndpoint.Name:     dest,
		seriesEndpoint.Name:       dest,
		sketchSeriesEndpoint.Name: dest,
	}
	return &MultiDomainResolver{
		mainEndpoint,
		apiKeys,
		overrides,
	}
}
