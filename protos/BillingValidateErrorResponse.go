package protos

type BillingValidateErrorResponse struct {
	Error_code int64 `json:",omitempty,omitzero"`
	Error string `json:",omitempty,omitzero"`
	Error_description string `json:",omitempty,omitzero"`
}
