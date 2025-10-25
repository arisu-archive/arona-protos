package protos

type TTSGetFileResponse struct {
	ResponsePacket
	Protocol     Protocol
	IsFileReady  bool
	TTSFileS3Uri string
}
