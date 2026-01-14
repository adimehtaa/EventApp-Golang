package main

import (
	"app-event/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) registerUser(c *gin.Context) {
	var register registerRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(register.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	register.Password = string(hashedPassword)

	user := database.Users{
		Email:    register.Email,
		Password: register.Password,
		Name:     register.Name,
	}

	if err := app.models.Users.Insert(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
		"data":    user,
	})
}

func (app *application) loginUser(c *gin.Context) {
	var login loginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (app *application) logoutUser(c *gin.Context) {}
