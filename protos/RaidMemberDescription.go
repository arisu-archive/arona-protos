package protos

type RaidMemberDescription struct {
	AccountId        int64
	AccountName      string
	CharacterId      int64
	DamageCollection RaidDamageCollection
}
