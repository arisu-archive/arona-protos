package protos

type AccountAuthRequest struct {
	RequestPacket
	Protocol               Protocol
	Version                int64
	DevId                  string
	IMEI                   int64
	AccessIP               string
	MarketId               string
	UserType               string
	AdvertisementId        string
	OSType                 string
	OSVersion              string
	DeviceUniqueId         string
	DeviceModel            string
	DeviceSystemMemorySize int32
	CountryCode            string
	Idfv                   string
	IsTeenVersion          bool
	DeviceLocaleCode       string
	GameOptionLanguage     string
}
