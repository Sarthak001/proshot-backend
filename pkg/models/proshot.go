package models

import "gorm.io/gorm"

type UserDetails struct {
	gorm.Model
	Username     string        `json:"userName" gorm:"type:VARCHAR(255);uniqueIndex;NOT NULL"`
	Firstname    string        `json:"firstName" gorm:"type:VARCHAR(255)"`
	Lastname     string        `json:"lastName" gorm:"type:VARCHAR(255)"`
	Email        string        `json:"email" gorm:"type:VARCHAR(255);uniqueIndex;NOT NULL"`
	Password     string        `json:"password" gorm:"type:VARCHAR(255);NOT NULL"`
	Secret       string        `json:"secret" gorm:"type:VARCHAR(255)"`
	Isenabled    bool          `json:"isEnabled" gorm:"default:false"`
	Isverified   bool          `json:"isVerified" gorm:"default:false"`
	Albums       []Album       `gorm:"foreignKey:Owner;references:Username"`
	Photos       []Photo       `gorm:"foreignKey:Owner;references:Username"`
	SharedAlbums []SharedAlbum `gorm:"foreignKey:Owner;references:Username"`
}

type Album struct {
	gorm.Model
	Name            string        `json:"name" gorm:"type:VARCHAR(255);uniqueIndex;NOT NULL"`
	Path            string        `json:"path" gorm:"type:VARCHAR(255)"`
	Owner           string        `json:"owner" gorm:"type:VARCHAR(255);"`
	Description     string        `json:"description" gorm:"type:VARCHAR(255);"`
	Sharedto        string        `json:"sharedTo" gorm:"type:VARCHAR(255);"`
	Photos          []Photo       `gorm:"foreignKey:Album;references:Name"`
	Sharedalbumname []SharedAlbum `gorm:"foreignKey:Albumname;references:Name"`
}

type Photo struct {
	gorm.Model
	Name  string `json:"firstName" gorm:"type:VARCHAR(255)"`
	Path  string `json:"path" gorm:"type:VARCHAR(255)"`
	Size  int64  `json:"size" gorm:"type:INT"`
	Owner string `json:"Owner" gorm:"type:VARCHAR(255);"`
	Album string `json:"album" gorm:"type:VARCHAR(255)"`
}

type SharedAlbum struct {
	gorm.Model
	Urlpath     string `json:"urlPath" gorm:"type:VARCHAR(255);"`
	Owner       string `json:"owner" gorm:"type:VARCHAR(255);"`
	Password    string `json:"password" gorm:"type:VARCHAR(255);"`
	Isexpirable bool   `json:"isExpirable" gorm:"default:true"`
	Albumname   string `json:"albumName" gorm:"type:VARCHAR(255);"`
}
