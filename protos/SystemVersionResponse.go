package protos

type SystemVersionResponse struct {
	ResponsePacket
	Protocol       Protocol
	CurrentVersion int64
	MinimumVersion int64
	IsDevelopment  bool
}
