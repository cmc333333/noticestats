package models

type Notice struct {
	DocumentNumber string
}

type Version struct {
	Id        int
	VersionId int `db:"version_id"`
}
