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
	urlOne  = fmt.Sprintf("%s/tag", url)
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
