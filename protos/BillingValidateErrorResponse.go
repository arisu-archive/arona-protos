package protos

type BillingValidateErrorResponse struct {
	Error_code        int64 `json:",omitempty,omitzero"`
	Error             string
	Error_description string
}
