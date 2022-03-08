package handlers

import (
	"net/http"
	"strconv"

	"github.com/OmarAouini/golang-api-starter/entities"
	"github.com/OmarAouini/golang-api-starter/service"
	"github.com/OmarAouini/golang-api-starter/store"
	"github.com/gin-gonic/gin"
)

func CompanyRoutes(r *gin.Engine) {
	controller := newCompanyController()
	companies := r.Group("/companies")
	companies.GET("/", controller.All)
	companies.POST("/", controller.Create)
	companies.GET("/id/:id", controller.ById)
	companies.GET("/name/:name", controller.ByName)
}

type CompanyController struct {
	Service service.CompanyService
}

func newCompanyController() *CompanyController {
	return &CompanyController{
		Service: service.CompanyService{
			Store: &store.MySqlCompanyStore{},
		},
	}
}

func (c *CompanyController) All(ctx *gin.Context) {
	comps, err := c.Service.All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comps)
}

func (c *CompanyController) ById(ctx *gin.Context) {
	param, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(param)
	comps, err := c.Service.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comps)
}

func (c *CompanyController) ByName(ctx *gin.Context) {
	name, _ := ctx.Params.Get("name")
	comps, err := c.Service.GetByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comps)
}

func (c *CompanyController) Create(ctx *gin.Context) {
	var requestBody entities.Company

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	err := c.Service.Create(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusCreated)
}
