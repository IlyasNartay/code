package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Product represents a product entity
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// PostgreSQLDB struct holds a reference to the SQL database
type PostgreSQLDB struct {
	DB *sqlx.DB
}

// NewPostgreSQLDB creates a new instance of PostgreSQLDB
func NewPostgreSQLDB(connectionString string) (*PostgreSQLDB, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreSQLDB{
		DB: db,
	}, nil
}

// GetProductByID retrieves product from database by ID
func (p *PostgreSQLDB) GetProductByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	err := p.DB.GetContext(ctx, &product, "SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
