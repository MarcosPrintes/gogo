package main

import (
	"net/http"

	. "github.com.br/MarcosPrintes/nms/models"
	"github.com.br/MarcosPrintes/nms/templates"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Service struct {
	Echo    *echo.Echo
	Request *Mail
}

type DataEmail struct {
	Type         string      `json:"type"`
	Subject      string      `json:"subject"`
	Destinations []string    `json:"destinations"`
	Message      interface{} `json:"message"`
}

var templatesModels = map[string]string{
	"header":         templates.TemplateHeader,
	"footer":         templates.TemplateFooter,
	"bank":           templates.TemplateBank,
	"transferToUser": templates.TemplateTransferUser,
	"recharge":       templates.TemplateRecharge,
	"payment":        templates.TemplatePayment,
}

func (service *Service) startService() {

	service.Echo = echo.New()

	service.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.POST, echo.GET, echo.OPTIONS, echo.PUT, echo.DELETE, echo.HEAD},
	}))
	service.Echo.POST("/send", service.s)
	service.Echo.Logger.Fatal(service.Echo.Start(":3008"))
}

//===================== handler =====================
func (service *Service) sendEmails(c echo.Context) error {
	email := new(DataEmail)
	if err := c.Bind(email); err != nil {
		return c.JSON(http.StatusInternalServerError, "bind error: "+err.Error())
	}

	r := NewMail(email.Destinations, email.Subject)

	template := templatesModels["header"] + templatesModels[email.Type] + templatesModels["footer"]

	if status := r.Send(template, email.Message); status {
		return c.JSON(http.StatusOK, "email enviado")
	}
	return c.JSON(http.StatusInternalServerError, "email não enviado")
}

func (service *Service) s(c echo.Context) error {
	return service.sendEmails(c)
}
