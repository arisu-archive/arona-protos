package protos

type AcademyGetInfoResponse struct {
	ResponsePacket
	Protocol           Protocol
	AcademyDB          AcademyDB
	AcademyLocationDBs []AcademyLocationDB
}
