package offer

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/BoyerDamien/ressources/tag"
	"github.com/BoyerDamien/ressources/testUtils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	url     = "/api/v1"
	app     = testUtils.SetupApp(url, &Offer{}, &tag.Tag{})
	urlOne  = fmt.Sprintf("%s/offer", url)
	urlList = fmt.Sprintf("%ss", urlOne)
	tester  = testUtils.TestApi{App: app}
)

/*****************************************************************************
 *					Test empty routes
 ****************************************************************************/

func Test_GET_Offer_Empty(t *testing.T) {

	url := fmt.Sprintf("%s/test", urlOne)

	resp, err := app.Test(httptest.NewRequest("GET", url, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

func Test_GET_Offer_List_Empty(t *testing.T) {

	resp, err := app.Test(httptest.NewRequest("GET", urlList, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
}

func Test_DELETE_Offer_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_DELETE_Offer_List_Empty(t *testing.T) {
	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s?name=test", urlList), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusAccepted, resp.StatusCode, "Status code")
}

func Test_PUT_Offer_Empty(t *testing.T) {
	data := Offer{
		Name: "test",
	}
	var result Offer
	resp, err := tester.Update(url, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *					Test create routes
 ****************************************************************************/

func Test_POST_Offer_without_tags(t *testing.T) {
	data := Offer{
		Name:        "test",
		Description: "description",
	}
	var result Offer
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusBadRequest, resp.StatusCode, "Status code")
}

func Test_POST_Offer_with_tags(t *testing.T) {
	data := Offer{
		Name:        "test",
		Description: "description",
		Tags: []tag.Tag{
			{
				Name: "Tag1",
			},
		},
	}

	var result Offer
	resp, err := tester.Create(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 Offer
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")

	var tagFound tag.Tag
	resp, err = tester.Retrieve(fmt.Sprintf("%s/tag/%s", url, data.Tags[0].Name), &tagFound)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data.Tags[0]), testUtils.ModelToString(tagFound), "Value")

}

/*****************************************************************************
 *			Test retrieve routes
 ****************************************************************************/
func Test_GET_Offer(t *testing.T) {
	data := Offer{
		Name:        "test",
		Description: "description",
		Tags: []tag.Tag{
			{
				Name: "Tag1",
			},
		},
	}

	tester.Create(urlOne, &data, nil)

	var result Offer
	resp, err := tester.Retrieve(urlOne+"/test", &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

}

/*****************************************************************************
 *			Test Update Routes
 ****************************************************************************/
func Test_PUT_Offer_simple(t *testing.T) {
	data := Offer{
		Name:        "test",
		Description: "description",
		Tags: []tag.Tag{
			{
				Name: "Tag1",
			},
		},
	}
	tester.Create(urlOne, &data, nil)

	data.Description = "Changed"

	var result Offer
	resp, err := tester.Update(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 Offer
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")
}

func Test_PUT_Offer_update_tag(t *testing.T) {
	data := Offer{
		Name:        "test",
		Description: "description",
		Tags: []tag.Tag{
			{
				Name: "Tag1",
			},
		},
	}
	tester.Create(urlOne, &data, nil)

	data.Tags[0].Name = "Tag2"

	var result Offer
	resp, err := tester.Update(urlOne, &data, &result)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result), "Value")

	var result2 Offer
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, testUtils.ModelToString(data), testUtils.ModelToString(result2), "Value")

	// 	var tagFound tag.Tag
	// 	resp, err = tester.Retrieve(fmt.Sprintf("%s/tag/%s", url, data.Tags[0].Name), &tagFound)
	// 	utils.AssertEqual(t, nil, err, "app.Test")
	// 	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	// 	utils.AssertEqual(t, testUtils.ModelToString(data.Tags[0]), testUtils.ModelToString(tagFound), "Value")
	//
}

/****************************************************************************
 *			Test Delete Routes
 ****************************************************************************/
func Test_DELETE_User(t *testing.T) {
	data := Offer{
		Name:        "test",
		Description: "description",
		Tags: []tag.Tag{
			{
				Name: "Tag1",
			},
		},
	}
	tester.Create(urlOne, &data, nil)

	resp, err := app.Test(httptest.NewRequest("DELETE", fmt.Sprintf("%s/test", urlOne), nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 Offer
	resp, err = tester.Retrieve(urlOne+"/test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode, "Status code")
}

/*****************************************************************************
 *			Test GET list Routes
 ****************************************************************************/
func Test_GET_Offer_List(t *testing.T) {
	data := []Offer{
		{
			Name:        "test3",
			Description: "Offre gratuite!",
			Tags: []tag.Tag{
				{
					Name: "Tag1",
				},
			},
		},
		{
			Name:        "test2",
			Description: "120 euros",
			Tags: []tag.Tag{
				{
					Name: "Payant",
				},
			},
		},
		{
			Name:        "test1",
			Description: "10 euros",
			Tags: []tag.Tag{
				{
					Name: "Minimum",
				},
			},
		},
	}
	for _, val := range data {
		tester.Create(urlOne, &val, nil)
	}

	var all []Offer
	resp, err := tester.Retrieve(urlList+"?orderBy=name", &all)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, len(data), len(all))
	utils.AssertEqual(t, "test1", all[0].Name)

	var results []Offer
	resp, err = tester.Retrieve(urlList+"?toFind=euros", &results)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 2, len(results))

	var found []Offer
	resp, err = tester.Retrieve(urlList+"?tofind=grat", &found)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 1, len(found), "Size")

	var limit []Offer
	resp, err = tester.Retrieve(urlList+"?limit=2&offset=1", &limit)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 2, len(limit), "Size")
}

/*****************************************************************************
 *			Test Delete list Routes
 ****************************************************************************/
func Test_DELETE_Offer_List(t *testing.T) {
	data := []Offer{
		{
			Name:        "test3",
			Description: "Offre gratuite!",
			Tags: []tag.Tag{
				{
					Name: "Tag1",
				},
			},
		},
		{
			Name:        "test2",
			Description: "120 euros",
			Tags: []tag.Tag{
				{
					Name: "Payant",
				},
			},
		},
		{
			Name:        "test1",
			Description: "10 euros",
			Tags: []tag.Tag{
				{
					Name: "Minimum",
				},
			},
		},
	}
	for _, val := range data {
		tester.Create(urlOne, &val, nil)
	}

	endpoint := urlList + "?names=test3,test2,test1"
	resp, err := app.Test(httptest.NewRequest("DELETE", endpoint, nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")

	var result2 []Offer
	resp, err = tester.Retrieve(urlList+"?toFind=test", &result2)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, fiber.StatusOK, resp.StatusCode, "Status code")
	utils.AssertEqual(t, 0, len(result2), "Size")
}
