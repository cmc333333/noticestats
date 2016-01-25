package stats

type Stats struct {
	Id           string
	XmlLen       int  `db:"xml_len"`
	RegtextLen   int  `db:"regtext_len"`
	PageLen      int  `db:"page_len"`
	IsCorrection bool `db:"is_correction"`
}
