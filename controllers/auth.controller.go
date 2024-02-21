package controllers

import (
	"net/http"
	"os"
	"strings"
	"taptoeat-be/models"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Declare model
	var user models.User

	// Declate struct from body
	var body struct {
		Username        string
		Email           string
		Password        string
		ConfirmPassword string
	}

	// bind from request body same or no
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Replace whitespace and convert to lowercase
	replaceUsername := strings.ToLower(strings.Replace(body.Username, " ", "", -1))

	// Find username or email from body
	models.DB.Where("username = ?", replaceUsername).Or("email = ?", body.Email).First(&user)

	// Validate email exist
	if user.Email == body.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exists, please using other email",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Validate username exist
	if user.Username == replaceUsername {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username already exists, please using other name",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Check email is valid
	err := checkmail.ValidateFormat(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email must be valid",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Validate password equals to confirm password
	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Password, Try again",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Generate random string for id and password
	genId, genErr := bcrypt.GenerateFromPassword([]byte(os.Getenv("ID")), 1)
	pass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	// Check hashing success or no
	if err != nil || genErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
			"code":    http.StatusInternalServerError,
		})
	}

	// Declare payload after all validate valid
	payload := models.User{
		Id:           "TAP-" + string(genId),
		Username:     replaceUsername,
		Email:        body.Email,
		Password:     string(pass),
		CurrentMoney: 100000,
		CreateAt:     time.Now(),
	}

	// Insert on table user using payload
	result := models.DB.Create(&payload)

	// Check insert proccess, success or fail
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to SignUp",
			"code":    http.StatusInternalServerError,
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Success SignUp",
			"code":    http.StatusCreated,
		})
	}
}


