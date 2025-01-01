package helpers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	 
}