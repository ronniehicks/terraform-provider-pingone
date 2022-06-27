package models

type SignOnPolicy struct {
	Environment map[string]string `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	ID          *string           `json:"id,omitempty" mapstructure:"policy_id,omitempty"`
	Name        *string           `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description *string           `json:"description,omitempty" mapstructure:"description,omitempty"`
	Default     *bool             `json:"default,omitempty" mapstructure:"default,omitempty"`
}

type SignOnPolicyCondition struct {
	Not          []SubCondition `json:"not,omitempty" mapstructure:"not,omitempty"`
	Or           []SubCondition `json:"or,omitempty" mapstructure:"or,omitempty"`
	And          []SubCondition `json:"and,omitempty" mapstructure:"and,omitempty"`
	SubCondition `mapstructure:",squash"`
}

type SubCondition struct {
	Value        *string  `json:"value,omitempty" mapstructure:"value,omitempty"`
	Equals       *string  `json:"equals,omitempty" mapstructure:"equals,omitempty"`
	Contains     *string  `json:"contains,omitempty" mapstructure:"contains,omitempty"`
	IpRange      []string `json:"ipRange,omitempty" mapstructure:"ip_range,omitempty"`
	Greater      *int     `json:"greater,omitempty" mapstructure:"greater,omitempty"`
	SecondsSince *string  `json:"secondsSince,omitempty" mapstructure:"seconds_since,omitempty"`
}

type SignOnPolicyAction struct {
	Environment  map[string]string      `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	SignOnPolicy map[string]string      `json:"signOnPolicy,omitempty" mapstructure:"sign_on_policy,omitempty"`
	ID           *string                `json:"id,omitempty" mapstructure:"action_id,omitempty"`
	Priority     *int                   `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	Type         *string                `json:"type,omitempty" mapstructure:"type,omitempty"`
	Condition    *SignOnPolicyCondition `json:"condition,omitempty" mapstructure:"condition,omitempty"`

	// Login & Identifier First props
	ConfirmIdentityProviderAttributes  *bool                           `json:"confirmIdentityProviderAttributes,omitempty" mapstructure:"confirm_identity_provider_attributes,omitempty"`
	EnforceLockoutForIdentityProviders *bool                           `json:"enforceLockoutForIdentityProviders,omitempty" mapstructure:"enforce_lockout_for_identity_providers,omitempty"`
	Recovery                           map[string]bool                 `json:"recovery,omitempty" mapstructure:"recovery,omitempty"`
	Registration                       *SignOnPolicyActionRegistration `json:"registration,omitempty" mapstructure:"registration,omitempty"`
	SocialProviders                    []map[string]string             `json:"socialProviders,omitempty" mapstructure:"social_providers,omitempty"`

	// Identifier First & Identity Provider prop
	IdentityProvider map[string]string `json:"identityProvider,omitempty" mapstructure:"identity_provider,omitempty"`

	AgreementAction                 `mapstructure:",squash"`
	IdentifierFirstAction           `mapstructure:",squash"`
	IdentityProviderAction          `mapstructure:",squash"`
	MultiFactorAuthenticationAction `mapstructure:",squash"`
	ProgressiveProfilingAction      `mapstructure:",squash"`
}

type SignOnPolicyActionRegistration struct {
	Enabled                           bool              `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	External                          map[string]string `json:"external,omitempty" mapstructure:"external,omitempty"`
	Population                        map[string]string `json:"population,omitempty" mapstructure:"population,omitempty"`
	ConfirmIdentityProviderAttributes bool              `json:"confirmIdentityProviderAttributes,omitempty" mapstructure:"confirm_identity_provider_attributes,omitempty"`
}
