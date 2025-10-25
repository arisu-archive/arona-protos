package protos

type BillingValidateNexonResponse struct {
	Stamp_token     string
	Order_id        string
	Product_id      []string
	User_id         string
	Character_id    string
	Service_payload string
	Market_type     string
	Purchase_type   string
	Currency        string
	Price           string
}
