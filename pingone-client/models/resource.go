package models

type Resource struct {
	ID                         *string `json:"id,omitempty" mapstructure:"resource_id,omitempty"`
	Name                       *string `json:"name,omitempty" mapstructure:"name,omitempty"`
	Description                *string `json:"description,omitempty" mapstructure:"description,omitempty"`
	Type                       *string `json:"type,omitempty" mapstructure:"type,omitempty"`
	Audience                   *string `json:"audience,omitempty" mapstructure:"audience,omitempty"`
	AccessTokenValiditySeconds *int    `json:"accessTokenValiditySeconds,omitempty" mapstructure:"access_token_validity_seconds,omitempty"`
}

type ResourceAttribute struct {
	Environment map[string]string `json:"environment,omitempty" mapstructure:"environment,omitempty"`
	Resource    map[string]string `json:"resource,omitempty" mapstructure:"resource,omitempty"`
	ID          *string           `json:"id,omitempty" mapstructure:"attribute_id,omitempty"`
	Name        *string           `json:"name,omitempty" mapstructure:"name,omitempty"`
	Type        *string           `json:"type,omitempty" mapstructure:"type,omitempty"`
	Value       *string           `json:"value,omitempty" mapstructure:"value,omitempty"`
}
