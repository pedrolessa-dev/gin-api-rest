package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pedrolessa-dev/gin-api-rest/controllers"
	"github.com/pedrolessa-dev/gin-api-rest/database"
	"github.com/pedrolessa-dev/gin-api-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Teste", CPF: "12345678911", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestStatusCodeTodosAlunos(t *testing.T) {
	database.DBConnect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRoutes()
	r.GET("/alunos", controllers.RetornaAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetByCPF(t *testing.T) {
	database.DBConnect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRoutes()
	r.GET("/alunos/cpf/:cpf", controllers.RetornaAlunoPeloCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetByID(t *testing.T) {
	database.DBConnect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRoutes()
	r.GET("/alunos/:id", controllers.RetornaAlunoPeloId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Teste", alunoMock.Nome)
	assert.Equal(t, "12345678911", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeletaAluno(t *testing.T) {
	database.DBConnect()
	CriaAlunoMock()
	r := SetupRoutes()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditaAluno(t *testing.T) {
	database.DBConnect()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRoutes()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Teste do teste", CPF: "12345678901", RG: "123456789"}
	alunoJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunoJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "Teste do teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "12345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456789", alunoMockAtualizado.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}
