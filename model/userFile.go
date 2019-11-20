package model

import "time"

//UserFile UserFile
type UserFile struct {
	BaseModel

	UserID     int
	FileHash   string
	FileSize   int
	FileName   string
	UploadAt   time.Time
	LastUpdate time.Time
	status     int
}
