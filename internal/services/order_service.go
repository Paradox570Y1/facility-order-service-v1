package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Paradox570Y1/facility-order-service-v1/internal/constants"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/dto"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/models"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/repository"
)

var (
	ErrFacilityDoesNotExist = errors.New("facility does not exist")
	ErrOrderNotFound        = errors.New("order not found")
	ErrInvalidID            = errors.New("order id cannot be empty")
	ErrInvalidFacilityCode  = errors.New("facility code cannot be empty")
)

func IsOrderNotFound(err error) bool {
	return errors.Is(err, ErrOrderNotFound)
}

func IsFacilityDoesNotExist(err error) bool {
	return errors.Is(err, ErrFacilityDoesNotExist)
}

func IsInvalidID(err error) bool {
	return errors.Is(err, ErrInvalidID)
}

type OrderService interface {
	GetAll(ctx context.Context) ([]dto.OrderResponse, error)
	GetByID(ctx context.Context, id string) (*dto.OrderResponse, error)
	Create(ctx context.Context, order dto.CreateOrderRequest) error
}

type orderService struct {
	orderRepo    repository.OrderRepository
	facilityRepo repository.FacilityRepository
}

func NewOrderService(orderRepo repository.OrderRepository, facilityRepo repository.FacilityRepository) OrderService {
	return &orderService{
		orderRepo:    orderRepo,
		facilityRepo: facilityRepo,
	}
}

func toOrderResponse(o models.Order) dto.OrderResponse {
	return dto.OrderResponse{
		ID:           o.ID,
		FacilityCode: o.FacilityCode,
		Status:       o.Status,
		CreatedAt:    o.CreatedAt,
	}
}

func (s *orderService) GetAll(ctx context.Context) ([]dto.OrderResponse, error) {
	res, err := s.orderRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]dto.OrderResponse, 0, len(res))
	for _, o := range res {
		orders = append(orders, toOrderResponse(o))
	}
	return orders, nil
}

func (s *orderService) GetByID(ctx context.Context, id string) (*dto.OrderResponse, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrInvalidID
	}

	o, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, ErrOrderNotFound
	}

	order := toOrderResponse(*o)
	return &order, nil
}
func (s *orderService) Create(ctx context.Context, order dto.CreateOrderRequest) error {

	if strings.TrimSpace(order.ID) == "" {
		return ErrInvalidID
	}
	if strings.TrimSpace(order.FacilityCode) == "" {
		return ErrInvalidFacilityCode
	}

	facility, err := s.facilityRepo.GetByCode(ctx, order.FacilityCode)
	if err != nil {
		return err
	}

	if facility == nil {
		return ErrFacilityDoesNotExist
	}

	newOrder := models.Order{
		ID:           order.ID,
		FacilityCode: order.FacilityCode,
		Status:       constants.OrderStatusCreated,
		CreatedAt:    time.Now(),
	}

	return s.orderRepo.Create(ctx, newOrder)
}
