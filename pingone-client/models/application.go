package models

type Icon struct {
	ID   *string `json:"id,omitempty" mapstructure:"id,omitempty"`
	Href *string `json:"href,omitempty" mapstructure:"href,omitempty"`
}

type ApplicationAccessControlGroup struct {
	Type   *string             `json:"type,omitempty" mapstructure:"type,omitempty"`
	Groups []map[string]string `json:"groups,omitempty" mapstructure:"groups,omitempty"`
}

type ApplicationAccessControl struct {
	Role  map[string]string              `json:"role,omitempty" mapstructure:"role,omitempty"`
	Group *ApplicationAccessControlGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
}

type ApplicationMobileIntegrityDetectionCacheDuration struct {
	Amount *int    `json:"amount,omitempty" mapstructure:"amount,omitempty"`
	Units  *string `json:"units,omitempty" mapstructure:"units,omitempty"`
}

type ApplicationMobileIntegrityDetection struct {
	Mode          *string                                           `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	CacheDuration *ApplicationMobileIntegrityDetectionCacheDuration `json:"cacheDuration,omitempty" mapstructure:"cache_duration,omitempty"`
}

type ApplicationMobilePasscodeRefreshDuration struct {
	Duration *int    `json:"duration,omitempty" mapstructure:"duration,omitempty"`
	TimeUnit *string `json:"timeUnit,omitempty" mapstructure:"time_unit,omitempty"`
}

type ApplicationMobile struct {
	BundleID                *string                                   `json:"bundleId,omitempty" mapstructure:"bundle_id,omitempty"`
	PackageName             *string                                   `json:"packageName,omitempty" mapstructure:"package_name,omitempty"`
	IntegrationDetection    *ApplicationMobileIntegrityDetection      `json:"integrationDetection,omitempty" mapstructure:"integration_detection,omitempty"`
	PasscodeRefreshDuration *ApplicationMobilePasscodeRefreshDuration `json:"passcodeRefreshDuration,omitempty" mapstructure:"passcode_refresh_duration,omitempty"`
}

type Application struct {
	Environment      map[string]string         `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	ID               *string                   `json:"id,omitempty" mapstructure:"application_id,omitempty"`
	Name             *string                   `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description      *string                   `json:"description,omitempty" mapstructure:"description,omitempty"`
	Enabled          *bool                     `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	Type             *string                   `json:"type,omitempty" mapstructure:"type,omitempty"`
	Protocol         *string                   `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	HomePageUrl      *string                   `json:"homePageUrl,omitempty" mapstructure:"home_page_url,omitempty"`
	LoginPageUrl     *string                   `json:"loginPageUrl,omitempty" mapstructure:"login_page_url,omitempty"`
	Icon             *Icon                     `json:"icon,omitempty" mapstructure:"icon,omitempty"`
	AssignActorRoles *bool                     `json:"assignActorRoles,omitempty" mapstructure:"assign_actor_roles,omitempty"`
	AccessControl    *ApplicationAccessControl `json:"accessControl,omitempty" mapstructure:"access_control,omitempty"`

	// OIDC
	GrantTypes                   []*string `json:"grantTypes,omitempty" mapstructure:"grant_types,omitempty"`
	PostLogoutRedirectUris       []*string `json:"postLogoutRedirectUris,omitempty" mapstructure:"post_logout_redirect_uris,omitempty"`
	RedirectUris                 []*string `json:"redirectUris,omitempty" mapstructure:"redirect_uris,omitempty"`
	ResponseTypes                []*string `json:"responseTypes,omitempty" mapstructure:"response_types,omitempty"`
	TokenEndpointAuthMethod      *string   `json:"tokenEndpointAuthMethod,omitempty" mapstructure:"token_endpoint_auth_method,omitempty"`
	PkceEnforcement              *string   `json:"pkceEnforcement,omitempty" mapstructure:"pkce_enforcement,omitempty"`
	RefreshTokenDuration         *int      `json:"refreshTokenDuration,omitempty" mapstructure:"refresh_token_duration,omitempty"`
	RefreshTokenRollingDuration  *int      `json:"refreshTokenRollingDuration,omitempty" mapstructure:"refresh_token_rolling_duration,omitempty"`
	SupportUnsignedRequestObject *bool     `json:"supportUnsignedRequestObject,omitempty" mapstructure:"support_unsigned_request_object,omitempty"`

	// Native App
	BundleID    *string            `json:"bundleId,omitempty" mapstructure:"bundle_id,omitempty"`
	PackageName *string            `json:"packageName,omitempty" mapstructure:"package_name,omitempty"`
	Mobile      *ApplicationMobile `json:"mobile,omitempty" mapstructure:"mobile,omitempty"`

	// SAML
	SpEntityId          *string                    `json:"spEntityId,omitempty" mapstructure:"sp_entity_id,omitempty"`
	AcsUrls             []*string                  `json:"acsUrls,omitempty" mapstructure:"acs_urls,omitempty"`
	AssertionDuration   *int                       `json:"assertionDuration,omitempty" mapstructure:"assertion_duration,omitempty"`
	DefaultTargetUrl    *string                    `json:"defaultTargetUrl,omitempty" mapstructure:"default_target_url,omitempty"`
	SloBinding          *string                    `json:"sloBinding,omitempty" mapstructure:"slo_binding,omitempty"`
	SloEndpoint         *string                    `json:"sloEndpoint,omitempty" mapstructure:"slo_endpoint,omitempty"`
	SloResponseEndpoint *string                    `json:"sloResponseEndpoint,omitempty" mapstructure:"slo_response_endpoint,omitempty"`
	ResponseSigned      *bool                      `json:"responseSigned,omitempty" mapstructure:"response_signed,omitempty"`
	AssertionSigned     *bool                      `json:"assertionSigned,omitempty" mapstructure:"assertion_signed,omitempty"`
	IdpSigning          *ApplicationIdpSigning     `json:"idpSigning,omitempty" mapstructure:"idp_signing,omitempty"`
	SpEncryption        *ApplicationSpEncryption   `json:"spEncryption,omitempty" mapstructure:"sp_encryption,omitempty"`
	SpVerification      *ApplicationSpVerification `json:"spVerification,omitempty" mapstructure:"sp_verification,omitempty"`
	NameIdFormat        *string                    `json:"nameIdFormat,omitempty" mapstructure:"name_id_format,omitempty"`
}

type ApplicationSpVerification struct {
	AuthnRequestSigned *bool                      `json:"authnRequestSigned,omitempty" mapstructure:"authn_request_signed,omitempty"`
	Certificates       []*ApplicationCertificates `json:"certificates,omitempty" mapstructure:"certificates,omitempty"`
}

type ApplicationIdpSigning struct {
	Key       map[string]string `json:"key,omitempty" mapstructure:"key,omitempty"`
	Algorithm *string           `json:"algorithm,omitempty" mapstructure:"algorithm,omitempty"`
}

type ApplicationCertificates struct {
	ID *string `json:"id,omitempty" mapstructure:"id,omitempty"`
}

type ApplicationSpEncryption struct {
	Algorithm    *string                  `json:"algorithm,omitempty" mapstructure:"algorithm,omitempty"`
	Certificates *ApplicationCertificates `json:"certificates,omitempty" mapstructure:"certificates,omitempty"`
}

type ApplicationAttribute struct {
	ID          *string           `json:"id,omitempty" mapstructure:"attribute_id,omitempty"`
	Name        *string           `json:"name,omitempty" mapstructure:"name,omitempty"`
	MappingType *string           `json:"mappingType,omitempty" mapstructure:"mapping_type,omitempty"`
	Application map[string]string `json:"application,omitempty" mapstructure:"application,omitempty"`
	Required    *bool             `json:"required,omitempty" mapstructure:"required,omitempty"`
	Value       *string           `json:"value,omitempty" mapstructure:"value,omitempty"`
}

type ApplicationGrant struct {
	ID          *string             `json:"id,omitempty" mapstructure:"grant_id,omitempty"`
	Environment map[string]string   `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	Application map[string]string   `json:"application,omitempty" mapstructure:"application,omitempty"`
	Resource    map[string]string   `json:"resource,omitempty" mapstructure:"resource,omitempty"`
	Scopes      []map[string]string `json:"scopes,omitempty" mapstructure:"scopes,omitempty"`
}

type ApplicationSignOnPolicyAssignment struct {
	ID           *string           `json:"id,omitempty" mapstructure:"assignment_id,omitempty"`
	Environment  map[string]string `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	Application  map[string]string `json:"application,omitempty" mapstructure:"application,omitempty"`
	SignOnPolicy map[string]string `json:"signOnPolicy,omitempty" mapstructure:"sign_on_policy,omitempty"`
	Priority     *int              `json:"priority,omitempty" mapstructure:"priority,omitempty"`
}
