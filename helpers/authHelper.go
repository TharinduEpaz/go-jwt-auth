package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

/*
CheckUserType:
Takes a Gin context, a role, and a userId as parameters
Retrieves the user's type from the context
Checks if the user's type matches the required role
Returns an error if there's a mismatch ("unauthorized to access this resource")

2. MatchUserTypeToUid:

Takes a Gin context and a userId as parameters
Gets both the user type and uid from the context
Performs two checks:
If the user is a regular "USER", verifies that their uid matches the provided userId

2. Calls CheckUserType to verify the user type matches

Returns an error if any check fails
This is a common pattern in web applications for implementing role-based access control (RBAC) 
and ensuring users can only access resources they're authorized for. For example:
Regular users can only access their own resources (checked by matching UIDs)
Different user types (like "ADMIN" vs "USER") might have different levels of access
*/

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}

	err = CheckUserType(c, userType)

	return err
}
