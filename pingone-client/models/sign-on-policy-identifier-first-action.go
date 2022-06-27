package models

type IdentifierFirstAction struct {
	DiscoveryRules []DiscoveryRule `json:"discoveryRules,omitempty" mapstructure:"discovery_rules,omitempty"`
}

type DiscoveryRule struct {
	Condition        map[string]string `json:"condition,omitempty" mapstructure:"condition,omitempty"`
	IdentityProvider map[string]string `json:"identityProvider,omitempty" mapstructure:"identity_provider,omitempty"`
}
