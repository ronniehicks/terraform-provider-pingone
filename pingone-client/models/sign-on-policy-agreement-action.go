package models

type AgreementAction struct {
	Agreement map[string]string `json:"agreement,omitempty" mapstructure:"agreement,omitempty"`
}
