package protos

type ManagementBannerListResponse struct {
	ResponsePacket
	Protocol  Protocol
	BannerDBs []BannerDB
}
