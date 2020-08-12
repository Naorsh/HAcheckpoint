package db

import (
	"database/sql"

	"checkpointapp/internal/pkg/models"
	"checkpointapp/internal/pkg/security"

	//getting pq for extending base sql library
	_ "github.com/lib/pq"
)

var db *sql.DB

// Connect - connection to db
func Connect(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return db.Ping()
}

// Disconnect - disconecting from db
func Disconnect() {
	db.Close()
}

// InsertNewApp inserts new app to DB
func InsertNewApp(app models.App) error {
	encKey, encErr := security.Encrypt(app.Key)
	if encErr != nil {
		return encErr
	}
	_, err := db.Exec(`INSERT INTO public."Applications" (id, name, key, creation_time) VALUES ($1, $2, $3, $4)`, app.ID, app.Name, encKey, app.CreationTime)
	if err != nil {
		return err
	}
	return nil
}

// GetApplications gets all applications from db
// in case error in one row, the function will return all valid rows
// to that point
func GetApplications() ([]models.App, error) {
	var apps []models.App
	rows, err := db.Query(`SELECT * FROM public."Applications" ORDER BY name ASC`)
	if err != nil {
		return apps, err
	}
	defer rows.Close()

	for rows.Next() {
		app := models.App{}
		err := rows.Scan(&app.ID, &app.Name, &app.Key, &app.CreationTime)
		if err != nil {
			return apps, err
		}
		appKey, decErr := security.Decrypt(app.Key)
		if decErr != nil {
			return apps, decErr
		}
		app.Key = appKey
		apps = append(apps, app)
	}
	if err = rows.Err(); err != nil {
		return apps, err
	}

	return apps, err
}
