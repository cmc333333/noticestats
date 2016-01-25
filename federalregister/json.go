package federalregister

type NoticeResult struct {
	DocumentNumber string   `json:"document_number"`
	Published      string   `json:"publication_date"`
	Agencies       []string `json:"agency_names"`
}

type NoticeResults struct {
	Count   int
	Results []NoticeResult
}
