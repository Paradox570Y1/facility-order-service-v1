package repository

import (
	"context"
	"database/sql"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/models"
)

type OrderRepository interface {
	GetAll(ctx context.Context) ([]models.Order, error)
	GetByID(ctx context.Context, id string) (*models.Order, error)
	Create(ctx context.Context, order models.Order) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(database *sql.DB) OrderRepository {
	return &orderRepository{
		db: database,
	}
}

func (r *orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	query := `SELECT id, facility_code, status, created_at FROM orders`
	rows,err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	
	orders := []models.Order{}
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.FacilityCode, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (*models.Order, error) {
	query := `SELECT id, facility_code, status, created_at FROM orders WHERE id=?`
	var o models.Order
	err := r.db.QueryRowContext(ctx, query, id).Scan(&o.ID, &o.FacilityCode, &o.Status, &o.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &o, nil
}

func (r *orderRepository) Create(ctx context.Context, order models.Order) error {
	query := `INSERT INTO orders (id,facility_code,status,created_at) VALUES(?,?,?,?)`
	_, err := r.db.ExecContext(ctx,query,
		order.ID,
		order.FacilityCode,
		order.Status,
		order.CreatedAt,
	)
	return err
}