package protos

type CharacterListResponse struct {
	ResponsePacket
	Protocol        Protocol
	CharacterDBs    []CharacterDB
	TSSCharacterDBs []CharacterDB
	WeaponDBs       []WeaponDB
	CostumeDBs      []CostumeDB
}
