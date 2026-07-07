package desa_v010

import (
	"rsudlampung/modules/kecamatan_v010"
	"time"
)

type Desa struct {
	ID          uint64                   `gorm:"primaryKey" json:"id"`
	Name        string                   `gorm:"type:varchar(255)" json:"name"`
	KecamatanID uint64                   `json:"kecamatan_id"`
	Kecamatan   kecamatan_v010.Kecamatan `gorm:"foreignKey:KecamatanID" json:"kecamatan"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	Status      bool                     `json:"status"`
	Fasprof     bool                     `json:"fasprof" gorm:"type:bool;default:false"`
	IN          bool                     `json:"in" gorm:"type:bool;default:false"`
}

type DesaCreate struct {
	Name        string `json:"name" binding:"required"`
	KecamatanID uint64 `json:"kecamatan_id" binding:"required"`
	Status      bool   `json:"status"`
}

type DesaEdit struct {
	ID          uint64 `json:"id" binding:"required"`
	Name        string `json:"name"`
	KecamatanID uint64 `json:"kecamatan_id"`
	Fasprof     bool   `json:"fasprof"`
	IN          bool   `json:"in"`
	Status      bool   `json:"status"`
}
