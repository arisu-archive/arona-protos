package protos

type BillingValidateNexonResponse struct {
	Stamp_token     string   `json:",omitempty,omitzero"`
	Order_id        string   `json:",omitempty,omitzero"`
	Product_id      []string `json:",omitempty,omitzero"`
	User_id         string   `json:",omitempty,omitzero"`
	Character_id    string   `json:",omitempty,omitzero"`
	Service_payload string   `json:",omitempty,omitzero"`
	Market_type     string   `json:",omitempty,omitzero"`
	Purchase_type   string   `json:",omitempty,omitzero"`
	Currency        string   `json:",omitempty,omitzero"`
	Price           string   `json:",omitempty,omitzero"`
}
