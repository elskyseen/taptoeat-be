package controllers

import (
	"net/http"
	"os"
	"strings"
	"taptoeat-be/models"
	"taptoeat-be/validations"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	// Validate username min 3 character
	if checkValidCharacter := validations.IsValidChar(replaceUsername, 3, 12); !checkValidCharacter {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username min 3 character or less than 12 character",
			"code":    http.StatusBadRequest,
			"len":     len(replaceUsername),
		})
		return
	}

	// Check username can't include simbol
	if usernameIsValid := validations.IsValidUsername(replaceUsername); usernameIsValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username must be valid",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Check email is valid
	if emailIsValid := validations.IsValidEmail(body.Email); !emailIsValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email must be valid",
			"code":    http.StatusBadRequest,
			"is":      emailIsValid,
		})
		return
	}

	// Validate email exist
	if checkExistEmail := validations.IsExistField(body.Email, user.Email); checkExistEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist, please using other email",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Validate username exist
	if checkExistUsername := validations.IsExistField(body.Email, user.Email); checkExistUsername {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username already exist, please using other name",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// Check length from passwordd
	if isValidPassword := validations.IsValidChar(body.Password, 7, 16); !isValidPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords can't be less than 8 and must not be more than 16",
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

func SignIn(c *gin.Context) {
	var user models.User
	var body struct {
		Email    string
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	replaceUsername := strings.ToLower(strings.Replace(body.Username, " ", "", -1))

	result := models.DB.Where("username = ?", replaceUsername).Or("email = ?", body.Email).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Login Failed",
			"code":    http.StatusBadRequest,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}
	c.SetCookie("Authorization", tokenString, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success sigin",
		"code":    http.StatusOK,
	})
}
