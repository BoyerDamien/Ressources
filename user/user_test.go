package user

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	url     = "/api/v1"
	app     = SetupApp(url)
	urlOne  = fmt.Sprintf("%s/user", url)
	urlList = fmt.Sprintf("%ss", urlOne)
	tester  = testApi{App: app}
)

/*****************************************************************************
 *					Test empty routes
 ****************************************************************************/

func Test_GET_User_Empty(t *testing.T) {

	url := fmt.Sprintf("%s/dboyer", urlOne)

	resp, err := app.Test(httptest.NewRequest("GET", url, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

func Test_GET_User_List_Empty(t *testing.T) {

	resp, err := app.Test(httptest.NewRequest("GET", urlList, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func Test_DELETE_User_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test@gmail.com", url), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

func Test_DELETE_User_List_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s?emails=test@gmail.com", urlList), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_PUT_User_Empty(t *testing.T) {
	tester := testApi{App: app}
	data := User{
		Email:    "test@gmail.com",
		Password: "test",
		Role:     "admin",
	}
	var result User
	resp, err := tester.Update(url, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test create routes
 ****************************************************************************/

func Test_POST_User(t *testing.T) {
	data := User{
		Email:    "test@gmail.com",
		Password: "test",
		Role:     "admin",
	}
	var result User
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	data.Password = ""
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")

	var result2 User
	resp, err = tester.Retrieve(urlOne+"/test@gmail.com", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result2), "Value")
}

func Test_POST_User_Wrong_Mail(t *testing.T) {
	data := User{
		Email:    "testgmail.com",
		Password: "test",
		Role:     "admin",
	}
	var result User
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

func Test_POST_User_Wrong_Role(t *testing.T) {
	data := User{
		Email:    "testgmail.com",
		Password: "test",
		Role:     "afr",
	}
	var result User
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

func Test_POST_User_Wrong_Password(t *testing.T) {
	data := User{
		Email:    "testgmail.com",
		Password: "",
		Role:     "admin",
	}
	var result User
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test retrieve routes
 ****************************************************************************/
func Test_GET_User(t *testing.T) {
	data := User{
		Email:    "test@gmail.com",
		Password: "test",
		Role:     "admin",
	}
	tester.Create(urlOne, &data, nil)
	data.Password = ""

	var result User
	resp, err := tester.Retrieve(urlOne+"/test@gmail.com", &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")
}

/*****************************************************************************
 *					Test Update Routes
 ****************************************************************************/
func Test_PUT_User(t *testing.T) {
	data := User{
		Email:    "test@gmail.com",
		Password: "test",
		Role:     "admin",
	}
	tester.Create(urlOne, &data, nil)
	data.Age = 30
	data.FirstName = "test"
	data.LastName = "oko"

	var result User
	resp, err := tester.Update(urlOne, &data, &result)

	data.Password = ""
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")

	var result2 User
	resp, err = tester.Retrieve(urlOne+"/test@gmail.com", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result2), "Value")
}

func Test_PUT_User_Invalid_Column(t *testing.T) {
	data := User{
		Email:     "test@gmail.com",
		Password:  "test",
		Role:      "admin",
		Age:       30,
		FirstName: "test",
		LastName:  "oko",
	}
	tester.Create(urlOne, &data, nil)

	data.Role = "user"
	data.Password = "okoko"
	var result User
	resp, err := tester.Update(urlOne, &data, &result)

	data.Password = ""
	data.Role = "admin"
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")

	var result2 User
	resp, err = tester.Retrieve(urlOne+"/test@gmail.com", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result2), "Value")
}
