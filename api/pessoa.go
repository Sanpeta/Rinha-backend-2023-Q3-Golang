package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	db "github.com/Sanpeta/rinha-backend-2023-q3-golang/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createPessoaRequest struct {
	Name      string   `json:"nome" binding:"required,max=32"`
	Nickname  string   `json:"apelido" binding:"required,max=100"`
	Birthdate string   `json:"nascimento" binding:"required,min=10,max=10"`
	Stack     []string `json:"stack" binding:"omitempty,dive,max=32"`
}

func (server *Server) createPessoa(context *gin.Context) {
	var request createPessoaRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !re.MatchString(request.Birthdate) {
		context.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("data de nascimento inválida")))
		return
	}

	uuid := uuid.New()

	id, err := server.store.CreatePessoa(context, db.CreatePessoaParams{
		ID:          uuid,
		Nome:        request.Name,
		Apelido:     request.Nickname,
		Nascimento:  request.Birthdate,
		Stack:       request.Stack,
		SearchIndex: fmt.Sprintf("%s %s %s", request.Name, request.Nickname, strings.ToLower(strings.Join(request.Stack, " "))),
	})
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{})
		return
	}

	context.Header("Location", fmt.Sprintf("/pessoas/%s", id))
	context.JSON(http.StatusCreated, gin.H{})
}

type getPessoaRequest struct {
	ID string `uri:"id" binding:"required"`
}

type getPessoaResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"nome"`
	Nickname  string    `json:"apelido"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

func (server *Server) getPessoa(context *gin.Context) {
	var request getPessoaRequest
	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	uuid, err := uuid.Parse(request.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	pessoa, err := server.store.GetPessoa(context, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{})
			return
		}
		context.JSON(http.StatusNotFound, gin.H{})
		return
	}

	response := getPessoaResponse{
		ID:        pessoa.ID,
		Name:      pessoa.Nome,
		Nickname:  pessoa.Apelido,
		Birthdate: pessoa.Nascimento,
		Stack:     pessoa.Stack,
	}

	context.JSON(http.StatusOK, response)
}

type listPessoasRequest struct {
	Term string `form:"t" binding:"required"`
}

func (server *Server) listPessoas(context *gin.Context) {
	var request listPessoasRequest
	err := context.ShouldBindQuery(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	pessoas, err := server.store.GetPessoas(context, sql.NullString{String: request.Term, Valid: true})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	context.JSON(http.StatusOK, pessoas)
}

func (server *Server) contagemPessoas(context *gin.Context) {
	count, err := server.store.CountPessoas(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	context.String(http.StatusOK, fmt.Sprintf("%d", count))
}
