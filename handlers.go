package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

//RESTREAMERS

func GetRestreamers(c *fiber.Ctx) error {
	claims := c.Context().Value("tokenClaims").(jwt.MapClaims)
	customer := GetCustomerNameFromJwt(claims)

	var restr []Restreamer
	err := DB.Preload("RestreamerSrt").Preload("RestreamerSettings").Where("owner = ?", customer).Find(&restr)
	if err != nil {
		return c.Status(500).JSON("error during fetch restreamer")
	}
	return c.Status(200).JSON(restr)
}

func GetRestreamersWithApiKey(c *fiber.Ctx) error {
	var requestBody struct {
		ApiKey   string `json:"api_key"`
		Customer string `json:"customer"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	//getting apikey info
	var apikeyInfo CustomerApiKey
	err := DB.Where("id = ?", fmt.Sprintf("%s_id", requestBody.Customer)).Find(&apikeyInfo)
	if err != nil {
		return c.Status(500).JSON("error during fetch apikey info")
	}
	if apikeyInfo.APIKey != requestBody.ApiKey && apikeyInfo.ID != fmt.Sprintf("%s_id", requestBody.Customer) {
		return c.Status(401).JSON("apikey does not match")
	}
	//getting restreamers
	var restr []Restreamer
	err = DB.Preload("RestreamerSrt").Preload("RestreamerSettings").Where("owner = ?", requestBody.Customer).Find(&restr)
	if err != nil {
		return c.Status(500).JSON("error during fetch restreamer")
	}
	return c.Status(200).JSON(restr)
}

func CreateRestreamer(c *fiber.Ctx) error {
	var requestBody Restreamer
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	err := DB.Save(&requestBody)
	if err != nil {
		return c.Status(500).JSON("error during create restreamer")
	}
	return c.SendStatus(201)
}

func UpdateRestreamer(c *fiber.Ctx) error {
	var requestBody Restreamer
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	err := DB.Save(&requestBody)
	if err != nil {
		return c.Status(500).JSON("error during update restreamer")
	}
	return c.SendStatus(200)
}

// //PROXY PUSH / PULL

// func ProxyPushStart(c *fiber.Ctx) error {

// }

// func ProxyPushStop(c *fiber.Ctx) error {

// }

// func ProxyPullStart(c *fiber.Ctx) error {

// }

// func ProxyPullStop(c *fiber.Ctx) error {

// }

// //EVENTS

func CreateEvent(c *fiber.Ctx) error {
	var requestBody struct {
		Customer     string    `json:"customer"`
		InstanceName string    `json:"instance_name"`
		StartDate    time.Time `json:"start_date"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	var customerId int // TODO: change to customer name

	newEvent := RestreamerEvent{
		EventId:      uuid.NewString(),
		StartAt:      requestBody.StartDate,
		UpdatedAt:    time.Now(),
		InstanceName: requestBody.InstanceName,
		CustomerId:   customerId,
		Active:       0,
		Completed:    0,
	}
	err := DB.Save(&newEvent)
	if err != nil {
		return c.Status(500).JSON("error during create event")
	}
	return c.SendStatus(201)
}

func StarEvent(c *fiber.Ctx) error {
	var requestBody struct {
		EventId string `json:"event_id"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	err := DB.Model(RestreamerEvent{}).Where("event_id = ?", requestBody.EventId).Updates(&RestreamerEvent{
		Active:    1,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return c.Status(500).JSON("error during update for start event")
	}
	return c.SendStatus(200)
}

func SetEventCompleted(c *fiber.Ctx) error {
	var requestBody struct {
		EventId string `json:"event_id"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	err := DB.Model(RestreamerEvent{}).Where("event_id = ?", requestBody.EventId).Updates(&RestreamerEvent{
		Active:    0,
		UpdatedAt: time.Now(),
		Completed: 1,
	})
	if err != nil {
		return c.Status(500).JSON("error during update for complete event")
	}
	return c.SendStatus(200)
}

func EventsHistory(c *fiber.Ctx) error {
	var requestBody struct {
		Customer string `json:"customer"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	var customerId int //TODO; change to customer name

	var events []RestreamerEvent
	err := DB.Model(RestreamerEvent{}).Where("customer_id = ?", customerId).Find(&events)
	if err != nil {
		return c.Status(500).JSON("error during getting restreamer events history")
	}
	return c.Status(200).JSON(events)
}

// //VIDEO UPLOADS

func GetVideos(c *fiber.Ctx) error {
	var requestBody struct {
		Amount      int     `json:"amount"`
		StartFrom   int     `json:"start_from"`
		Category    *string `json:"category"`
		TitleFilter *string `json:"q"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	claims := c.Context().Value("tokenClaims").(jwt.MapClaims)
	customer := GetCustomerNameFromJwt(claims)

	var uploads []Upload
	query := DB.Where("client_id = ?", fmt.Sprintf("%s_client", customer))

	if requestBody.Category != nil {
		query.Where("category = ?", requestBody.Category)
	}

	if requestBody.TitleFilter != nil {
		query.Where("title like ?", requestBody.TitleFilter)
	}

	err := query.Offset(requestBody.StartFrom).Limit(requestBody.Amount).Find(&uploads)
	if err != nil {
		return c.Status(500).JSON("error during fetch videouploads")
	}

	return c.Status(200).JSON(uploads)
}

func GetVideosWithApiKey(c *fiber.Ctx) error {
	var requestBody struct {
		ApiKey      string  `json:"api_key"`
		Customer    string  `json:"customer"`
		Amount      int     `json:"amount"`
		StartFrom   int     `json:"start_from"`
		Category    *string `json:"category"`
		TitleFilter *string `json:"q"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}
	//getting apikey info
	var apikeyInfo CustomerApiKey
	err := DB.Where("id = ?", fmt.Sprintf("%s_id", requestBody.Customer)).Find(&apikeyInfo)
	if err != nil {
		return c.Status(500).JSON("error during fetch apikey info")
	}
	if apikeyInfo.APIKey != requestBody.ApiKey && apikeyInfo.ID != fmt.Sprintf("%s_id", requestBody.Customer) {
		return c.Status(401).JSON("apikey does not match")
	}

	var uploads []Upload
	query := DB.Where("client_id = ?", fmt.Sprintf("%s_client", requestBody.Customer))

	if requestBody.Category != nil {
		query.Where("category = ?", requestBody.Category)
	}

	if requestBody.TitleFilter != nil {
		query.Where("title like ?", requestBody.TitleFilter)
	}

	err = query.Offset(requestBody.StartFrom).Limit(requestBody.Amount).Find(&uploads)
	if err != nil {
		return c.Status(500).JSON("error during fetch videouploads")
	}
	return c.Status(200).JSON(uploads)
}

// func CreateVideo(c *fiber.Ctx) error {

// }

// func DeleteVideo(c *fiber.Ctx) error {

// }

// func GetCategories(c *fiber.Ctx) error {

// }

// func GetCategoriesWithApiKey(c *fiber.Ctx) error {

// }

// func CreateCategory(c *fiber.Ctx) error {

// }

// //CUSTOMERS

// func CreateCustomer(c *fiber.Ctx) error {

// }

// func GetCustomer(c *fiber.Ctx) error {

// }

// func CreateUser(c *fiber.Ctx) error {

// }

// func UpdateUser(c *fiber.Ctx) error {

// }

// func DeleteUser(c *fiber.Ctx) error {

// }
