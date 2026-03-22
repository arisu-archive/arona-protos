package protos

type AccountAuthRequest struct {
	RequestPacket
	Version                int64 `json:",omitempty,omitzero"`
	DevId                  string
	IMEI                   int64 `json:",omitempty,omitzero"`
	AccessIP               string
	MarketId               string
	UserType               string
	AdvertisementId        string
	OSType                 string
	OSVersion              string
	DeviceUniqueId         string
	DeviceModel            string
	DeviceSystemMemorySize int32 `json:",omitempty,omitzero"`
	CountryCode            string
	Idfv                   string
	IsTeenVersion          bool `json:",omitempty,omitzero"`
	DeviceLocaleCode       string
	GameOptionLanguage     string
}
