package handlers

import (
	"app/entity"
	usecase_repository "app/usecase/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRepositoriesHandle(c *gin.Context, usecaseRepository usecase_repository.IUsecaseRepository) {
	repositories, err := usecaseRepository.GetRepositories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repositories)
}

func CreateRepositoryHandle(c *gin.Context, usecaseRepository usecase_repository.IUsecaseRepository) {
	var repository entity.EntityRepository

	if err := c.ShouldBindJSON(&repository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecaseRepository.Create(&repository); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, repository)
}

func DeleteRepositoryHandle(c *gin.Context, usecaseRepository usecase_repository.IUsecaseRepository) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo, err := usecaseRepository.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := usecaseRepository.Delete(repo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, nil)
}
