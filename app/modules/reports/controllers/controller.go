package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simulator/app/modules/reports/models"
)

func AddRoutes(router *gin.RouterGroup) {
	report := router.Group("/reports")

	report.POST("/", postReportConfiguration)
}

func postReportConfiguration(context *gin.Context) {
	var newReportConfiguration models.ReportConfigurationForm

	if err := context.BindJSON(&newReportConfiguration); err != nil {
		return
	}

	context.String(http.StatusCreated, "")
}
