package services

import (
	"dcard-2024-backend-intern-assignment/models"
	"errors"
)

type AdService struct {
}

func NewAdService() *AdService {
	return &AdService{}
}

func (s *AdService) ValidateCreateAdConditions(conditions models.AdCondition) error {
	if conditions.AgeStart == 0 && conditions.AgeEnd == 0 {
		return nil
	}

	if conditions.AgeStart == 0 {
		return errors.New("ageStart is required")
	}

	if conditions.AgeEnd == 0 {
		return errors.New("ageEnd is required")
	}

	if conditions.AgeStart > conditions.AgeEnd {
		return errors.New("ageStart must be less than or equal to ageEnd")
	}

	return nil
}
