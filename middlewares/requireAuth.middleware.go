package middlewares

//func RequireAuthMiddleware(c *fiber.Ctx) error {
//	// get token from header
//	token := c.Get("Authorization")
//	// check if token is empty
//	if token == "" {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//			"message": "unauthorized",
//		})
//	}
//
//	// check if token is valid
//	//err := utils.VerifyToken(token)
//	//if err != nil {
//	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//	//		"message": "unauthorized",
//	//	})
//	//}
//	return c.Next()
//
//}
