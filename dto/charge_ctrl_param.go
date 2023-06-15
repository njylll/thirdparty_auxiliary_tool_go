package dto

type ChargeCtrlParam struct {
	BillCode      string `json:"billCode" binding:"numeric"`
	ChargeType    int    `json:"chargeType" binding:"gte=0, lte=8"`
	CtrlAddr      string `json:"ctrlAddr" binding:"required"`
	DevType       int    `json:"devType" binding:"required"`
	GunCode       int    `json:"gunCode" binding:"required"`
	OperationType int    `json:"operationType" binding:"required"`
}