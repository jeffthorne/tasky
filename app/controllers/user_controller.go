package controllers

/*
type LoginForm struct {
	Email    string `json:"email",validate:"required,email,min=6,max=32"`
	Password string `json:"password",validate:"required,min=5,max=32"`
}

type AuthResponse struct {
	AccessToken string      `json:"access_token"`
	Error       bool        `json:"error"`
	Msg         interface{} `json:"msg"`
}

type ErrorResponse struct {
	FailedField string
	NameSpace   string
	Tag         string
	Value       string
	Message     string
}

func ValidateStruct(form any) []interface{} {
	var errors []interface{}

	err := validator.New().Struct(form)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s %s", err.Tag(), err.Type())
			v := map[string]interface{}{err.StructField(): message}
			errors = append(errors, &v)
		}
	}

	return errors
}

// Login
// @Summary Authenticate to pocdb API
// @Description login
// @Tags Auth
// @Accept json
// @Param user body controllers.LoginForm true "User Data"
// @Success 200 {object} controllers.AuthResponse
// @Failure 400,500 {object} object
// @Router /api/v1/login [post]
func Login(c *fiber.Ctx) error {

	loginForm := LoginForm{}
	if err := c.BodyParser(&loginForm); err != nil {
		return c.JSON(map[string]interface{}{"message": err.Error()})
	}

	if err := ValidateStruct(loginForm); err != nil {
		return c.JSON(err)
	}

	user := models.User{}
	if err := database.DB.Where("email = ?", loginForm.Email).First(&user).Error; err != nil {
		return c.JSON(map[string]interface{}{"message": "Invalid Login"})
	}

	if !user.CheckPassword(loginForm.Password) {
		return c.JSON(map[string]interface{}{"message": "Invalid Login"})
	}

	token, err := utils.GenerateJWT(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": true, "msg": err.Error()}) // Return status 500 and token generation error.
	}

	return c.JSON(fiber.Map{"error": false, "msg": nil, "access_token": token})

}

func Test(c *fiber.Ctx) error {
	fmt.Println("IN PROTECTED", c.GetReqHeaders())
	if user, err := utils.UserFromJWT(c); err != nil {
		return c.JSON(map[string]interface{}{"message": "Invalid Login"})
	} else {
		return c.JSON(fiber.Map{"message": user.Email})
	}
}

*/
