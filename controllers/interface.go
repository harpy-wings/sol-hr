package controllers

import "github.com/kataras/iris/v12"

type IController interface {
	Register(App *iris.Application) error
}

type GenericResponse struct {
	Success      bool                   `json:"success"`
	Message      string                 `json:"message"`
	ErrorCode    int                    `json:"error_code"`
	Meta         map[string]interface{} `json:"meta"`
	TotalRecords int                    `json:"total_records"`
	// Offset       int                    `json:"offset"`
	// Limit        int                    `json:"limit"`
	Data interface{} `json:"data"`
}
