package protos

type AccountCheckYostarRequest struct {
	RequestPacket
	UID int64 `json:",omitempty,omitzero"`
	YostarToken string `json:",omitempty,omitzero"`
	EnterTicket string `json:",omitempty,omitzero"`
	PassCookieResult bool `json:",omitempty,omitzero"`
	Cookie string `json:",omitempty,omitzero"`
}
