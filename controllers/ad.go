package controllers

import (
	"dcard-2024-backend-intern-assignment/models"
	"dcard-2024-backend-intern-assignment/repositories"
	"dcard-2024-backend-intern-assignment/services"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAdRequest struct {
	Title   string `json:"title" binding:"required"`
	StartAt string `json:"startAt" binding:"required"`
	EndAt   string `json:"endAt" binding:"required"`

	Conditions *CreateAdConditionRequest `json:"conditions"`
}

type CreateAdConditionRequest struct {
	AgeStart int      `json:"ageStart" binding:"omitempty,numeric,gte=1,lte=100"`
	AgeEnd   int      `json:"ageEnd" binding:"omitempty,numeric,gte=1,lte=100"`
	Gender   string   `json:"gender" binding:"omitempty,oneof=M F"`
	Country  []string `json:"country" binding:"dive,iso3166_1_alpha2"`
	Platform []string `json:"platform" binding:"dive,oneof=android ios web"`
}

type AdController struct {
	adRepository repositories.AdRepository
	adService    services.AdService
}

func validateCheckDuplicateAd(arrayData []string, fieldName string) error {
	occurredString := map[string]bool{}
	for v := range arrayData {
		if occurredString[arrayData[v]] {
			return fmt.Errorf("%s has duplicate value", fieldName)
		}
		occurredString[arrayData[v]] = true
	}

	return nil
}

func validateAdConditionAge(ageStart int, ageEnd int) error {
	if ageStart == 0 && ageEnd == 0 {
		return nil
	}

	if ageStart == 0 {
		return errors.New("ageStart is required")
	}

	if ageEnd == 0 {
		return errors.New("ageEnd is required")
	}

	if ageStart > ageEnd {
		return errors.New("ageStart must be less than or equal to ageEnd")
	}

	return nil
}

func parseCreateAdTimeFormat(datetime string) (string, error) {
	const layout string = "2006-01-02T15:04:05.000Z"

	parsedTime, err := time.Parse(layout, datetime)
	if err != nil {
		return "", err
	}
	return parsedTime.UTC().Format(time.DateTime), nil
}

func validateCreateAd(request *CreateAdRequest) error {
	var err error

	if err = validateAdConditionAge(request.Conditions.AgeStart, request.Conditions.AgeEnd); err != nil {
		return err
	}

	if err = validateCheckDuplicateAd(request.Conditions.Country, "country"); err != nil {
		return err
	}

	if err = validateCheckDuplicateAd(request.Conditions.Platform, "platform"); err != nil {
		return err
	}

	request.StartAt, err = parseCreateAdTimeFormat(request.StartAt)
	if err != nil {
		return err
	}

	request.EndAt, err = parseCreateAdTimeFormat(request.EndAt)
	if err != nil {
		return err
	}

	return nil
}

func NewAdController(adRepository repositories.AdRepository, adService services.AdService) *AdController {
	return &AdController{adRepository, adService}
}

func (c *AdController) CreateAd(context *gin.Context) {
	var request CreateAdRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validateCreateAd(&request); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ad, err := c.adRepository.CreateAd(models.Ad{
		Title:   request.Title,
		StartAt: request.StartAt,
		EndAt:   request.EndAt,
		Conditions: models.AdCondition{
			AgeStart: request.Conditions.AgeStart,
			AgeEnd:   request.Conditions.AgeEnd,
			Gender:   request.Conditions.Gender,
			Country:  request.Conditions.Country,
			Platform: request.Conditions.Platform,
		},
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "CreateAd",
		"ad":      ad,
	})
}
