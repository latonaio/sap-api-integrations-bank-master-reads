package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bank-master-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBank(raw []byte, l *logger.Logger) ([]Bank, error) {
	pm := &responses.Bank{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Bank. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	bank := make([]Bank, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		bank = append(bank, Bank{
			BankCountry:         data.BankCountry,
			BankInternalID:      data.BankInternalID,
			BankName:            data.BankName,
			Region:              data.Region,
			ShortStreetName:     data.ShortStreetName,
			ShortCityName:       data.ShortCityName,
			SWIFTCode:           data.SWIFTCode,
			BankNetworkGrouping: data.BankNetworkGrouping,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
			Bank:                data.Bank,
			BankBranch:          data.BankBranch,
			BankCategory:        data.BankCategory,
		})
	}

	return bank, nil
}
