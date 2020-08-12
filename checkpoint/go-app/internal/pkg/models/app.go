package models

import (
	"time"

	"github.com/google/uuid"
)

// App struct
type App struct {
	ID           string `json:"ID"`
	Name         string `json:"Name"`
	Key          string `json:"Key"`
	CreationTime string `json:"CreationTime"`
}

// HealthResponse struct
type HealthResponse struct {
	Key string `json:"Key"`
}

// NewApp creats New App struct
func NewApp(name, key string) App {
	return App{
		ID:           uuid.New().String(),
		Name:         name,
		Key:          key,
		CreationTime: time.Now().Format("2006.01.02 15:04:05"),
	}

}
