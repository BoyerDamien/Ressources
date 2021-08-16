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
 *			Test retrieve routes
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

/*****************************************************************************
 *			Test Update Routes
 ****************************************************************************/
func Test_PUT_PortFolio_simple(t *testing.T) {
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
	data.Description = "Changed"
	data.Website = "https://changed.com"
	resp, err = tester.Update(urlOne, &data, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")

	var result3 PortFolio
	resp, err = tester.Retrieve(urlOne+"/"+data.Name, &result3)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result3), "Value")
}

func Test_PUT_PortFolio_Tag(t *testing.T) {
	data := PortFolio{
		Name:        "test",
		Description: "Changed",
		Gallery:     []media.Media{},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
		Website: "https://changed.com",
	}

	var result PortFolio
	resp, err := tester.Create(urlOne, &data, &result)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 PortFolio
	data.Tags = append(data.Tags, tag.Tag{Name: "test2"})
	resp, err = tester.Update(urlOne, &data, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")

	var result3 PortFolio
	resp, err = tester.Retrieve(urlOne+"/"+data.Name, &result3)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result3), "Value")
}

func Test_PUT_PortFolio_Gallery(t *testing.T) {
	var mediaResult media.Media
	resp, err := tester.CreateForm(fmt.Sprintf("%s/media", url), "../testFile2.json", "media", &mediaResult)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	data := PortFolio{
		Name:        "test",
		Description: "Changed",
		Gallery:     []media.Media{},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
		Website: "https://changed.com",
	}

	var result PortFolio
	resp, err = tester.Create(urlOne, &data, &result)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 PortFolio
	data.Gallery = append(data.Gallery, mediaResult)
	resp, err = tester.Update(urlOne, &data, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")

	var result3 PortFolio
	resp, err = tester.Retrieve(urlOne+"/"+data.Name, &result3)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result3), "Value")
}

func Test_PUT_PortFolio_Wrong_Gallery(t *testing.T) {
	data := PortFolio{
		Name:        "test2",
		Description: "Changed",
		Gallery:     []media.Media{},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
		Website: "https://changed.com",
	}

	var result PortFolio
	resp, err := tester.Create(urlOne, &data, &result)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 PortFolio
	data.Gallery = append(data.Gallery, media.Media{
		Name: "wrongMedia",
		Path: "wrongMedia",
	})
	resp, err = tester.Update(urlOne, &data, &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *			Test Delete Routes
 ****************************************************************************/
func Test_DELETE_User(t *testing.T) {
	var mediaResult media.Media
	resp, err := tester.CreateForm(fmt.Sprintf("%s/media", url), "../testFile2.json", "media", &mediaResult)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	data := PortFolio{
		Name:        "test",
		Description: "Changed",
		Gallery:     []media.Media{},
		Tags: []tag.Tag{
			{
				Name: "test",
			},
		},
		Website: "https://changed.com",
	}
	tester.Create(urlOne, &data, nil)

	resp, err = app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 PortFolio
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}
