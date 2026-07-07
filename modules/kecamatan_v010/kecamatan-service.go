package kecamatan_v010

import (
	"log"
	"rsudlampung/helper"
	"time"

	"gorm.io/gorm"
)

type KecamatanService interface {
	Create(Kecamatan) (Kecamatan, error)
	Update(Kecamatan) error
	Delete(Kecamatan) error
	FindAll() []Kecamatan
	FindById(kecamatanId uint64) Kecamatan
	FindByKabkot(kabkotId uint64) []Kecamatan
}

type kecamatanService struct {
	conn *gorm.DB
}

func NewKecamatanService(db *gorm.DB) KecamatanService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		db.AutoMigrate(&Kecamatan{})
	}

	return &kecamatanService{
		conn: db,
	}
}

func (service *kecamatanService) Create(kecamatan Kecamatan) (Kecamatan, error) {
	kecamatan.CreatedAt = time.Now()
	err := service.conn.Create(&kecamatan).Error
	if err != nil {
		return Kecamatan{}, err
	}
	return kecamatan, nil
}

func (service *kecamatanService) Update(kecamatan Kecamatan) error {
	kecamatan.UpdatedAt = time.Now()
	err := service.conn.Save(&kecamatan).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *kecamatanService) Delete(kecamatan Kecamatan) error {
	err := service.conn.Delete(&kecamatan).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *kecamatanService) FindAll() []Kecamatan {
	var kecamatans []Kecamatan
	service.conn.Preload("Kabkot.Provinsi").Find(&kecamatans)
	return kecamatans
}

func (service *kecamatanService) FindById(kecamatanId uint64) Kecamatan {
	var kecamatan Kecamatan
	service.conn.Preload("Kabkot.Provinsi").Where("id=?", kecamatanId).Find(&kecamatan)
	return kecamatan
}

func (service *kecamatanService) FindByKabkot(Id uint64) []Kecamatan {
	var kecamatans []Kecamatan
	service.conn.Preload("Kabkot.Provinsi").Where("kabkot_id=?").Find(&kecamatans)
	return kecamatans
}
