package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Paradox570Y1/facility-order-service-v1/internal/dto"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/models"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/repository"
)

var (
	ErrFacilityNotFound = errors.New("facility not found")
	ErrInvalidCode      = errors.New("facility code cannot be empty")
	ErrInvalidName      = errors.New("facility name cannot be empty")
	ErrInvalidAddress   = errors.New("facility address cannot be empty")
)

func IsFacilityNotFound(err error) bool {
	return errors.Is(err, ErrFacilityNotFound)
}

func IsInvalidCode(err error) bool {
	return errors.Is(err, ErrInvalidCode)
}

type FacilityService interface {
	GetAll(ctx context.Context) ([]dto.FacilityResponse, error)
	GetByCode(ctx context.Context, code string) (*dto.FacilityResponse, error)
	Create(ctx context.Context, req dto.CreateFacilityRequest) error
}

type facilityService struct {
	facilityRepo repository.FacilityRepository
}

func NewFacilityService(repo repository.FacilityRepository) FacilityService {
	return &facilityService{
		facilityRepo: repo,
	}
}

func (s *facilityService) GetAll(ctx context.Context) ([]dto.FacilityResponse, error) {
	facilities, err := s.facilityRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.FacilityResponse, 0, len(facilities))

	for _, f := range facilities {
		res = append(res, dto.FacilityResponse{
			Code:    f.Code,
			Name:    f.Name,
			Address: f.Address,
		})
	}

	return res, nil
}

func (s *facilityService) GetByCode(ctx context.Context, code string) (*dto.FacilityResponse, error) {
	if strings.TrimSpace(code) == "" {
		return nil, ErrInvalidCode
	}

	f, err := s.facilityRepo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if f == nil {
		return nil, ErrFacilityNotFound
	}

	return &dto.FacilityResponse{
		Code:    f.Code,
		Name:    f.Name,
		Address: f.Address,
	}, nil
}

func (s *facilityService) Create(ctx context.Context, req dto.CreateFacilityRequest) error {

	if strings.TrimSpace(req.Code) == "" {
		return ErrInvalidCode
	}
	if strings.TrimSpace(req.Name) == "" {
		return ErrInvalidName
	}
	if strings.TrimSpace(req.Address) == "" {
		return ErrInvalidAddress
	}

	facility := models.Facility{
		Code:    req.Code,
		Name:    req.Name,
		Address: req.Address,
	}

	return s.facilityRepo.Create(ctx, facility)
}
