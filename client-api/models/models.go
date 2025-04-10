package models

import "time"

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Stock struct {
	Available int `json:"available"`
	Reserved  int `json:"reserved"`
}

type Image struct {
	URL string `json:"url"`
	Alt string `json:"alt"`
}

type Product struct {
	ProductID      string            `json:"product_id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Price          Price             `json:"price"`
	Category       string            `json:"category"`
	Brand          string            `json:"brand"`
	Stock          Stock             `json:"stock"`
	SKU            string            `json:"sku"`
	Tags           []string          `json:"tags"`
	Images         []Image           `json:"images"`
	Specifications map[string]string `json:"specifications"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	Index          string            `json:"index"`
	StoreID        string            `json:"store_id"`
}

type AnalyticsRecommendation struct {
	Hour  int `json:"hour"`
	Count int `json:"count"`
}

type ClientRequest struct {
	Query     string    `json:"query"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}
