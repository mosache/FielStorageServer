package model

import "time"

//UserFile UserFile
type UserFile struct {
	BaseModel

	UserID     int64
	FileHash   string
	FileSize   int64
	FileName   string
	UploadAt   time.Time
	LastUpdate time.Time
	status     int
}
