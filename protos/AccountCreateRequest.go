package protos

type AccountCreateRequest struct {
	RequestPacket
	Protocol        Protocol
	DevId           string
	Version         int64
	IMEI            int64
	AccessIP        string
	MarketId        string
	UserType        string
	AdvertisementId string
	OSType          string
	OSVersion       string
	CountryCode     string
}
