package kecamatan_v010

import (
	"net/http"

	"rsudlampung/middlewares/mid_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type KecamatanServer interface {
	Init()
}

type kecamatanServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewKecamatanServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) KecamatanServer {

	return &kecamatanServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *kecamatanServer) Init() {

	kecamatanControl := NewKecamatanController(s.database)

	s.apiRoutes.GET("/"+s.version+"/kecamatan/all", func(ctx *gin.Context) {
		ctx.JSON(200, kecamatanControl.FindAll())
	})

	s.apiRoutes.GET("/"+s.version+"/kecamatan/bykabkot/:kabkot_id", func(ctx *gin.Context) {
		ctx.JSON(200, kecamatanControl.FindByKabkot(ctx))
	})

	s.apiRoutes.POST("/"+s.version+"/kecamatan", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := kecamatanControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/kecamatan", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := kecamatanControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}

	})

	s.apiRoutes.DELETE("/"+s.version+"/kecamatan/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := kecamatanControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/kecamatan/ubahstatus/:id/:status", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := kecamatanControl.ChangeStatus(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})
}
