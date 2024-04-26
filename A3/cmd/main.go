package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func main() {
	redisClient := initRedisClient()

	db := initDB()

	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", getProduct(redisClient, db)).Methods("GET")

	port := ":5432"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func initRedisClient() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5432",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")
	return rdb
}

func initDB() *sqlx.DB {

	db, err := sqlx.Open("postgres", "postgres://postgres:9999@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to database")
	return db
}

func getProduct(redisClient *redis.Client, db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		productID := vars["id"]

		product, err := getProductFromCache(redisClient, productID)
		if err != nil {
			http.Error(w, "Error retrieving product from cache", http.StatusInternalServerError)
			return
		}

		if product != nil {
			respondWithJSON(w, http.StatusOK, product)
			return
		}

		product, err = getProductFromDB(db, productID)
		if err != nil {
			http.Error(w, "Error retrieving product from database", http.StatusInternalServerError)
			return
		}

		if product == nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		err = cacheProduct(redisClient, productID, product)
		if err != nil {
			fmt.Println("Warning: Error caching product in Redis:", err)
		}

		respondWithJSON(w, http.StatusOK, product)
	}
}

func getProductFromCache(redisClient *redis.Client, productID string) (*Product, error) {

	return nil, nil
}

func getProductFromDB(db *sqlx.DB, productID string) (*Product, error) {
	return nil, nil
}

func cacheProduct(redisClient *redis.Client, productID string, product *Product) error {

	return nil
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
}
