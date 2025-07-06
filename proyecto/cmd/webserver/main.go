package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear una nueva instancia de Gin
	r := gin.Default()

	// Ruta básica
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test server funcionando !",
		})
	})

	// Ruta para obtener información del proyecto
	r.GET("/api/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"proyecto":  "test server",
			"framework": "Gin",
			"version":   "1.10.1",
		})
	})

	// Ruta con parámetro
	r.GET("/api/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Hola " + nombre + "!",
		})
	})

	// Ruta POST de ejemplo
	r.POST("/api/datos", func(c *gin.Context) {
		var datos map[string]interface{}
		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Datos JSON inválidos",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Datos recibidos correctamente",
			"datos":   datos,
		})
	})

	// Servir archivos estáticos desde la carpeta assets
	r.Static("/assets", "./assets")

	// Servir archivos HTML desde la carpeta templates
	r.LoadHTMLGlob("templates/*")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "test server",
		})
	})

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
