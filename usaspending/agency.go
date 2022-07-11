package usaspending

import (
	"encoding/json"
	"fmt"
)

type Agency struct {
	FiscalYear                    *int          `json:"fiscal_year"`
	ToptierCode                   *string       `json:"toptier_code"`
	Name                          *string       `json:"name"`
	Abbreviation                  *string       `json:"abbreviation"`
	AgencyID                      *int          `json:"agency_id"`
	IconFilename                  *string       `json:"icon_filename"`
	Mission                       *string       `json:"mission"`
	Website                       *string       `json:"website"`
	CongressionalJustificationURL *string       `json:"congressional_justification_url"`
	AboutAgencyData               interface{}   `json:"about_agency_data"`
	SubtierAgencyCount            *int          `json:"subtier_agency_count"`
	DefCodes                      []DefCode     `json:"def_codes"`
	Messages                      []interface{} `json:"messages"`
}

type DefCode struct {
	Code      *string `json:"code"`
	PublicLaw *string `json:"public_law"`
	Title     *string `json:"title"`
	URLs      *string `json:"urls"`
	Disaster  *string `json:"disaster"`
}

// GetAgency returns the agency with the given ID.
func (ussapi *USASpendingClient) GetAgency(id string) (*Agency, error) {
	var agency *Agency
	url := fmt.Sprintf("%s/%s/%s", USASpendingBaseURL, "agency", id)

	resp, err := ussapi.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget: %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&agency)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return agency, nil
}
