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
	companies.DELETE("/id/:id", controller.Delete)
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

// All godoc
// @Summary      All
// @Description  All companies
// @Tags         companies
// @Produce      json
// @Success      200 {array} entities.Company
// @Router       /companies [get]
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

// ById godoc
// @Summary      ById
// @Description  companies ById
// @Tags         companies
// @Param        id   path      int  true  "company id"
// @Produce      json
// @Success      200 {object} entities.Company
// @Router       /companies/id/:id [get]
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

// ByName godoc
// @Summary      ByName
// @Description  companies ByName
// @Tags         companies
// @Param        name   path      string  true  "company name"
// @Produce      json
// @Success      200 {object} entities.Company
// @Router       /companies/name/:name [get]
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

// Create godoc
// @Summary      Create
// @Description  companies create
// @Tags         companies
// @Param        company body  entities.Company  true  "company"
// @Produce      json
// @Success      201
// @Router       /companies [post]
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

// Delete godoc
// @Summary      Delete
// @Description  companies Delete
// @Tags         companies
// @Param        id   path      int  true  "company id"
// @Produce      json
// @Success      200
// @Router       /companies [delete]
func (c *CompanyController) Delete(ctx *gin.Context) {
	param, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(param)
	err := c.Service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
