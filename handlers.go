package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

//RESTREAMERS

// GetRestreamers godoc
// @Summary      GetRestreamers
// @Description  need jwt with customer_name
// @Tags         restreamer
// @Produce      json
// @Success      200 {array} Restreamer
// @Router       /restreamers [get]
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
	var restreamersWithApikey struct {
		ApiKey   string `json:"api_key"`
		Customer string `json:"customer"`
	}
	if err := c.BodyParser(&restreamersWithApikey); err != nil {
		return err
	}
	//getting apikey info
	var apikeyInfo CustomerApiKey
	err := DB.Where("id = ?", fmt.Sprintf("%s_id", restreamersWithApikey.Customer)).Find(&apikeyInfo)
	if err != nil {
		return c.Status(500).JSON("error during fetch apikey info")
	}
	if apikeyInfo.APIKey != restreamersWithApikey.ApiKey && apikeyInfo.ID != fmt.Sprintf("%s_id", restreamersWithApikey.Customer) {
		return c.Status(401).JSON("apikey does not match")
	}
	//getting restreamers
	var restr []Restreamer
	err = DB.Preload("RestreamerSrt").Preload("RestreamerSettings").Where("owner = ?", restreamersWithApikey.Customer).Find(&restr)
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
		Tags        *string `json:"tags"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	claims := c.Context().Value("tokenClaims").(jwt.MapClaims)
	customer := GetCustomerNameFromJwt(claims)

	var uploads []Upload
	query := DB.
		Where("client_id = ?", fmt.Sprintf("%s_client", customer))

	if requestBody.Category != nil {
		query.
			Where("category = ?", requestBody.Category)
	}

	if requestBody.TitleFilter != nil {
		query.
			Where("title like ?", requestBody.TitleFilter)
	}

	if requestBody.Tags != nil {
		query.
			Where("tags like ?", requestBody.TitleFilter)
	}

	err := query.
		Offset(requestBody.StartFrom).
		Limit(requestBody.Amount).
		Find(&uploads)
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
		Tags        *string `json:"tags"`
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
	query := DB.
		Where("client_id = ?", fmt.Sprintf("%s_client", requestBody.Customer))

	if requestBody.Category != nil {
		query.
			Where("category = ?", requestBody.Category)
	}

	if requestBody.TitleFilter != nil {
		query.
			Where("title like ?", requestBody.TitleFilter)
	}

	if requestBody.Tags != nil {
		query.
			Where("tags like ?", requestBody.TitleFilter)
	}
	err = query.
		Offset(requestBody.StartFrom).
		Limit(requestBody.Amount).
		Find(&uploads)
	if err != nil {
		return c.Status(500).JSON("error during fetch videouploads")
	}
	return c.Status(200).JSON(uploads)
}

// func GetSingleVideoWithApiKey(c *fiber.Ctx) error {

// }

func CreateVideo(c *fiber.Ctx) error {
	var createVideoRequest struct {
		Filename      string `json:"filename"`
		Size          int    `json:"size"`
		Category      int    `json:"category"`
		Title         string `json:"title"`
		Tags          string `json:"tags"`
		Location      string `json:"location"`
		Customer      string `json:"customer"`
		CustomerId    int    `json:"customer_id"`
		ReportedEmail string `json:"reporter_email"`
	}
	if err := c.BodyParser(&createVideoRequest); err != nil {
		return err
	}

	type uploadEncodingRequest struct {
		Jobid           string   `json:"jobid"`
		SrcVideo        string   `json:"src_video"`
		Profiles        []string `json:"profiles"`
		APIKey          string   `json:"api_key"`
		ThumbnailNumber int      `json:"thumbnail_number"`
		CustomerID      string   `json:"customer_id"`
	}

	type ManifestUrl struct {
		High720  string `json:"high_720"`
		High1080 string `json:"high_1080"`
		High360  string `json:"high_360"`
		High480  string `json:"high_480"`
	}

	type ThumbnailsUrl struct {
		High1080 []string `json:"high_1080"`
		High360  []string `json:"high_360"`
		Original []string `json:"original"`
		High720  []string `json:"high_720"`
		High480  []string `json:"high_480"`
	}

	type UploadEncodingResponse struct {
		Result        *string          `json:"result"`
		Jobid         *int             `json:"jobid"`
		OriginalURL   *string          `json:"original_url"`
		ManifestURL   *ManifestUrl     `json:"manifest_url"`
		ThumbnailsURL *[]ThumbnailsUrl `json:"thumbnails_Url"`
		Message       *string          `json:"message"`
	}

	//defult profiles for encoding upload
	encodingProfiles := make([]string, 0)
	encodingProfiles = append(encodingProfiles, "hls_high_1080p", "hls_high_720p", "hls_high_480p", "hls_high_360p")

	//getting apikey info
	var apikeyInfo CustomerApiKey
	err := DB.Where("id = ?", fmt.Sprintf("%s_id", createVideoRequest.Customer)).Find(&apikeyInfo)
	if err != nil {
		return c.Status(500).JSON("error during fetch apikey info")
	}
	//call restreamer to get jobid for encoding job
	encodingApiUrl := "https://api.tngrm.io/api/tangram/customer/v1/encoding"

	client := resty.New()
	client.SetDebug(true)

	//location like /mnt/vackstage/vackstageBucket/upload/splash_video.mp4
	//creating path for bucket mount, needed for encoding api to search for the video on the bucket folder
	filePathBucket := "/mnt/" + createVideoRequest.Customer + "/" + createVideoRequest.Customer + createVideoRequest.Filename

	response, errz := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&UploadEncodingResponse{}).
		SetBody(&uploadEncodingRequest{
			Jobid:           "0",
			SrcVideo:        filePathBucket,
			APIKey:          apikeyInfo.APIKey,
			Profiles:        encodingProfiles,
			ThumbnailNumber: 4,
			CustomerID:      fmt.Sprint(createVideoRequest.CustomerId)}).
		Post(encodingApiUrl)
	if errz != nil {
		logrus.Errorf(errz.Error())
		return errz
	}
	responseEncoding := response.Result().(*UploadEncodingResponse)

	if !response.IsSuccess() {
		logrus.Errorf("error during call encoding restreamer api, cause: %s", string(*responseEncoding.Message))
		return fmt.Errorf("error during call encoding restreamer api, cause: %s", string(*responseEncoding.Message))
	}

	if *responseEncoding.Result == "KO" {
		logrus.Errorf("error during call encoding restreamer api, cause: %v", responseEncoding)
		return fmt.Errorf("error during call encoding restreamer api, cause: %v", responseEncoding)
	}

	//TODO

	return c.Status(http.StatusCreated).JSON("OK")

}

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

// func GetEncodingSettings(c *fiber.Ctx) error {

// }

// func UpdateEncodingSettings(c *fiber.Ctx) error {

// }
