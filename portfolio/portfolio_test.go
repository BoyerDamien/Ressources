package portfolio

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/BoyerDamien/ressources/media"
	"github.com/BoyerDamien/ressources/tag"
	"github.com/BoyerDamien/ressources/testUtils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	url     = "/api/v1"
	app     = testUtils.SetupApp(url, &PortFolio{}, &media.Media{}, &tag.Tag{})
	urlOne  = fmt.Sprintf("%s/portfolio", url)
	urlList = fmt.Sprintf("%ss", urlOne)
	tester  = testUtils.TestApi{App: app}
)

/****************************************************************************************
*				Test empty routes
****************************************************************************************/

func Test_GET_PortFolio_Empty(t *testing.T) {

	url := fmt.Sprintf("%s/test", urlOne)

	resp, err := app.Test(httptest.NewRequest("GET", url, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

func Test_GET_PortFolio_List_Empty(t *testing.T) {

	resp, err := app.Test(httptest.NewRequest("GET", urlList, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func Test_DELETE_PortFolio_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_DELETE_PortFolio_List_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s?names=test", urlList), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *			Test create routes
 ****************************************************************************/

func Test_POST_PortFolio(t *testing.T) {
	var mediaResult media.Media
	resp, err := tester.CreateForm(fmt.Sprintf("%s/media", url), "../testFile.txt", "media", &mediaResult)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	data := PortFolio{
		Name:        "test",
		Description: "test",
		Gallery:     []media.Media{mediaResult},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
		Website: "https://test.com",
	}

	var result PortFolio
	resp, err = tester.Create(urlOne, &data, &result)
	data.Gallery = []media.Media{}
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 PortFolio
	resp, err = tester.Retrieve(urlOne+"/"+data.Name, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")
}

func Test_POST_PortFolio_without_website(t *testing.T) {
	var mediaResult media.Media
	resp, err := tester.CreateForm(fmt.Sprintf("%s/media", url), "../testFile.txt", "media", &mediaResult)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	data := PortFolio{
		Name:        "test",
		Description: "test",
		Gallery:     []media.Media{mediaResult},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
	}

	var result PortFolio
	resp, err = tester.Create(urlOne, &data, &result)
	data.Gallery = []media.Media{}
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

func Test_POST_PortFolio_without_tag(t *testing.T) {
	var mediaResult media.Media
	resp, err := tester.CreateForm(fmt.Sprintf("%s/media", url), "../testFile.txt", "media", &mediaResult)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	data := PortFolio{
		Name:        "test",
		Description: "test",
		Gallery:     []media.Media{mediaResult},
		Website:     "https://test.com",
	}
	var result PortFolio
	resp, err = tester.Create(urlOne, &data, &result)
	data.Gallery = []media.Media{}
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test retrieve routes
 ****************************************************************************/
func Test_GET_PortFolio(t *testing.T) {
	data := PortFolio{
		Name:        "test",
		Description: "test",
		Gallery:     []media.Media{},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
		Website: "https://test.com",
	}

	var result PortFolio
	resp, err := tester.Create(urlOne, &data, &result)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 PortFolio
	resp, err = tester.Retrieve(urlOne+"/"+data.Name, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")
}
