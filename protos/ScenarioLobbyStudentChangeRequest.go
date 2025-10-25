package protos

type ScenarioLobbyStudentChangeRequest struct {
	RequestPacket
	Protocol            Protocol
	LobbyStudents       []int64
	LobbyStudentsBefore []int64
}
