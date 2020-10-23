package apis

import (
	"net/http"
	"chaca/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

func DocenteGetId(c *gin.Context) {
	id := c.Params.ByName("id")
	var cur models.Docente
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	if err := conn.First(&cur, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &cur)
}




func DocenteGet(c *gin.Context) {
	var lis []models.Docente
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}



func DocentePost(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	cur := models.Docente{Name: c.PostForm("name"), Paternal: c.PostForm("paternal"), Maternal: c.PostForm("maternal"), Age:c.PostForm("age")}
	conn.Create(&cur)
	c.JSON(http.StatusOK, &cur)
}

func DocentePut(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var cur models.Docente
	if err := conn.First(&cur, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	cur.Name = c.PostForm("name")
	cur.Age = c.PostForm("age")
	conn.Save(&cur)
	c.JSON(http.StatusOK, &cur)
}

func DocenteDelete(c *gin.Context){
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var cur models.Docente
	if err := conn.Where("id = ?", id).First(&cur).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&cur)
}