package handler

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	CarRepository "hvalfangst/golang-fasthttp-pgx/src/repository/car"
	"strconv"
)

func GetCarById(pool *pgxpool.Pool) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		id, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			_, err := ctx.WriteString(`{"error": "Invalid Car ID"}`)
			if err != nil {
				return
			}
			return
		}

		// Acquire a connection from the pool
		conn, err := pool.Acquire(ctx)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(`{"error": "Failed to acquire database connection"}`)
			return
		}
		defer conn.Release()

		// Query car by ID
		car, err := CarRepository.GetCarByID(ctx, conn, id)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(`{"error": "CarDetails doesn't exist"}`)
			return
		}

		// Serialize the car object to JSON
		carJSON, err := json.Marshal(car)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(`{"error": "Internal Server Error"}`)
			return
		}

		ctx.SetContentType("application/json")
		ctx.SetBody(carJSON)
	}
}
