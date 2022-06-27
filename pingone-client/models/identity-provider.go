package models

type IdpVerification struct {
	Certificates []map[string]string `json:"certificates,omitempty" mapstructure:"certificates,omitempty"`
	// handle the mapping in the expander to get to id
}

type SpSigning struct {
	Key map[string]string `json:"key,omitempty" mapstructure:"key,omitempty"`
}

type Registration struct {
	Population map[string]string `json:"population,omitempty" mapstructure:"population,omitempty"`
}

type ProviderAttribute struct {
	ID               *string           `json:"id,omitempty" mapstructure:"attribute_id,omitempty"`
	MappingType      *string           `json:"mappingType,omitempty" mapstructure:"mapping_type,omitempty"`
	Name             *string           `json:"name,omitempty" mapstructure:"name,omitempty"`
	Value            *string           `json:"value,omitempty" mapstructure:"value,omitempty"`
	Update           *string           `json:"update,omitempty" mapstructure:"update,omitempty"`
	IdentityProvider map[string]string `json:"identityProvider,omitempty" mapstructure:"identity_provider,omitempty"`
}

type IdentityProvider struct {
	// Base IdP data model
	ID                 *string              `json:"id,omitempty" mapstructure:"identity_provider_id,omitempty"`
	Name               *string              `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description        *string              `json:"description,omitempty" mapstructure:"description,omitempty"`
	Enabled            *bool                `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	Environment        map[string]string    `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	Icon               *Icon                `json:"icon" mapstructure:"icon,omitempty"`
	LoginButtonIcon    *Icon                `json:"loginButtonIcon" mapstructure:"login_button_icon,omitempty"`
	Registration       *Registration        `json:"registration,omitempty" mapstructure:"registration,omitempty"`
	Type               *string              `json:"type,omitempty" mapstructure:"type,omitempty"`
	ProviderAttributes *[]ProviderAttribute `json:"providerAttributes,omitempty" mapstructure:"provider_attributes,omitempty"`

	// OpenID Connect Request Model
	AuthorizationEndpoint   *string  `json:"authorizationEndpoint,omitempty" mapstructure:"authorization_endpoint,omitempty"`
	ClientId                *string  `json:"clientId,omitempty" mapstructure:"client_id,omitempty"`
	ClientSecret            *string  `json:"clientSecret,omitempty" mapstructure:"client_secret,omitempty"`
	DiscoveryEndpoint       *string  `json:"discoveryEndpoint,omitempty" mapstructure:"discovery_endpoint,omitempty"`
	Issuer                  *string  `json:"issuer,omitempty" mapstructure:"issuer,omitempty"`
	JwksEndpoint            *string  `json:"jwksEndpoint,omitempty" mapstructure:"jwks_endpoint,omitempty"`
	Scopes                  []string `json:"scopes,omitempty" mapstructure:"scopes,omitempty"`
	TokenEndpoint           *string  `json:"tokenEndpoint,omitempty" mapstructure:"token_endpoint,omitempty"`
	TokenEndpointAuthMethod *string  `json:"tokenEndpointAuthMethod,omitempty" mapstructure:"token_endpoint_auth_method,omitempty"`
	UserInfoEndpoint        *string  `json:"userInfoEndpoint,omitempty" mapstructure:"user_info_endpoint,omitempty"`

	// SAML Request Model
	AuthnRequestSigned *bool            `json:"authnRequestSigned,omitempty" mapstructure:"authn_request_signed,omitempty"`
	IdpEntityId        *string          `json:"idpEntityId,omitempty" mapstructure:"idp_entity_id,omitempty"`
	IdpVerification    *IdpVerification `json:"idpVerification" mapstructure:"idp_verification,omitempty"`
	SpEntityId         *string          `json:"spEntityId,omitempty" mapstructure:"sp_entity_id,omitempty"`
	SpSigning          *SpSigning       `json:"spSigning,omitempty" mapstructure:"sp_signing,omitempty"`
	SsoBinding         *string          `json:"ssoBinding,omitempty" mapstructure:"sso_binding,omitempty"`
	SsoEndpoint        *string          `json:"ssoEndpoint,omitempty" mapstructure:"sso_endpoint,omitempty"`
}
