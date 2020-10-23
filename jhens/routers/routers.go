package routers


import (
	"chaca/apis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"chaca/models"
)


func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database"+err.Error())
	}

	r := gin.Default()
	r.Use(dbMiddleware(*conn))

	cu := r.Group("/cu")
	{
		cu.GET("/cursos/:id", 	apis.CursosGetId)
		cu.GET("/cursos/", 		apis.CursosGet)
		cu.POST("/cursos/", 		apis.CursosPost)
		cu.PUT("/cursos/:id", 	apis.CursosPut)
		cu.DELETE("/cursos/:id", apis.CursosDelete)
	}
	do := r.Group("/do")
	{

		do.GET("/docent/:id", 	apis.DocenteGetId)
		do.GET("/docent/", 		apis.DocenteGet)
		do.POST("/docent/", 		apis.DocentePost)
		do.PUT("/docent/:id", 	apis.DocentePut)
		do.DELETE("/docent/:id", apis.DocenteDelete)
	}

	return r
}

func connectDB() (c *gorm.DB, err error) {

	dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	conn.AutoMigrate(&models.Cursos{})
	conn.AutoMigrate(&models.Docente{})

	if err != nil {
		panic("failed to connect database"+err.Error())
	}
	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

