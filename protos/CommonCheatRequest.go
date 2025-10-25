package protos

type CommonCheatRequest struct {
	RequestPacket
	Protocol              Protocol
	Cheat                 string
	CharacterCustomPreset []CheatCharacterCustomPreset
}
