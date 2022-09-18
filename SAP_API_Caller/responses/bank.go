package responses

type Bank struct {
	OdataContext      string `json:"@odata.context"`
	OdataMetadataEtag string `json:"@odata.metadataEtag"`
	Value             []struct {
		BankCountry         string `json:"BankCountry"`
		BankInternalID      string `json:"BankInternalID"`
		BankName            string `json:"BankName"`
		Region              string `json:"Region"`
		ShortStreetName     string `json:"ShortStreetName"`
		ShortCityName       string `json:"ShortCityName"`
		SWIFTCode           string `json:"SWIFTCode"`
		BankNetworkGrouping string `json:"BankNetworkGrouping"`
		IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
		Bank                string `json:"Bank"`
		BankBranch          string `json:"BankBranch"`
		BankCategory        string `json:"BankCategory"`
	} `json:"value"`
}
