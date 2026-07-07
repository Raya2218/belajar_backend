package kecamatan_v010

import (
	"rsudlampung/modules/kabkot_v010"
	"time"
)

type Kecamatan struct {
	ID        uint64             `gorm:"primaryKey" json:"id"`
	Name      string             `gorm:"type:varchar(255)" json:"name"`
	KabkotID  uint64             `json:"kabkot_id"`
	Kabkot    kabkot_v010.Kabkot `gorm:"foreignKey:KabkotID" json:"kabkot"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Status    bool               `json:"status"`
	Fasprof   bool               `json:"fasprof" gorm:"type:bool;default:false"`
	IN        bool               `json:"in" gorm:"type:bool;default:false"`
}

type KecamatanCreate struct {
	Name     string `json:"name" binding:"required"`
	KabkotID uint64 `json:"kabkot_id" binding:"required"`
	Status   bool   `json:"status"`
}

type KecamatanEdit struct {
	ID       uint64 `json:"id" binding:"required"`
	Name     string `json:"name"`
	KabkotID uint64 `json:"kabkot_id"`
	Fasprof  bool   `json:"fasprof"`
	IN       bool   `json:"in"`
	Status   bool   `json:"status"`
}
