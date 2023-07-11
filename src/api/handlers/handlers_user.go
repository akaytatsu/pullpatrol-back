package handlers

import (
	"app/entity"
	usecase_user "app/usecase/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordData struct {
	Email           string `json:"email"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func LoginHandler(c *gin.Context, usecaseUser usecase_user.IUsecaseUser) {

	var loginData LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := usecaseUser.LoginUser(loginData.Email, loginData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, refreshToken, err := usecase_user.JWTTokenGenerator(*user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})

}

func GetMeHandler(c *gin.Context, usecaseUser usecase_user.IUsecaseUser) {
	user, err := usecaseUser.GetUserByToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUserHandler(c *gin.Context, usecaseUser usecase_user.IUsecaseUser) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := usecaseUser.Create(&entityUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUserHandler(c *gin.Context, usecaseUser usecase_user.IUsecaseUser) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := usecaseUser.Update(&entityUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUserHandler(c *gin.Context, usecaseUser usecase_user.IUsecaseUser) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := usecaseUser.Delete(&entityUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func UpdatePasswordHandler(c *gin.Context, usecaseUser usecase_user.IUsecaseUser) {

	var updatePasswordData UpdateUserPasswordData

	if err := c.ShouldBindJSON(&updatePasswordData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// id to int
	id, _ := strconv.Atoi(c.Param("id"))

	err := usecaseUser.UpdatePassword(id, updatePasswordData.OldPassword, updatePasswordData.NewPassword, updatePasswordData.ConfirmPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
