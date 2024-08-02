package databases

import (
	"time"
)

type TempUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Masterkey string `json:"masterkey"`
	Storage   int    `json:"storage"`
}

type User struct {
	UUID      string   `gorm:"type:char(36);primary_key;" json:"uuid"`
	Username  string   `json:"username"`
	Email     string   `gorm:"unique;not null" json:"email"`
	Masterkey string   `json:"masterkey"`
	Folders   []Folder `gorm:"foreignkey:UserID" json:"folders"`
	Files     []File   `gorm:"foreignkey:UserID" json:"files"`
	Storage   Storage  `gorm:"foreignkey:UserID" json:"storage"`
}

type Folder struct {
	UUID     string `gorm:"type:char(36);primary_key;" json:"uuid"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
	UserID   string `json:"user_id"`
	Files    []File `gorm:"foreignkey:FolderID" json:"files"`
}

type File struct {
	UUID      string    `gorm:"type:char(36);primary_key;"`
	Name      string    `json:"name"`
	FolderID  string    `json:"folder_id"`
	UserID    string    `json:"user_id"`
	Size      int       `json:"size"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Checksum  string    `json:"checksum"`
}

type Storage struct {
	UUID       string `gorm:"type:char(36);primary_key;" json:"uuid"`
	UserID     string `json:"user_id"`
	UsedSpace  int    `json:"used_space"`
	TotalSpace int    `json:"total_space"`
	MaxSpace   int    `json:"max_space"`
	BonusSpace int    `json:"bonus_space"`
}
