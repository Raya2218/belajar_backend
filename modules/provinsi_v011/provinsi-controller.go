package provinsi_v011

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProvinsiController interface {
	FindAll() []Provinsi
	FindByID(ctx *gin.Context) (Provinsi, error)
	Create(ctx *gin.Context) (Provinsi, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerProvinsi struct {
	service ProvinsiService
}

func NewProvinsiController(db *gorm.DB) ProvinsiController {

	return &controllerProvinsi{
		service: NewProvinsiService(db),
	}
}

func (c *controllerProvinsi) FindAll() []Provinsi {
	return c.service.FindAll()
}

func (c *controllerProvinsi) FindByID(ctx *gin.Context) (Provinsi, error) {
	var provinsi Provinsi
	id := ctx.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return Provinsi{}, errors.New("Invalid ID format")
	}
	provinsi = c.service.FindByID(idInt)
	if (provinsi == Provinsi{}) {
		return Provinsi{}, errors.New("Provinsi tidak valid")
	}
	return provinsi, nil
}

func (c *controllerProvinsi) Create(ctx *gin.Context) (Provinsi, error) {
	var provinsi Provinsi
	err := ctx.ShouldBindJSON(&provinsi)
	if err != nil {
		return Provinsi{}, err
	}
	result, err := c.service.Create(provinsi)
	if err != nil {
		return Provinsi{}, err
	}
	return result, nil
}

func (c *controllerProvinsi) Update(ctx *gin.Context) error {
	var provinsi Provinsi
	err := ctx.ShouldBindJSON(&provinsi)
	if err != nil {
		return err
	}

	err = c.service.Update(provinsi)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerProvinsi) Delete(ctx *gin.Context) error {
	var provinsi Provinsi
	id := ctx.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("Invalid ID format")
	}

	provinsi = c.service.FindByID(idInt)
	if (provinsi == Provinsi{}) {
		return errors.New("Provinsi tidak valid")
	}

	provinsi.ID = idInt
	err = c.service.Delete(provinsi)
	if err != nil {
		return err
	}
	return nil
}
