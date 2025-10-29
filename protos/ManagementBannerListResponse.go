package protos

type ManagementBannerListResponse struct {
	ResponsePacket
	BannerDBs []BannerDB `json:",omitempty,omitzero"`
}
