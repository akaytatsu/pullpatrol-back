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

type UserHandlers struct {
	UsecaseUser usecase_user.IUsecaseUser
}

func NewUserHandler(usecaseUser usecase_user.IUsecaseUser) *UserHandlers {
	return &UserHandlers{UsecaseUser: usecaseUser}
}

func (h UserHandlers) LoginHandler(c *gin.Context) {

	var loginData LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		handleError(c, err)
		return
	}

	user, err := h.UsecaseUser.LoginUser(loginData.Email, loginData.Password)

	if exception := handleError(c, err); exception {
		return
	}

	token, refreshToken, err := usecase_user.JWTTokenGenerator(*user)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})

}

func (h UserHandlers) GetMeHandler(c *gin.Context) {
	user, err := h.UsecaseUser.GetUserByToken(c.GetHeader("Authorization"))

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, user)
}

func (h UserHandlers) CreateUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.Create(&entityUser)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User created successfully"})

}

func (h UserHandlers) UpdateUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.Update(&entityUser)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h UserHandlers) DeleteUserHandler(c *gin.Context) {

	var entityUser entity.EntityUser

	if err := c.ShouldBindJSON(&entityUser); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.Delete(&entityUser)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h UserHandlers) UpdatePasswordHandler(c *gin.Context) {

	var updatePasswordData UpdateUserPasswordData

	if err := c.ShouldBindJSON(&updatePasswordData); err != nil {
		handleError(c, err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := h.UsecaseUser.UpdatePassword(id, updatePasswordData.OldPassword, updatePasswordData.NewPassword, updatePasswordData.ConfirmPassword)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "Password updated successfully"})
}
