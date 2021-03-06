// @APIVersion 1.0.0
// @Title gongan API
// @Description gongan API usually works as expected. But sometimes its not true
// @Contact tiandage719@outlook.com
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
// @BasePath http://127.0.0.1:8000/api/v1
// @SubApi Auth management API [/]
// @SubApi Admins management API [/admins]
// @SubApi Users management API [/users]
// @SubApi Upload management API [/upload]

package api

import (
	"../config"
	"./v1"	
	"github.com/labstack/echo"
)

// RouteAPI contains router groups for API
func RouteAPI(parentRoute *echo.Echo) {
	route := parentRoute.Group(config.APIURL)
	{
		v1.InitBlockchain(route)
	}
}
