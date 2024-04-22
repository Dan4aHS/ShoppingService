package fns_api

type FNSResponse struct {
	Code    int             `json:"code"`
	First   int             `json:"first"`
	Data    DataResponse    `json:"data"`
	Request RequestResponse `json:"request"`
}

type DataResponse struct {
	Json JsonResponse `json:"json"`
	Html string       `json:"html"`
}

type JsonResponse struct {
	Code                      int                `json:"code"`
	User                      string             `json:"user"`
	Items                     []ItemResponse     `json:"items"`
	Nds10                     int                `json:"nds10"`
	Nds18                     int                `json:"nds18"`
	FnsUrl                    string             `json:"fnsUrl"`
	Region                    string             `json:"region"`
	UserInn                   string             `json:"userInn"`
	DateTime                  string             `json:"dateTime"`
	KktRegId                  string             `json:"kktRegId"`
	Metadata                  MetadataResponse   `json:"metadata"`
	Operator                  string             `json:"operator"`
	TotalSum                  int                `json:"totalSum"`
	CreditSum                 int                `json:"creditSum"`
	NumberKkt                 string             `json:"numberKkt"`
	FiscalSign                int                `json:"fiscalSign"`
	PrepaidSum                int                `json:"prepaidSum"`
	Properties                PropertiesResponse `json:"properties"`
	OperatorInn               string             `json:"operatorInn"`
	RetailPlace               string             `json:"retailPlace"`
	ShiftNumber               int                `json:"shiftNumber"`
	CashTotalSum              int                `json:"cashTotalSum"`
	ProvisionSum              int                `json:"provisionSum"`
	EcashTotalSum             int                `json:"ecashTotalSum"`
	OperationType             int                `json:"operationType"`
	RedefineMask              int                `json:"redefine_mask"`
	RequestNumber             int                `json:"requestNumber"`
	FiscalDriveNumber         string             `json:"fiscalDriveNumber"`
	MessageFiscalSign         float64            `json:"messageFiscalSign"`
	RetailPlaceAddress        string             `json:"retailPlaceAddress"`
	AppliedTaxationType       int                `json:"appliedTaxationType"`
	FiscalDocumentNumber      int                `json:"fiscalDocumentNumber"`
	FiscalDocumentFormatVer   int                `json:"fiscalDocumentFormatVer"`
	CheckingLabeledProdResult int                `json:"checkingLabeledProdResult"`
}

type ItemResponse struct {
	Nds                           int                   `json:"nds"`
	Sum                           int                   `json:"sum"`
	Name                          string                `json:"name"`
	Price                         int                   `json:"price"`
	Quantity                      float32               `json:"quantity"`
	PaymentType                   int                   `json:"paymentType"`
	ProductType                   int                   `json:"productType"`
	ProductCodeNew                ProductCodeResponse   `json:"productCodeNew"`
	LabelCodeProcesMode           int                   `json:"labelCodeProcesMode"`
	ItemsIndustryDetails          ItemsIndustryResponse `json:"itemsIndustryDetails"`
	ItemsQuantityMeasure          int                   `json:"itemsQuantityMeasure"`
	CheckingProdInformationResult int                   `json:"checkingProdInformationResult"`
}

type ProductCodeResponse struct {
	Gs1m  BarCodeResponse `json:"gs1m"`
	Ean13 BarCodeResponse `json:"ean13"`
}

type BarCodeResponse struct {
	Gtin           string `json:"gtin"`
	Sernum         string `json:"sernum"`
	ProductIdType  int    `json:"productIdType"`
	RawProductCode string `json:"rawProductCode"`
}

type ItemsIndustryResponse struct {
	IdFoiv                string `json:"idFoiv"`
	IndustryPropValue     string `json:"industryPropValue"`
	FoundationDocNumber   string `json:"foundationDocNumber"`
	FoundationDocDateTime string `json:"foundationDocDateTime"`
}

type MetadataResponse struct {
	Id          int    `json:"id"`
	OfdId       string `json:"ofdId"`
	Address     string `json:"address"`
	Subtype     string `json:"subtype"`
	ReceiveDate string `json:"receiveDate"`
}

type PropertiesResponse struct {
	PropertyName  string `json:"propertyName"`
	PropertyValue string `json:"propertyValue"`
}

type RequestResponse struct {
	Qrurl  string         `json:"qrurl"`
	Qrfile string         `json:"qrfile"`
	Qrraw  string         `json:"qrraw"`
	Manual ManualResponse `json:"manual"`
}

type ManualResponse struct {
	Fn        string `json:"fn"`
	Fd        string `json:"fd"`
	Fp        string `json:"fp"`
	CheckTime string `json:"check_time"`
	Type      string `json:"type"`
	Sum       string `json:"sum"`
}
