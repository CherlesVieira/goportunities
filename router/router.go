package router

import "github.com/gin-gonic/gin" //quando nao der o nome na frente ser o nome do ultimo diretorio (gin)

func Initialize() {

	//Initialize Router
	router := gin.Default() //definindo e declarando a variavel ao mesmo tempo

	//Inicialize Routes
	initializeRoutes(router)

	//Run Server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
