package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrolessa-dev/gin-api-rest/controllers"
	"github.com/pedrolessa-dev/gin-api-rest/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	docs.SwaggerInfo.BasePath = "/"
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/alunos", controllers.RetornaAlunos)
	r.POST("/alunos", controllers.CadastraAluno)
	r.GET("/alunos/:id", controllers.RetornaAlunoPeloId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.RetornaAlunoPeloCpf)
	r.GET("/", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.Page404)
	r.Run()
}
