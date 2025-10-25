package protos

type AccountLoginSyncRequest struct {
	RequestPacket
	Protocol         Protocol
	SyncProtocols    []Protocol
	SkillCutInOption string
}
