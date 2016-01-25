package federalregister

type NoticeResult struct {
	DocumentNumber string `json:"document_number"`
	Published      string `json:"publication_date"`
}

type NoticeResults struct {
	Count   int
	Results []NoticeResult
}
