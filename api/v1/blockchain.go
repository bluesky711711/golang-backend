package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"../response"
	"github.com/labstack/echo"
)

type Inputdata struct {
	Api_url  string    `json:"api_url"`
}

// InitAdmins inits admin CRUD apis
// @Title Admins
// @Description Admins's router group.
func InitBlockchain(parentRoute *echo.Group) {
	route := parentRoute.Group("/blockchain")
	route.POST("/callingapi", callingAPI)
}

func callingAPI(c echo.Context) error {
	data := &Inputdata{}
	if err := c.Bind(data); err != nil {
		return response.KnownErrJSON(c, "err.blockchain.bind", err)
	}
  resp, err := http.Get(data.Api_url)
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Println(resp.Body)
  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body));
  var sensor1 map[string]interface{}
  jsonErr := json.Unmarshal(body, &sensor1)
  if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return response.SuccessInterface(c, sensor1)
}
