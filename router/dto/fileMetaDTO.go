package dto

//FileMetaDTO FileMetaDTO
type FileMetaDTO struct {
	FileName string `json:"file_name"`

	FileHash string `json:"file_hash"`

	FileSize int `json:"file_size"`
}
