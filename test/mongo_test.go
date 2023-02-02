package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffthorne/tasky/app/database"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var app *fiber.App

func init() {
	//database.Init()
	//app = fiber.New()
	//routes.SetupRoutes(app)
}

func TestCreateOrg(t *testing.T) {
	database := database.InitMongo()
	assert.Equal(t, os.Getenv("MONGO_DB"), database.Name(), "Expecting true for successful mongo connection")
}

/*
func TestAddUser(t *testing.T) {
	wiz := models.Organization{}
	database.DB.Preload("Users").Where("name = ?", "Wiz Inc").First(&wiz)
	fmt.Println("Returned Org:", len(wiz.Users))
}

func TestRequiredField(t *testing.T) {
	jen := models.User{Email: "jen@u.washington.edu", Password: "test", FirstName: "Jen", LastName: "Smith"}
	jen.HashPassword()
	org := models.Organization{Users: []models.User{jen}}
	result := database.DB.Create(&org)
	fmt.Println("RESULT: ", result)
}

func TestGetUserOrgEager(t *testing.T) {
	jeff := models.User{}
	database.DB.Preload("Organization").Where("email = ?", "jthorne@u.washington.edu").First(&jeff)

	if jeff.Organization.Name != "Wiz Inc" {
		t.Errorf("Org not found")
	}
	if jeff.CheckPassword("testing") != true {
		t.Errorf("Password incorrect")
	}
}

func TestInvalidSignup(t *testing.T) {
	testBody := map[string]interface{}{"FirstName": "jeff", "LastName": "thorne", "Email": "jthorneu.washington.edu", "Password": "testing", "ConfirmPassword": "testing"}
	body, _ := json.Marshal(testBody)

	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println("BODY: ", string(body))
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	var jsonData = map[string][]map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &jsonData); err != nil {
		fmt.Println("ERROR: ", err)
	}

	assert.Equal(t, jsonData["message"][0]["Email"], "Invalid email")
	assert.Equal(t, 400, resp.StatusCode)
}
*/
