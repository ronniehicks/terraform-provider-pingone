package models

type IdentityProviderAction struct {
	AcrValues       *string `json:"acrValues,omitempty" mapstructure:"acr_values,omitempty"`
	PassUserContext *bool   `json:"passUserContext,omitempty" mapstructure:"pass_user_context,omitempty"`
}
