package protos

type ProofTokenSubmitRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
