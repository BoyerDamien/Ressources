package media

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/BoyerDamien/ressources/testUtils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	url     = "/api/v1"
	app     = testUtils.SetupApp(url, &Media{})
	urlOne  = fmt.Sprintf("%s/media", url)
	urlList = fmt.Sprintf("%ss", urlOne)
	tester  = testUtils.TestApi{App: app}
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
	resp, err := tester.CreateForm(urlOne, "../testFile.txt", "media", &result)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")

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

	tester.CreateForm(urlOne, "../testFile.txt", "media", nil)

	var result Media
	resp, err := tester.Retrieve(urlOne+"/testFile.txt", &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

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
	tester.CreateForm(urlOne, "../testFile.txt", "media", nil)
	data.Status = "open"

	var result Media
	resp, err := tester.Update(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")
	data.Status = "protected"
	tester.Update(urlOne, &data, &result)
}

func Test_PUT_Media_Wrong_Status(t *testing.T) {
	data := Media{
		Name:   "testFile.txt",
		Size:   12,
		Type:   "application/octet-stream",
		Url:    "testFile.txt",
		Status: "protected",
	}
	tester.CreateForm(urlOne, "../testFile.txt", "media", nil)
	data.Status = "fergrt"

	var result Media
	resp, err := tester.Update(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(&Media{}), testUtils.ModelToString(result), "Value")

	data.Status = "protected"
	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	fmt.Println("result2 = ", result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value2")
}

/*****************************************************************************
 *					Test Delete Routes
 ****************************************************************************/
func Test_DELETE_Media(t *testing.T) {
	tester.CreateForm(urlOne, "../testFile.txt", "media", nil)

	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/testFile.txt", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 Media
	resp, err = tester.Retrieve(urlOne+"/testFile.txt", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test GET list Routes
 ****************************************************************************/
func Test_GET_Media_List(t *testing.T) {
	data := []Media{
		{
			Name:   "testFile.txt",
			Status: "open",
		},
		{
			Name:   "testFile2.json",
			Status: "protected",
		},
		{
			Name:   "testFile3.txt",
			Status: "protected",
		},
		{
			Name:   "testFile4.txt",
			Status: "open",
		},
	}
	for _, val := range data {
		tester.CreateForm(urlOne, "../"+val.Name, "media", nil)
	}

	tester.Update(urlOne, &data[0], nil)
	tester.Update(urlOne, &data[3], nil)

	var all []Media
	resp, err := tester.Retrieve(urlList+"?orderBy=size", &all)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, len(data), len(all))
	utils.AssertEqual(t, "testFile4.txt", all[0].Name)
	utils.AssertEqual(t, "testFile.txt", all[1].Name)

	var protected []Media
	resp, err = tester.Retrieve(urlList+"?status=open", &protected)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 2, len(protected))

	var found []Media
	resp, err = tester.Retrieve(urlList+"?tofind=json", &found)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 1, len(found), "Size")

	var limit []Media
	resp, err = tester.Retrieve(urlList+"?limit=2&offset=3", &limit)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 1, len(limit), "Size")
}

/*****************************************************************************
 *					Test Delete list Routes
 ****************************************************************************/
func Test_DELETE_Media_List(t *testing.T) {
	data := []Media{
		{
			Name:   "testFile.txt",
			Status: "open",
		},
		{
			Name:   "testFile2.json",
			Status: "protected",
		},
		{
			Name:   "testFile3.txt",
			Status: "protected",
		},
		{
			Name:   "testFile4.txt",
			Status: "open",
		},
	}
	endpoint := urlList + "?Names="
	for _, val := range data {
		tester.CreateForm(urlOne, "../"+val.Name, "media", nil)
		endpoint += val.Name + ","
	}

	endpoint = endpoint[:len(endpoint)-1]
	resp, err := app.Test(httptest.NewRequest("DELETE", endpoint, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 []Media
	resp, err = tester.Retrieve(urlList, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 0, len(result2), "Size")
}
