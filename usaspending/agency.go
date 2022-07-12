package usaspending

import (
	"encoding/json"
	"fmt"
)

// ------
// Agency
// ------

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
func (usasc *USASpendingClient) GetAgency(id string) (*Agency, error) {
	var agency *Agency
	url := fmt.Sprintf("%s/%s/%s", USASpendingBaseURL, "agency", id)

	resp, err := usasc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget (%d): %w", resp.StatusCode, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&agency)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return agency, nil
}

// -------------
// Agency Awards
// -------------

type AgencyAwards struct {
	FiscalYear       *int          `json:"fiscal_year"`
	LatestActionDate *string       `json:"latest_action_date"`
	ToptierCode      *string       `json:"toptier_code"`
	TransactionCount *int          `json:"transaction_count"`
	Obligations      *float64      `json:"obligations"`
	Messages         []interface{} `json:"messages"`
}

func (usasc *USASpendingClient) GetAgencyAwards(id string) (*AgencyAwards, error) {
	var agencyAwards *AgencyAwards
	url := fmt.Sprintf("%s/%s/%s/%s", USASpendingBaseURL, "agency", id, "awards")

	resp, err := usasc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget (%d): %w", resp.StatusCode, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&agencyAwards)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return agencyAwards, nil
}

// ----------------
// New Awards Count
// ----------------

type NewAwardsCount struct {
	ToptierCode    *string     `json:"toptier_code"`
	FiscalYear     *int        `json:"fiscal_year"`
	AgencyType     *string     `json:"agency_type"`
	AwardTypeCodes interface{} `json:"award_type_codes"`
	NewAwardCount  *int        `json:"new_award_count"`
}

func (usasc *USASpendingClient) GetNewAwardsCount(id string) (*NewAwardsCount, error) {
	var newAwardCount *NewAwardsCount
	url := fmt.Sprintf("%s/%s/%s/%s", USASpendingBaseURL, "agency", id, "awards/new/count")

	resp, err := usasc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget (%d): %w", resp.StatusCode, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&newAwardCount)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return newAwardCount, nil
}

// --------------------
// Budget Function List
// --------------------

type BudgetFunctionList struct {
	TopTierCode  *string                        `json:"toptier_code"`
	FiscalYear   *int                           `json:"fiscal_year"`
	Results      BudgetFunctionListResult       `json:"results"`
	Messages     []interface{}                  `json:"messages"`
	PageMetadata BudgetFunctionListPageMetadata `json:"page_metadata"`
}

type BudgetFunctionListResult struct {
	Name              *string                          `json:"name"`
	Children          BudgetFunctionListResultChildren `json:"children"`
	ObligatedAmount   *float64                         `json:"obligated_amount"`
	GrossOutlayAmount *float64                         `json:"gross_outlay_amount"`
}

type BudgetFunctionListResultChildren struct {
	Name              *string  `json:"name"`
	ObligatedAmount   *float64 `json:"obligated_amount"`
	GrossOutlayAmount *float64 `json:"gross_outlay_amount"`
}

type BudgetFunctionListPageMetadata struct {
	Page        *int        `json:"page"`
	Total       *int        `json:"total"`
	Limit       *int        `json:"limit"`
	Next        interface{} `json:"next"`
	Previous    interface{} `json:"previous"`
	HasNext     *bool       `json:"hasNext"`
	HasPrevious *bool       `json:"hasPrevious"`
}

func (usasc *USASpendingClient) GetBudgetFunctionList(id string) (*BudgetFunctionList, error) {
	var budgetFunctionList *BudgetFunctionList
	url := fmt.Sprintf("%s/%s/%s/%s", USASpendingBaseURL, "agency", id, "budget_function")

	resp, err := usasc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget (%d): %w", resp.StatusCode, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&budgetFunctionList)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return budgetFunctionList, nil
}

// ---------------------
// Budget Function Count
// ---------------------

type BudgetFunctionCount struct {
	TopTierCode            *string       `json:"toptier_code"`
	FiscalYear             *int          `json:"fiscal_year"`
	BudgetFunctionCount    *int          `json:"budget_function_count"`
	BudgetSubFunctionCount *int          `json:"budget_sub_function_count"`
	Messages               []interface{} `json:"messages"`
}

func (usasc *USASpendingClient) GetBudgetFunctionCount(id string) (*BudgetFunctionCount, error) {
	var budgetFunctionCount *BudgetFunctionCount
	url := fmt.Sprintf("%s/%s/%s/%s", USASpendingBaseURL, "agency", id, "budget_function/count")

	resp, err := usasc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget (%d): %w", resp.StatusCode, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&budgetFunctionCount)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return budgetFunctionCount, nil
}

// -------------------
// Budgetary Resources
// -------------------

type BudgetaryResources struct {
	TopTierCode      *string                                  `json:"toptier_code"`
	AgencyDataByYear []BudgetaryResourcesAgencyDataSingleYear `json:"agency_data_by_year"`
	Messages         []interface{}                            `json:"messages"`
}

type BudgetaryResourcesAgencyDataSingleYear struct {
	FiscalYear               *int                                                     `json:"fiscal_year"`
	AgencyBudgetaryResources *float64                                                 `json:"agency_budgetary_resources"`
	AgencyTotalObligated     *float64                                                 `json:"agency_total_obligated"`
	TotalBudgetaryResources  *float64                                                 `json:"total_budgetary_resources"`
	AgencyObligationByPeriod []BudgetaryResourcesAgencyDataSingleYearObligationPeriod `json:"agency_obligation_by_period"`
}

type BudgetaryResourcesAgencyDataSingleYearObligationPeriod struct {
	Period    *int     `json:"period"`
	Obligated *float64 `json:"obligated"`
}

func (usasc *USASpendingClient) GetBudgetaryResources(id string) (*BudgetaryResources, error) {
	var budgetaryResources *BudgetaryResources
	url := fmt.Sprintf("%s/%s/%s/%s", USASpendingBaseURL, "agency", id, "budgetary_resources")

	resp, err := usasc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("httpget (%d): %w", resp.StatusCode, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&budgetaryResources)
	if err != nil {
		return nil, fmt.Errorf("jsondecoder: %w", err)
	}

	return budgetaryResources, nil
}
