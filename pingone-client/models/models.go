package models

type GenericResponse[T AllTheThings] struct {
	Links    map[string]map[string]string `json:"_links,omitempty"`
	Embedded map[string][]T               `json:"_embedded"`
}

type AllTheThings interface {
	Application |
	ApplicationAttribute |
	ApplicationGrant |
	ApplicationSignOnPolicyAssignment |
	Certificate |
	CertificateApplication |
	CustomDomain |
	Environment |
	Group |
	IdentityProvider |
	Key |
	KeyApplication |
	ProviderAttribute |
	Population |
	Resource |
	ResourceAttribute |
	Scope |
	SignOnPolicy |
	SignOnPolicyAction
}

type Environment struct {
	ID           *string           `json:"id,omitempty" mapstructure:"id"`
	Name         *string           `json:"name,omitempty" mapstructure:"name"`
	Description  *string           `json:"description,omitempty" mapstructure:"description,omitempty"`
	Type         *string           `json:"type,omitempty" mapstructure:"type,omitempty"`
	Region       *string           `json:"region,omitempty" mapstructure:"region,omitempty"`
	Organization map[string]string `json:"organization,omitempty" mapstructure:"organization,omitempty"`
	License      map[string]string `json:"license,omitempty" mapstructure:"license,omitempty"`
}

type Group struct {
	ID          *string                `json:"id,omitempty" mapstructure:"group_id,omitempty"`
	Name        *string                `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description *string                `json:"description,omitempty" mapstructure:"description,omitempty"`
	Population  map[string]string      `json:"population,omitempty" mapstructure:"population,omitempty"`
	UserFilter  *string                `json:"userFilter,omitempty" mapstructure:"user_filter,omitempty"`
	ExternalId  *string                `json:"externalId,omitempty" mapstructure:"external_id,omitempty"`
	CustomData  map[string]interface{} `json:"customData,omitempty" mapstructure:"custom_data,omitempty"`
	Type        *string                `json:"type,omitempty" mapstructure:"type,omitempty"`
}

type Population struct {
	ID             *string           `json:"id,omitempty" mapstructure:"population_id,omitempty"`
	Name           *string           `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description    *string           `json:"description,omitempty" mapstructure:"description,omitempty"`
	PasswordPolicy map[string]string `json:"passwordPolicy,omitempty" mapstructure:"password_policy,omitempty"`
}

type Scope struct {
	ID               *string           `json:"id,omitempty" mapstructure:"scope_id,omitempty"`
	Name             *string           `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description      *string           `json:"description,omitempty" mapstructure:"description,omitempty"`
	Resource         map[string]string `json:"resource,omitempty" mapstructure:"resource,omitempty"`
	SchemaAttributes []*string         `json:"schemaAttributes,omitempty" mapstructure:"schema_attributes,omitempty"`
}

type Certificate struct {
	Algorithm          *string `json:"algorithm,omitempty" mapstructure:"algorithm,omitempty"`
	CreatedAt          *string `json:"createdAt,omitempty" mapstructure:"created_at,omitempty"`
	Default            *bool   `json:"default,omitempty" mapstructure:"default,omitempty"`
	ExpiresAt          *string `json:"expiresAt,omitempty" mapstructure:"expires_at,omitempty"`
	ID                 *string `json:"id,omitempty" mapstructure:"certificate_id,omitempty"`
	IssuerDN           *string `json:"issuerDN,omitempty" mapstructure:"issuer_dn,omitempty"`
	KeyLength          *int    `json:"keyLength,omitempty" mapstructure:"key_length,omitempty"`
	Name               *string `json:"name,omitempty" mapstructure:"name,omitempty"`
	SerialNumber       *int    `json:"serialNumber,omitempty" mapstructure:"serial_number,omitempty"`
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" mapstructure:"signature_algorithm,omitempty"`
	StartsAt           *string `json:"startsAt,omitempty" mapstructure:"starts_at,omitempty"`
	Status             *string `json:"status,omitempty" mapstructure:"status,omitempty"`
	SubjectDN          *string `json:"subjectDN,omitempty" mapstructure:"subject_dn,omitempty"`
	UsageType          *string `json:"usageType,omitempty" mapstructure:"usage_type,omitempty"`
	ValidityPeriod     *int    `json:"validityPeriod,omitempty" mapstructure:"validity_period,omitempty"`
}

type Key struct {
	Algorithm          *string `json:"algorithm,omitempty" mapstructure:"algorithm,omitempty"`
	CreatedAt          *string `json:"createdAt,omitempty" mapstructure:"created_at,omitempty"`
	Default            *bool   `json:"default,omitempty" mapstructure:"default,omitempty"`
	ExpiresAt          *string `json:"expiresAt,omitempty" mapstructure:"expires_at,omitempty"`
	ID                 *string `json:"id,omitempty" mapstructure:"key_id,omitempty"`
	IssuerDN           *string `json:"issuerDN,omitempty" mapstructure:"issuer_dn,omitempty"`
	KeyLength          *int    `json:"keyLength,omitempty" mapstructure:"key_length,omitempty"`
	Name               *string `json:"name,omitempty" mapstructure:"name,omitempty"`
	SerialNumber       *int    `json:"serialNumber,omitempty" mapstructure:"serial_number,omitempty"`
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" mapstructure:"signature_algorithm,omitempty"`
	StartsAt           *string `json:"startsAt,omitempty" mapstructure:"starts_at,omitempty"`
	Status             *string `json:"status,omitempty" mapstructure:"status,omitempty"`
	SubjectDN          *string `json:"subjectDN,omitempty" mapstructure:"subject_dn,omitempty"`
	UsageType          *string `json:"usageType,omitempty" mapstructure:"usage_type,omitempty"`
	ValidityPeriod     *int    `json:"validityPeriod,omitempty" mapstructure:"validity_period,omitempty"`
}

type CertificateApplication struct {
	ID   *string `json:"id,omitempty" mapstructure:"application_id,omitempty"`
	Name *string `json:"name,omitempty" mapstructure:"name,omitempty"`
}

type KeyApplication struct {
	ID   *string `json:"id,omitempty" mapstructure:"application_id,omitempty"`
	Name *string `json:"name,omitempty" mapstructure:"name,omitempty"`
}
type CustomDomain struct {
	Environment   map[string]string `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	ID            *string           `json:"id,omitempty" mapstructure:"custom_domain_id,omitempty"`
	CanonicalName *string           `json:"canonicalName,omitempty" mapstructure:"canonical_name,omitempty"`
	DomainName    *string           `json:"domainName,omitempty" mapstructure:"domain_name,omitempty"`
	Status        *string           `json:"status,omitempty" mapstructure:"status,omitempty"`
	Certificate   map[string]string `json:"certificate,omitempty" mapstructure:"certificate,omitempty"`
}
