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

	url := fmt.Sprintf("%s/test@gmail.com", urlOne)

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
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test@gmail.com", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
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
		LastName: "test",
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
		LastName: "test",
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
		LastName: "test",
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
		LastName: "test",
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
		LastName: "test",
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
		LastName: "test",
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
		LastName: "test",
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

/*****************************************************************************
 *					Test Delete Routes
 ****************************************************************************/
func Test_DELETE_User(t *testing.T) {
	data := User{
		Email:     "test@gmail.com",
		Password:  "test",
		Role:      "admin",
		Age:       30,
		FirstName: "test",
		LastName:  "oko",
	}
	tester.Create(urlOne, &data, nil)

	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test@gmail.com", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 User
	resp, err = tester.Retrieve(urlOne+"/test@gmail.com", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *			Test GET list Routes
 ****************************************************************************/
func Test_GET_User_List(t *testing.T) {
	data := []User{
		{
			Email:     "a@gmail.com",
			Password:  "test",
			Role:      "user",
			Age:       70,
			FirstName: "test",
			LastName:  "list",
		},
		{
			Email:     "b@gmail.com",
			Password:  "test",
			Role:      "customer",
			Age:       50,
			FirstName: "ok",
			LastName:  "list",
		},
		{
			Email:     "c@gmail.com",
			Password:  "test",
			Role:      "admin",
			Age:       40,
			FirstName: "ok",
			LastName:  "list",
		},
		{
			Email:     "d@gmail.com",
			Password:  "test",
			Role:      "admin",
			Age:       10,
			FirstName: "test",
			LastName:  "list",
		},
	}
	for _, val := range data {
		tester.Create(urlOne, &val, nil)
		val.Password = ""
	}
	var all []User
	resp, err := tester.Retrieve(urlList+"?orderBy=email", &all)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, len(data), len(all))
	utils.AssertEqual(t, "a@gmail.com", all[0].Email)

	var admins []User
	resp, err = tester.Retrieve(urlList+"?role=admin", &admins)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 2, len(admins))

	var found []User
	resp, err = tester.Retrieve(urlList+"?tofind=ok", &found)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 2, len(found), "Size")

	var limit []User
	resp, err = tester.Retrieve(urlList+"?limit=2&offset=3", &limit)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 1, len(limit), "Size")
}

/*****************************************************************************
 *					Test Delete list Routes
 ****************************************************************************/
func Test_DELETE_User_List(t *testing.T) {
	data := []User{
		{
			Email:     "a@gmail.com",
			Password:  "test",
			Role:      "user",
			Age:       70,
			FirstName: "test",
			LastName:  "list",
		},
		{
			Email:     "b@gmail.com",
			Password:  "test",
			Role:      "customer",
			Age:       50,
			FirstName: "test",
			LastName:  "list",
		},
		{
			Email:     "c@gmail.com",
			Password:  "test",
			Role:      "admin",
			Age:       40,
			FirstName: "test",
			LastName:  "list",
		},
		{
			Email:     "d@gmail.com",
			Password:  "test",
			Role:      "admin",
			Age:       10,
			FirstName: "test",
			LastName:  "list",
		},
	}
	for _, val := range data {
		tester.Create(urlOne, &val, nil)
	}

	endpoint := urlList + "?emails=a@gmail.com,b@gmail.com,c@gmail.com,d@gmail.com"
	resp, err := app.Test(httptest.NewRequest("DELETE", endpoint, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 []User
	resp, err = tester.Retrieve(urlList+"?toFind=list", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 0, len(result2), "Size")
}
