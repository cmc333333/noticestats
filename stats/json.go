package stats

type NoticeJSON struct {
	FullTextXMLUrl string `json:"full_text_xml_url"`
	PageLength     int    `json:"page_length"`
	CorrectionOf   string `json:"correction_of"`
}
