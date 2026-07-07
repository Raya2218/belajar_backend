package modules

import (
	"rsudlampung/helper"
	groupv010 "rsudlampung/modules/group_v010"
	groupv011 "rsudlampung/modules/group_v011"

	userv010 "rsudlampung/modules/user_v010"
	userv011 "rsudlampung/modules/user_v011"

	kabkotv010 "rsudlampung/modules/kabkot_v010"
	provinsiv010 "rsudlampung/modules/provinsi_v010"
	provinsiv011 "rsudlampung/modules/provinsi_v011"

	desav010 "rsudlampung/modules/desa_v010"
	kecamatanv010 "rsudlampung/modules/kecamatan_v010"

	"github.com/gin-gonic/gin"
)

type Versions interface {
	Run()
}

type versions struct {
	configEnv  helper.Config
	mainServer *gin.Engine
}

func NewVersion(configEnv helper.Config, mainServer *gin.Engine) Versions {
	return &versions{
		configEnv:  configEnv,
		mainServer: mainServer,
	}
}

func (s *versions) Run() {
	db_010 := helper.OpenDB(s.configEnv.DB, s.configEnv.SCHEMA, "v010")
	db_011 := helper.OpenDB(s.configEnv.DB, s.configEnv.SCHEMA, "v011")

	registerRoutes := func(apiRoutes *gin.RouterGroup) {
		groupV010 := groupv010.NewGroupServer(apiRoutes, db_010, "v010")
		groupV010.Init()

		groupV011 := groupv011.NewGroupServer(apiRoutes, db_011, "v011")
		groupV011.Init()

		userV010 := userv010.NewUserServer(apiRoutes, db_010, "v010")
		userV010.Init()

		userV011 := userv011.NewUserServer(apiRoutes, db_011, "v011")
		userV011.Init()

		provinsiV010 := provinsiv010.NewProvinsiServer(apiRoutes, db_010, "v010")
		provinsiV010.Init()

		provinsiV011 := provinsiv011.NewProvinsiServer(apiRoutes, db_011, "v011")
		provinsiV011.Init()

		kabkotV010 := kabkotv010.NewKabkotServer(apiRoutes, db_010, "v010")
		kabkotV010.Init()

		kecamatanV010 := kecamatanv010.NewKecamatanServer(apiRoutes, db_010, "v010")
		kecamatanV010.Init()

		desaV010 := desav010.NewDesaServer(apiRoutes, db_010, "v010")
		desaV010.Init()
	}

	registerRoutes(s.mainServer.Group("/"))
	registerRoutes(s.mainServer.Group("/sistem"))

}
