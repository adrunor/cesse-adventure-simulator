package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simulator/app/modules/level_simulation/models"
	"simulator/app/modules/level_simulation/repository"
	"simulator/app/modules/level_simulation/services"
	"simulator/systems"
	"simulator/systems/api"
	"strconv"
)

var router systems.Router
var reportRepository repository.LevelConfigRepository

func Setup(r systems.Router, rcr repository.LevelConfigRepository) {
	router = r
	reportRepository = rcr

	configureRoute()
}

func configureRoute() {
	report := router.GetRouter().Group("/level_simulation")

	report.GET("/", getReportConfigurations)
	report.GET("/:id", getReportConfigurationById)
	report.POST("/", postReportConfiguration)
	report.POST("/:id/result", postReportResult)
}

func getReportConfigurations(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "")
}

func getReportConfigurationById(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)

	//reportRepository.Find(id)

	ctx.String(http.StatusNotImplemented, "")
}

func postReportConfiguration(ctx *gin.Context) {
	var newReportConfiguration models.LevelConfig
	code := http.StatusCreated

	if err := ctx.ShouldBind(&newReportConfiguration); err != nil {
		code = http.StatusBadRequest
		ctx.JSON(code, &api.ErrorResponse{Code: code, Message: err.Error()})
		return
	}

	_, err := reportRepository.Create(&newReportConfiguration)
	if err != nil {
		code = http.StatusBadRequest
		ctx.JSON(code, &api.ErrorResponse{Code: code, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func postReportResult(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	levelConfig := models.LevelConfig{}
	err = reportRepository.Find(&levelConfig, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
	}

	fmt.Println(levelConfig)

	levels := services.GenerateLevelTable(&levelConfig)

	result := api.ArrayResponse{
		Items: levels,
		Count: len(*levels),
	}

	ctx.JSON(http.StatusOK, result)
}
