package helpers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func CheckUserType(c *fiber.Ctx, role string) (err error) {
	userType := c.Get("userType")
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this route")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *fiber.Ctx, userId string) (err error) {
	uid := c.Get("uid")
	userType := c.Get("userType")
	err = nil
	// if user type is user and uid is not equal to userId
	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this route")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}
