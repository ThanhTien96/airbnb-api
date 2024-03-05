package common

type FilterParams struct {
	Limit     string   `json:"limit" query:"limit"`
	Page      string   `json:"page" query:"page"`
	StartDate string   `json:"start_date" query:"start_date"`
	EndDate   string   `json:"end_date" query:"end_date"`
	Sorts     []string `json:"sorts" query:"sorts"`
	Filters   string   `json:"filters" query:"filters"`
	Search    string   `json:"search" query:"search"`
}
