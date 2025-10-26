package protos

type ManagementBannerListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
