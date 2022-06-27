package models

type ProgressiveProfilingAction struct {
	PreventMultiplePromptsPerFlow *bool                    `json:"preventMultiplePromptsPerFlow,omitempty" mapstructure:"prevent_multiple_prompts_per_flow,omitempty"`
	PromptIntervalSeconds         *int                     `json:"promptIntervalSeconds,omitempty" mapstructure:"prompt_interval_seconds,omitempty"`
	PromptText                    *string                  `json:"promptText,omitempty" mapstructure:"prompt_text,omitempty"`
	Attributes                    []map[string]interface{} `json:"attributes,omitempty" mapstructure:"attributes,omitempty"`
}
