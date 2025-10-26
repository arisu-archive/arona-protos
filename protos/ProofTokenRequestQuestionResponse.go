package protos

type ProofTokenRequestQuestionResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
