package protos

type BillingPurchaseFreeProductResponse struct {
	ResponsePacket
	Protocol        Protocol
	ParcelResult    ParcelResultDB
	MailDB          MailDB
	PurchaseProduct PurchaseCountDB
}
