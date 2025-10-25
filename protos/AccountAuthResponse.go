package protos

type AccountAuthResponse struct {
	ResponsePacket
	Protocol                            Protocol
	CurrentVersion                      int64
	MinimumVersion                      int64
	IsDevelopment                       bool
	BattleValidation                    bool
	UpdateRequired                      bool
	TTSCdnUri                           string
	AccountDB                           AccountDB
	AttendanceBookRewards               []AttendanceBookReward
	AttendanceHistoryDBs                []AttendanceHistoryDB
	RepurchasableMonthlyProductCountDBs []PurchaseCountDB
	MonthlyProductParcel                []ParcelInfo
	MonthlyProductMail                  []ParcelInfo
	BiweeklyProductParcel               []ParcelInfo
	BiweeklyProductMail                 []ParcelInfo
	WeeklyProductParcel                 []ParcelInfo
	WeeklyProductMail                   []ParcelInfo
	EncryptedUID                        string
	AccountRestrictionsDB               AccountRestrictionsDB
	IssueAlertInfos                     []IssueAlertInfoDB
	AccountBanByNexonDBs                []AccountBanByNexonDB
}
