package repository

import (
	"context"
	"database/sql"
	"github.com/Paradox570Y1/facility-order-service-v1/internal/models"
)

type FacilityRepository interface {
	GetAll(ctx context.Context) ([]models.Facility, error)
	GetByCode(ctx context.Context, code string) (*models.Facility, error)
	Create(ctx context.Context, facility models.Facility) error
}

type facilityRepository struct {
	db *sql.DB
}

func NewFacilityRepository(database *sql.DB) FacilityRepository {
	return &facilityRepository{
		db: database,
	}
}

func (r *facilityRepository) GetAll(ctx context.Context) ([]models.Facility, error) {
	query := `SELECT code, name, address FROM facilities`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	facilities := []models.Facility{}
	for rows.Next() {
		var f models.Facility
		err := rows.Scan(&f.Code, &f.Name, &f.Address)
		if err != nil {
			return nil, err
		}
		facilities = append(facilities, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return facilities, nil
}

func (r *facilityRepository) GetByCode(ctx context.Context, code string) (*models.Facility, error) {
	query := `SELECT code, name, address FROM facilities WHERE code=?`
	var f models.Facility
	err := r.db.QueryRowContext(ctx, query, code).Scan(&f.Code, &f.Name, &f.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &f, nil
}

func (r *facilityRepository) Create(ctx context.Context, facility models.Facility) error {
	query := `INSERT INTO facilities (code,name,address) VALUES(?,?,?)`
	_, err := r.db.ExecContext(ctx, query,
		facility.Code,
		facility.Name,
		facility.Address,
	)

	return err
}
