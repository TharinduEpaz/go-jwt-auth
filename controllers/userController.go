package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	database "github.com/TharinduEpaz/go-jwt-auth/database"
	helper "github.com/TharinduEpaz/go-jwt-auth/helpers"
	models "github.com/TharinduEpaz/go-jwt-auth/models"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {}

func Signup() gin.HandlerFunc {}

func Login() gin.HandlerFunc {}

func GetUsers() gin.HandlerFunc {}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // gin.H - header
			return
		}
		
	}
	
}
