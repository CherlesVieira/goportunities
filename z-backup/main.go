package main

import "github.com/gin-gonic/gin" //quando nao der o nome na frente ser o nome do ultimo diretorio (gin)

func main() {
	router := gin.Default() //definindo e declarando a variavel ao mesmo tempo

	router.GET("/ping",
		func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong - reposta ->>>>"})
		})

	router.GET("/ping2",
		func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong2222222 - reposta"})
		})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
