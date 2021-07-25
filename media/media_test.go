package media

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	url     = "/api/v1"
	app     = SetupApp(url)
	urlOne  = fmt.Sprintf("%s/media", url)
	urlList = fmt.Sprintf("%ss", urlOne)
	tester  = testApi{App: app}
)

/*****************************************************************************
 *					Test empty routes
 ****************************************************************************/

func Test_GET_Media_Empty(t *testing.T) {

	url := fmt.Sprintf("%s/test", urlOne)

	resp, err := app.Test(httptest.NewRequest("GET", url, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

func Test_GET_Media_List_Empty(t *testing.T) {

	resp, err := app.Test(httptest.NewRequest("GET", urlList, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func Test_DELETE_Media_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_DELETE_Media_List_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s?names=test", urlList), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_PUT_Media_Empty(t *testing.T) {
	tester := testApi{App: app}
	data := Media{
		Name: "test",
	}
	var result Media
	resp, err := tester.Update(url, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test create routes
 ****************************************************************************/

func Test_POST_Media(t *testing.T) {
	data := Media{
		Name:   "testFile.txt",
		Size:   12,
		Type:   "application/octet-stream",
		Url:    "testFile.txt",
		Status: "protected",
	}

	var result Media
	resp, err := tester.Create(urlOne, "../testFile.txt", &result)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")

	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result2), "Value")

	_, err = os.Stat("testFile.txt")
	utils.AssertEqual(t, false, os.IsNotExist(err), "Exist test")
}

/*****************************************************************************
 *					Test retrieve routes
 ****************************************************************************/
func Test_GET_Media(t *testing.T) {
	data := Media{
		Name:   "testFile.txt",
		Size:   12,
		Type:   "application/octet-stream",
		Url:    "testFile.txt",
		Status: "protected",
	}

	tester.Create(urlOne, "../testFile.txt", nil)

	var result Media
	resp, err := tester.Retrieve(urlOne+"/testFile.txt", &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")

	resp, err = app.Test(httptest.NewRequest("GET", "/static/testFile.txt", nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test Update Routes
 ****************************************************************************/
func Test_PUT_Media(t *testing.T) {
	data := Media{
		Name:   "testFile.txt",
		Size:   12,
		Type:   "application/octet-stream",
		Url:    "testFile.txt",
		Status: "protected",
	}
	tester.Create(urlOne, "../testFile.txt", nil)
	data.Status = "open"

	var result Media
	resp, err := tester.Update(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result), "Value")

	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result2), "Value")
}

func Test_PUT_Media_Wrong_Status(t *testing.T) {
	data := Media{
		Name:   "testFile.txt",
		Size:   12,
		Type:   "application/octet-stream",
		Url:    "testFile.txt",
		Status: "open",
	}
	tester.Create(urlOne, "../testFile.txt", nil)
	data.Status = "fergrt"

	var result Media
	resp, err := tester.Update(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(&Media{}), ModelToString(result), "Value")

	data.Status = "open"
	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, ModelToString(data), ModelToString(result2), "Value2")

}
