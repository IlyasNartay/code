package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Handler struct holds references to Redis client and PostgreSQL database
type Handler struct {
	RedisClient *redis.Client
	DB          *sqlx.DB
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// NewHandler creates a new instance of Handler
func NewHandler(redisClient *redis.Client, db *sqlx.DB) *Handler {
	return &Handler{
		RedisClient: redisClient,
		DB:          db,
	}
}

// GetProduct retrieves a product by its ID
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from request URL
	vars := mux.Vars(r)
	productID := vars["id"]

	// Check Redis cache for product
	product, err := h.getProductFromCache(productID)
	if err != nil {
		http.Error(w, "Error retrieving product from cache", http.StatusInternalServerError)
		return
	}

	if product != nil {
		// Product found in cache, return it
		respondWithJSON(w, http.StatusOK, product)
		return
	}

	// Product not found in cache, retrieve from database
	product, err = h.getProductFromDB(productID)
	if err != nil {
		http.Error(w, "Error retrieving product from database", http.StatusInternalServerError)
		return
	}

	if product == nil {
		// Product not found in database
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Cache product in Redis
	err = h.cacheProduct(productID, product)
	if err != nil {
		fmt.Println("Warning: Error caching product in Redis:", err)
	}

	// Return product to client
	respondWithJSON(w, http.StatusOK, product)
}

func (h *Handler) getProductFromCache(productID string) (*Product, error) {
	// Implement logic to retrieve product from Redis cache
	// Return nil if product not found in cache
	return nil, nil
}

func (h *Handler) getProductFromDB(productID string) (*Product, error) {
	// Implement logic to retrieve product from PostgreSQL database
	// Return nil if product not found in database
	return nil, nil
}

func (h *Handler) cacheProduct(productID string, product *Product) error {
	// Implement logic to cache product in Redis
	return nil
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
