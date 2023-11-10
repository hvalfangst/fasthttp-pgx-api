package car

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"hvalfangst/golang-fasthttp-pgx/src/model"
	"log"
)

func GetCarByID(ctx context.Context, conn *pgxpool.Conn, carID int64) (*model.Car, error) {
	// Initialize a car instance.
	car := &model.Car{}

	err := conn.QueryRow(
		ctx,
		"SELECT id, make, model, year, vin, color, purchase_date, owner_id FROM cars WHERE id = $1",
		carID,
	).Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.VIN, &car.Color, &car.PurchaseDate, &car.OwnerID)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Printf("Car with ID %d not found", carID)
		return nil, nil // Return nil and no error for not found
	} else if err != nil {
		log.Printf("Error retrieving car by ID: %v", err)
		return nil, err
	}

	// If no error occurred, return the car.
	return car, nil
}
