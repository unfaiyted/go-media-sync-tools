package controllers

import (
	"media-sync-tools/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body CreateUserInput true "User Data"
// @Success 200 {object} types.User
// @Router /users [post]
func CreateUser(c *gin.Context, db *gorm.DB) {
	// Define a variable to hold input data
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := types.User{Name: input.Name, Email: input.Email, Password: input.Password}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} types.User
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context, db *gorm.DB) {
	user := types.User{}

	db.Where("id = ?", c.Param("id")).Delete(&user)
}

// GetUsers godoc
// @Summary Gets a list of users
// @Description Retrieves a list of users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of users per page"
// @Success 200 {array} types.User
// @Router /users [get]
func GetUsers(c *gin.Context, db *gorm.DB) {
	// Your implementation here
	var users []types.User
	var count int64
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	db.Model(&types.User{}).Count(&count)
	db.Offset(offset).Limit(pageSize).Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users, "total": count, "page": page, "pageSize": pageSize})
}

// GetUserByIDOrUsername godoc
// @Summary Get a single user
// @Description Retrieves a user by their ID or username
// @Tags users
// @Accept json
// @Produce json
// @Param identifier path string true "User ID or Username"
// @Success 200 {object} types.User
// @Router /user/{identifier} [get]
func GetUserByIDOrUsername(c *gin.Context, db *gorm.DB) {
	// Your implementation here
	//
	var user types.User

	identifier := c.Param("identifier")

	if err := db.Where("id = ? OR username = ?", identifier, identifier).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
