package protos

type RecipeCraftRequest struct {
	RequestPacket
	Protocol                 Protocol
	RecipeCraftUniqueId      int64
	RecipeIngredientUniqueId int64
}
