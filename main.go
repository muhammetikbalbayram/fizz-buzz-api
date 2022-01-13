package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func getFizzBuzz(c echo.Context) error {
	count := c.Param("count")
	fizzBuzzCount, _ := strconv.Atoi(count)

	var fizzBuzzResult []string

	result := map[string][]string{}

	for i := 1; i <= fizzBuzzCount; i++ {
		if i%15 == 0 {
			fizzBuzzResult = append(fizzBuzzResult, "FizzBuzz")
		} else if i%3 == 0 {
			fizzBuzzResult = append(fizzBuzzResult, "Fizz")
		} else if i%5 == 0 {
			fizzBuzzResult = append(fizzBuzzResult, "Buzz")
		} else {
			fizzBuzzResult = append(fizzBuzzResult, strconv.Itoa(i))
		}
	}

	result["fizzbuzz"] = fizzBuzzResult

	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()

	e.GET("/fizzbuzz/:count", getFizzBuzz)

	if err := e.Start(":5000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
