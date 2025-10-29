package protos

type AccountAuthRequest struct {
	RequestPacket
	Version int64 `json:",omitempty,omitzero"`
	DevId string `json:",omitempty,omitzero"`
	IMEI int64 `json:",omitempty,omitzero"`
	AccessIP string `json:",omitempty,omitzero"`
	MarketId string `json:",omitempty,omitzero"`
	UserType string `json:",omitempty,omitzero"`
	AdvertisementId string `json:",omitempty,omitzero"`
	OSType string `json:",omitempty,omitzero"`
	OSVersion string `json:",omitempty,omitzero"`
	DeviceUniqueId string `json:",omitempty,omitzero"`
	DeviceModel string `json:",omitempty,omitzero"`
	DeviceSystemMemorySize int32 `json:",omitempty,omitzero"`
	CountryCode string `json:",omitempty,omitzero"`
	Idfv string `json:",omitempty,omitzero"`
	IsTeenVersion bool `json:",omitempty,omitzero"`
	DeviceLocaleCode string `json:",omitempty,omitzero"`
	GameOptionLanguage string `json:",omitempty,omitzero"`
}
