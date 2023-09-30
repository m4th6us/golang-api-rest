package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}


func main() {
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.DELETE("/cars", sendDelete)
	e.Logger.Fatal(e.Start(":8080"))
	err := getCars
	println(err)
	
}

func getCars(c echo.Context) error {
	var cars []Car
	db, err := sql.Open("sqlite3", "cars.db")


	if err != nil {
		return err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM cars")

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var car Car
		err := rows.Scan(&car.Name, &car.Price)

		if err != nil {
			return nil
		}

		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return c.JSON(200, cars)
}

func createCars(c echo.Context) error {
	var cars []Car

	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")

	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES ($1, $2)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}

	return nil
}

func sendDelete(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	deleteCar(*car)
	return c.JSON(200, "Delete car " + car.Name)
}

func deleteCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")

	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM cars WHERE name = ($1)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name)

	if err != nil {
		return nil
	}

	return nil

}