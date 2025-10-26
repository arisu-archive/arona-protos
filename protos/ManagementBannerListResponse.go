package protos

type ManagementBannerListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BannerDBs []BannerDB `json:",omitempty,omitzero"`
}
