package auth

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jeffthorne/tasky/app/database"
	"github.com/jeffthorne/tasky/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

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

func msgForTag(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "min of 5 characters required"
	case "taken":
		return "email taken"
	}

	return fe.Error() // default error
}

func emailExistsValidate(sl validator.StructLevel) {
	fmt.Println("FORM PR: ", sl.Current().Interface().(SignupForm).Email)
	email := sl.Current().Interface().(SignupForm).Email
	user := models.User{}
	col := database.DB.Collection("Users")
	err := col.FindOne(context.TODO(), bson.D{{"Email", email}}).Decode(&user)
	if err == nil {
		fmt.Println("CHECK EMAIL ERROR: ", err)
		sl.ReportError(sl.Current().Interface(), "Email", "Email", "taken", "")
	}

	fmt.Println("USER: ", user)
}

func ValidateStruct(form any, formKind string) []interface{} {
	var errors []interface{}
	fmt.Println("FORM TYPE: ", form)
	vald := validator.New()
	if formKind == "signup" {
		vald.RegisterStructValidation(emailExistsValidate, form)
	}
	err := vald.Struct(form)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := msgForTag(err)
			//message := fmt.Sprintf("%s %s", err.Tag(), err.Type())
			v := map[string]interface{}{err.StructField(): msg}
			errors = append(errors, &v)
		}
	}

	return errors
}
