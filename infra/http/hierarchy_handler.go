package http

import (
	"net/http"
	"personia/app"

	"github.com/labstack/echo/v4"
)

func RegisterHierarchy(api *echo.Group, hrUC app.HRUC) {
	hier := api.Group("/hierarchy")
	hier.PUT("", updateHierarchy(hrUC))
	hier.GET("", getHierarchy(hrUC))
	hier.GET("/:name", getSupervisor(hrUC))
}

func getSupervisor(hrUC app.HRUC) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		resp, err := hrUC.GetSupervisor(Context(c), name)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	}
}

func getHierarchy(hrUC app.HRUC) echo.HandlerFunc {
	return func(c echo.Context) error {
		hier, err := hrUC.GetHierarchy(Context(c))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, hier)
	}
}

func updateHierarchy(hrUC app.HRUC) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req map[string]string
		if err := c.Bind(&req); err != nil {
			return err
		}
		if err := hrUC.UpdateHierachy(Context(c), req); err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	}
}
