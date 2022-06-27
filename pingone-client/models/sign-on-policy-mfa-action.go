package models

type MultiFactorAuthenticationActionApplicationDeviceAuthorization struct {
	Enabled           *bool   `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	ExtraVerification *string `json:"extraVerification,omitempty" mapstructure:"extra_verification,omitempty"`
}

type MultiFactorAuthenticationActionApplication struct {
	ID                  *string                                                        `json:"id,omitempty" mapstructure:"id,omitempty"`
	AutoEnrollment      map[string]bool                                                `json:"autoEnrollment,omitempty" mapstructure:"auto_enrollment,omitempty"`
	DeviceAuthorization *MultiFactorAuthenticationActionApplicationDeviceAuthorization `json:"deviceAuthorization,omitempty" mapstructure:"device_authorization,omitempty"`
}

type MultiFactorAuthenticationAction struct {
	Authenticator   map[string]bool                              `json:"authenticator,omitempty" mapstructure:"authenticator,omitempty"`
	BoundBiometrics map[string]bool                              `json:"boundBiometrics,omitempty" mapstructure:"bound_biometrics,omitempty"`
	Email           map[string]bool                              `json:"email,omitempty" mapstructure:"email,omitempty"`
	SecurityKey     map[string]bool                              `json:"securityKey,omitempty" mapstructure:"security_key,omitempty"`
	Sms             map[string]bool                              `json:"sms,omitempty" mapstructure:"sms,omitempty"`
	Voice           map[string]bool                              `json:"voice,omitempty" mapstructure:"voice,omitempty"`
	Applications    []MultiFactorAuthenticationActionApplication `json:"applications,omitempty" mapstructure:"applications,omitempty"`
	NoDeviceMode    *string                                      `json:"noDeviceMode,omitempty" mapstructure:"no_device_mode,omitempty"`
}
