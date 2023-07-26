package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_user "app/usecase/user"
	"database/sql"
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

func (h UserHandlers) GetUsersHandler(c *gin.Context) {

	users, err := h.UsecaseUser.GetUsers()

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, users)
}

func (h UserHandlers) GetUserHandler(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UsecaseUser.GetUser(id)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, user)
}

func (h UserHandlers) GetGroupsHandler(c *gin.Context) {

	groups, err := h.UsecaseUser.GetGroups()

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, groups)
}

func (h UserHandlers) GetUsersByGroupHandler(c *gin.Context) {

	groupID, _ := strconv.Atoi(c.Param("groupID"))

	users, err := h.UsecaseUser.GetUsersByGroup(groupID)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, users)
}

func (h UserHandlers) GetGroupHandler(c *gin.Context) {

	groupID, _ := strconv.Atoi(c.Param("groupID"))

	group, err := h.UsecaseUser.GetGroup(groupID)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, group)
}

func (h UserHandlers) CreateGroupHandler(c *gin.Context) {

	var entityGroup entity.EntityGroup

	if err := c.ShouldBindJSON(&entityGroup); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.CreateGroup(&entityGroup)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "Group created successfully"})

}

func (h UserHandlers) UpdateGroupHandler(c *gin.Context) {

	var entityGroup entity.EntityGroup

	if err := c.ShouldBindJSON(&entityGroup); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.UpdateGroup(&entityGroup)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "Group updated successfully"})

}

func (h UserHandlers) DeleteGroupHandler(c *gin.Context) {

	var entityGroup entity.EntityGroup

	if err := c.ShouldBindJSON(&entityGroup); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseUser.DeleteGroup(&entityGroup)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "Group deleted successfully"})

}

func (h UserHandlers) AddUserToGroupHandler(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("userID"))
	groupID, _ := strconv.Atoi(c.Param("groupID"))

	err := h.UsecaseUser.AddUserToGroup(userID, groupID)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User added to group successfully"})

}

func (h UserHandlers) RemoveUserFromGroupHandler(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("userID"))
	groupID, _ := strconv.Atoi(c.Param("groupID"))

	err := h.UsecaseUser.RemoveUserFromGroup(userID, groupID)

	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, gin.H{"message": "User removed from group successfully"})

}

func MountUsersHandlers(gin *gin.Engine, conn *sql.DB) {

	userHandlers := NewUserHandler(
		usecase_user.NewService(
			repository.NewRepositoryUser(conn),
		),
	)

	gin.GET("/", HomeHandler)
	gin.POST("/api/login", userHandlers.LoginHandler)

	gin.POST("/login", userHandlers.LoginHandler)
	gin.GET("/me", userHandlers.GetMeHandler)

	// user
	group := gin.Group("/api/user")
	group.POST("/", userHandlers.CreateUserHandler)
	group.PUT("/", userHandlers.UpdateUserHandler)
	group.DELETE("/", userHandlers.DeleteUserHandler)
	group.PUT("/password/:id", userHandlers.UpdatePasswordHandler)
	group.GET("/", userHandlers.GetUsersHandler)
	group.GET("/:id", userHandlers.GetUserHandler)

	// group
	group = gin.Group("/api/group")
	group.GET("/", userHandlers.GetGroupsHandler)
	group.GET("/:groupID", userHandlers.GetGroupHandler)
	group.GET("/:groupID/users", userHandlers.GetUsersByGroupHandler)
	group.POST("/:groupID/users/:userID", userHandlers.AddUserToGroupHandler)
	group.DELETE("/:groupID/users/:userID", userHandlers.RemoveUserFromGroupHandler)
	group.POST("/", userHandlers.CreateGroupHandler)
	group.PUT("/", userHandlers.UpdateGroupHandler)
	group.DELETE("/", userHandlers.DeleteGroupHandler)
}
