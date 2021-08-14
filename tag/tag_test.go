package tag

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/BoyerDamien/ressources/testUtils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	url     = "/api/v1"
	app     = testUtils.SetupApp(url, &Tag{})
	urlOne  = fmt.Sprintf("%s/tag", url)
	urlList = fmt.Sprintf("%ss", urlOne)
	tester  = testUtils.TestApi{App: app}
)

/****************************************************************************************
*				Test empty routes
****************************************************************************************/
func Test_GET_Tag_Empty(t *testing.T) {

	url := fmt.Sprintf("%s/test", urlOne)

	resp, err := app.Test(httptest.NewRequest("GET", url, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

func Test_GET_Tag_List_Empty(t *testing.T) {

	resp, err := app.Test(httptest.NewRequest("GET", urlList, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func Test_DELETE_Tag_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_DELETE_Tag_List_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s?names=test", urlList), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test create routes
 ****************************************************************************/

func Test_POST_Tag(t *testing.T) {
	data := Tag{
		Name: "test",
	}
	var result Tag
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 Tag
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")
}

/*****************************************************************************
 *					Test retrieve routes
 ****************************************************************************/
func Test_GET_Tag(t *testing.T) {
	data := Tag{
		Name: "test",
	}
	tester.Create(urlOne, &data, nil)

	var result Tag
	resp, err := tester.Retrieve(urlOne+"/test", &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")
}

/*****************************************************************************
 *					Test Delete Routes
 ****************************************************************************/
func Test_DELETE_Tag(t *testing.T) {
	data := Tag{
		Name: "test",
	}
	tester.Create(urlOne, &data, nil)

	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 Tag
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test GET list Routes
 ****************************************************************************/
func Test_GET_Tag_List(t *testing.T) {
	data := []Tag{
		{
			Name: "test",
		},
		{
			Name: "okok",
		},
		{
			Name: "aaaa",
		},
		{
			Name: "tata",
		},
		{
			Name: "kaka",
		},
	}

	for _, val := range data {
		tester.Create(urlOne, &val, nil)
	}

	var all []Tag
	resp, err := tester.Retrieve(urlList+"?orderBy=name", &all)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, len(data), len(all))
	utils.AssertEqual(t, "aaaa", all[0].Name)

	var filter []Tag
	resp, err = tester.Retrieve(urlList+"?toFind=a", &filter)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 3, len(filter))

	var found []Tag
	resp, err = tester.Retrieve(urlList+"?tofind=ok", &found)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 1, len(found), "Size")

	var limit []Tag
	resp, err = tester.Retrieve(urlList+"?limit=2&offset=3", &limit)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 2, len(limit), "Size")
}

/*****************************************************************************
 *					Test Delete list Routes
 ****************************************************************************/
func Test_DELETE_Tag_List(t *testing.T) {
	data := []Tag{
		{
			Name: "test",
		},
		{
			Name: "okok",
		},
		{
			Name: "aaaa",
		},
		{
			Name: "tata",
		},
		{
			Name: "kaka",
		},
	}

	for _, val := range data {
		tester.Create(urlOne, &val, nil)
	}

	endpoint := urlList + "?names=test,okok,aaaa,tata,kaka"
	resp, err := app.Test(httptest.NewRequest("DELETE", endpoint, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 []Tag
	resp, err = tester.Retrieve(urlList+"?toFind=list", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 0, len(result2), "Size")
}
