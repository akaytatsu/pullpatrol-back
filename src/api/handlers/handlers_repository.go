package handlers

import (
	"app/entity"
	"app/infrastructure/git/github"
	"app/infrastructure/repository"
	usecase_repository "app/usecase/repository"
	"database/sql"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RepositoryHandlers struct {
	UsecaseRepository usecase_repository.IUsecaseRepository
}

func NewRepositoryHandler(usecaseRepo usecase_repository.IUsecaseRepository) *RepositoryHandlers {
	return &RepositoryHandlers{UsecaseRepository: usecaseRepo}
}

func (h RepositoryHandlers) GetRepositoriesHandle(c *gin.Context) {
	repositories, err := h.UsecaseRepository.GetRepositories()
	if handleError(c, err) {
		return
	}
	jsonResponse(c, http.StatusOK, repositories)
}

func (h RepositoryHandlers) CreateRepositoryHandle(c *gin.Context) {
	var repository entity.EntityRepository
	if err := c.ShouldBindJSON(&repository); handleError(c, err) {
		return
	}
	if err := h.UsecaseRepository.Create(&repository); handleError(c, err) {
		return
	}
	jsonResponse(c, http.StatusOK, repository)
}

func (h RepositoryHandlers) GetRepositoryHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repository, err := h.UsecaseRepository.Get(id)

	if handleError(c, err) {
		return
	}

	jsonResponse(c, http.StatusOK, repository)
}

func (h RepositoryHandlers) UpdateRepositoryHandle(c *gin.Context) {
	var repository entity.EntityRepository

	if err := c.ShouldBindJSON(&repository); handleError(c, err) {
		return
	}

	if err := h.UsecaseRepository.Update(&repository); handleError(c, err) {
		return
	}

	jsonResponse(c, http.StatusOK, repository)
}

func (h RepositoryHandlers) DeleteRepositoryHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := h.UsecaseRepository.Get(id)
	if handleError(c, err) {
		return
	}
	if err := h.UsecaseRepository.Delete(repo); handleError(c, err) {
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, nil)
}

func (h RepositoryHandlers) GitWebhookHandler(c *gin.Context) {
	var github github.Github

	jsonData, _ := io.ReadAll(c.Request.Body)

	err := h.UsecaseRepository.ProccessPullRequest(github, jsonData)

	if handleError(c, err) {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "ok"})
}

func MountRepositoryHandlers(r *gin.Engine, conn *sql.DB) {

	repoHandlers := NewRepositoryHandler(
		usecase_repository.NewService(
			repository.NewRepositoryRepository(conn),
		),
	)

	repoGroup := r.Group("/api/repository")
	repoGroup.GET("", repoHandlers.GetRepositoriesHandle)
	repoGroup.POST("", repoHandlers.CreateRepositoryHandle)
	repoGroup.GET("/:id", repoHandlers.GetRepositoryHandle)
	repoGroup.PUT("/:id", repoHandlers.UpdateRepositoryHandle)
	repoGroup.DELETE("/:id", repoHandlers.DeleteRepositoryHandle)
	r.POST("/git-webhook", repoHandlers.GitWebhookHandler)
}
