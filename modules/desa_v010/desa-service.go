package desa_v010

import (
	"log"
	"time"

	"rsudlampung/helper"

	"gorm.io/gorm"
)

type DesaService interface {
	Create(Desa) (Desa, error)
	Update(Desa) error
	Delete(Desa) error
	FindAll() []Desa
	FindById(desaId uint64) Desa
	FindByKecamatan(kecamatanId uint64) []Desa
}

type desaService struct {
	conn *gorm.DB
}

func NewDesaService(db *gorm.DB) DesaService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		db.AutoMigrate(&Desa{})
	}

	return &desaService{
		conn: db,
	}
}

func (service *desaService) Create(desa Desa) (Desa, error) {
	desa.CreatedAt = time.Now()
	err := service.conn.Create(&desa).Error
	if err != nil {
		return Desa{}, err
	}
	return desa, nil
}

func (service *desaService) Update(desa Desa) error {
	desa.UpdatedAt = time.Now()
	err := service.conn.Save(&desa).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *desaService) Delete(desa Desa) error {
	err := service.conn.Delete(&desa).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *desaService) FindAll() []Desa {
	var desas []Desa
	service.conn.Preload("Kecamatan.Kabkot.Provinsi").Find(&desas)
	return desas
}

func (service *desaService) FindById(desaId uint64) Desa {
	var desa Desa
	service.conn.Preload("Kecamatan.Kabkot.Provinsi").Where("id=?", desaId).Find(&desa)
	return desa
}

func (service *desaService) FindByKecamatan(kecamatanId uint64) []Desa {
	var desas []Desa
	service.conn.Preload("Kecamatan.Kabkot.Provinsi").Where("kecamatan_id=?", kecamatanId).Find(&desas)
	return desas
}
