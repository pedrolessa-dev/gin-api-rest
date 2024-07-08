package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrolessa-dev/gin-api-rest/database"
	"github.com/pedrolessa-dev/gin-api-rest/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// RetornaAlunos godoc
// @Summary Retorna todos os alunos
// @Description Retorna uma lista de todos os alunos
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func RetornaAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// RetornaAlunoPeloId godoc
// @Summary Retorna um aluno pelo ID
// @Description Retorna os dados de um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "Aluno ID"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} Resposta
// @Router /alunos/{id} [get]
func RetornaAlunoPeloId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, Resposta{
			Mensagem: "Nenhum aluno encontrado com o id '" + id + "'",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// CadastraAluno godoc
// @Summary Cadastra um novo aluno
// @Description Adiciona um novo aluno ao banco de dados
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param aluno body models.Aluno true "Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} Resposta
// @Router /alunos [post]
func CadastraAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, Resposta{
			Mensagem: err.Error()})
		return
	}
	if err := models.ValidaAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, Resposta{
			Mensagem: err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
// @Summary Deleta um aluno pelo ID
// @Description Remove um aluno do banco de dados pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "Aluno ID"
// @Success 200 {object} Resposta
// @Failure 404 {object} Resposta
// @Router /alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, Resposta{
			Mensagem: "Nenhum aluno encontrado com o id '" + id + "'",
		})
		return
	}
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, Resposta{
		Mensagem: "aluno com id '" + id + "' foi deletado com sucesso",
	})
}

// EditaAluno godoc
// @Summary Edita um aluno pelo ID
// @Description Atualiza os dados de um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "Aluno ID"
// @Param aluno body models.Aluno true "Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} Resposta
// @Failure 404 {object} Resposta
// @Router /alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, Resposta{
			Mensagem: "Nenhum aluno encontrado com o id '" + id + "'",
		})
		return
	}
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, Resposta{
			Mensagem: err.Error(),
		})
		return
	}
	if err := models.ValidaAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, Resposta{
			Mensagem: err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// RetornaAlunoPeloCpf godoc
// @Summary Retorna um aluno pelo CPF
// @Description Retorna os dados de um aluno específico pelo CPF
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param cpf path string true "Aluno CPF"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} Resposta
// @Router /alunos/cpf/{cpf} [get]
func RetornaAlunoPeloCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, Resposta{
			Mensagem: "Nenhum aluno encontrado com o cpf '" + cpf + "'",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func Page404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
