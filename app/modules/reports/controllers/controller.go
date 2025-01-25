package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simulator/app/modules/reports/models"
	"simulator/app/modules/reports/repository"
	"simulator/systems"
)

var router systems.Router
var reportRepository repository.ReportConfigurationRepository

func Setup(r systems.Router, rcr repository.ReportConfigurationRepository) {
	router = r
	reportRepository = rcr

	configureRoute()
}

func configureRoute() {
	report := router.GetRouter().Group("/reports")

	report.POST("/", postReportConfiguration)
}

func postReportConfiguration(context *gin.Context) {
	var newReportConfiguration models.ReportConfiguration

	if err := context.BindJSON(&newReportConfiguration); err != nil {
		log.Fatal(err)
	}

	_, err := reportRepository.Create(&newReportConfiguration)
	if err != nil {
		return
	}

	context.String(http.StatusCreated, "")
}
