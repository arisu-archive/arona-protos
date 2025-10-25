package protos

type ArenaUserDB struct {
	AccountServerId             int64
	RepresentCharacterUniqueId  int64
	RepresentCharacterCostumeId int64
	NickName                    string
	Rank                        int64
	Level                       int64
	Exp                         int64
	TeamSettingDB               ArenaTeamSettingDB
	AccountAttachmentDB         AccountAttachmentDB
	UserName                    string
}
