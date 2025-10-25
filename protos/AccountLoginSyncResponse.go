package protos

type AccountLoginSyncResponse struct {
	ResponsePacket
	Protocol                                 Protocol
	Responses                                ResponsePacket
	CafeGetInfoResponse                      CafeGetInfoResponse
	AccountCurrencySyncResponse              AccountCurrencySyncResponse
	CharacterListResponse                    CharacterListResponse
	EquipmentItemListResponse                EquipmentItemListResponse
	CharacterGearListResponse                CharacterGearListResponse
	ItemListResponse                         ItemListResponse
	EchelonListResponse                      EchelonListResponse
	MemoryLobbyListResponse                  MemoryLobbyListResponse
	CampaignListResponse                     CampaignListResponse
	ArenaLoginResponse                       ArenaLoginResponse
	RaidLoginResponse                        RaidLoginResponse
	EliminateRaidLoginResponse               EliminateRaidLoginResponse
	CraftInfoListResponse                    CraftInfoListResponse
	ClanLoginResponse                        ClanLoginResponse
	MomotalkOutlineResponse                  MomoTalkOutLineResponse
	ScenarioListResponse                     ScenarioListResponse
	ShopGachaRecruitListResponse             ShopGachaRecruitListResponse
	TimeAttackDungeonLoginResponse           TimeAttackDungeonLoginResponse
	EventContentPermanentListResponse        EventContentPermanentListResponse
	AttachmentGetResponse                    AttachmentGetResponse
	BillingPurchaseListByNexonResponse       BillingPurchaseListByNexonResponse
	AttachmentEmblemListResponse             AttachmentEmblemListResponse
	ContentSweepMultiSweepPresetListResponse ContentSweepMultiSweepPresetListResponse
	StickerListResponse                      StickerLoginResponse
	MultiFloorRaidSyncResponse               MultiFloorRaidSyncResponse
	FriendCount                              int64
	FriendCode                               string
	PickupFirstGetHistoryDBs                 []PickupFirstGetHistoryDB
	AccountLevelRewardIds                    []int64
}
