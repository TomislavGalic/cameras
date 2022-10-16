package controllers

import (
	"strconv"

	"github.com/TomislavGalic/cameras/database"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {

	latitude := c.Query("latitude", "")
	longitude := c.Query("longitude", "")
	radius := c.Query("radius", "")

	var errors []string

	if !checkValue(latitude, -90, 90) {
		errors = append(errors, "latitude")
	}

	if !checkValue(longitude, -180, 180) {
		errors = append(errors, "longitude")
	}
	if !checkValue(radius, 0, 81226144) {
		errors = append(errors, "radius")
	}

	if len(errors) > 0 {
		c.Redirect("/", 400)
	}

	type Camera struct {
		LocationName string
		Latitude     float64
		Longitude    float64
	}
	var cameras []Camera

	database.DB.Raw(
		"SELECT location_name, latitude, longitude FROM cameras WHERE 1 = 1 AND 2 * 3961 * asin(sqrt(power((sin(radians((? - cast(cameras.latitude as decimal(10,8))) / 2))) , 2) + cast(cos(radians(cast(cameras.latitude as decimal(18,8)))) * cos(radians(?)) * power((sin(radians((? - cast(cameras.longitude as decimal(18,8))) / 2))) , 2) as decimal(18,10) ))) <= ?", latitude, latitude, longitude, radius).Scan(&cameras)
	return c.JSON(cameras)
}

func checkValue(value string, min float64, max float64) bool {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}

	if v < min || v > max || value == "" {
		return false
	}
	return true
}
