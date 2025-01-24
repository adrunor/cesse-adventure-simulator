package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simulator/app/modules/reports/models"
	"simulator/systems"
)

var router systems.Router
var database systems.Database

func Setup(context *systems.Context) {
	router = context.Router()
	database = context.Database()

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

	context.String(http.StatusCreated, "")
}
