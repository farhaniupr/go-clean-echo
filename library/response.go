package library

import (
	"github.com/labstack/echo/v4"
)

type JsonResponse struct {
	RequestId string      `json:"request_id"`
	Status    int         `json:"status"`
	Messages  string      `json:"messages"`
	Data      interface{} `json:"data"`
}

type JsonResponseTotal struct {
	RequestId string      `json:"request_id"`
	Status    int         `json:"status"`
	Messages  string      `json:"messages"`
	Total     int         `json:"total"`
	Data      interface{} `json:"data"`
}

func ResponseInterface(c echo.Context, statusServer int, res interface{}, msg string) error {
	if res != nil {
		c.JSON(statusServer, JsonResponse{
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			Status:    statusServer,
			Messages:  msg,
			Data:      res,
		})
	} else {
		c.JSON(statusServer, JsonResponse{
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			Status:    statusServer,
			Messages:  msg,
			Data:      map[string]interface{}{},
		})
	}

	return nil
}

func ResponseInterfaceTotal(c echo.Context, statusServer int, res interface{}, msg string, total int) error {

	if res != nil {
		c.JSON(statusServer, JsonResponseTotal{
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			Status:    statusServer,
			Messages:  msg,
			Data:      res,
			Total:     total,
		})
	} else {
		c.JSON(statusServer, JsonResponseTotal{
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			Status:    statusServer,
			Messages:  msg,
			Total:     total,
			Data:      []interface{}{},
		})
	}

	return nil
}
