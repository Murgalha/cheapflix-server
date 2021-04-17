package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const ADDRESS = "localhost"
const PORT = "3000"

func setMoviesURL(movies []Movie) {
	Domain := strings.Join([]string{ADDRESS, PORT}, ":")
	for i := 0; i < len(movies); i++ {
		BaseURL := strings.Join(
			[]string{Domain, "movie", string(movies[i].ID)}, "/")
		movies[i].SubURL = strings.Join([]string{BaseURL, "sub"}, "/")
		movies[i].MovieURL = strings.Join([]string{BaseURL, "movie"}, "/")
		movies[i].PosterURL = strings.Join([]string{BaseURL, "poster"}, "/")
	}
}

func findMovie(id uint64, movies []Movie) *Movie {
	for _, element := range movies {
		if element.ID == id {
			return &element
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments. Please provide 1")
	}
	path := os.Args[1]
	rand.Seed(time.Now().UnixNano())

	movies := GetAllMovies(path)
	setMoviesURL(movies)
	fmt.Println(movies)
	app := fiber.New()

	app.Get("/movie/:id", func(c *fiber.Ctx) error {
		// return everything if 'id' is "all"
		if c.Params("id") == "all" {
			return c.JSON(fiber.Map{
				"result": "ok",
				"error":  "",
				"data":   movies,
			})
		}
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid ID",
				"data":   nil,
			})
		}

		mov := findMovie(id, movies)
		if mov == nil {
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  "Movie not found",
				"data":   nil,
			})
		}

		return c.JSON(fiber.Map{
			"status": "ok",
			"error":  "",
			"data":   mov,
		})
	})

	app.Get("/movie/:id/:data", func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		data := c.Params("data")

		if err != nil {
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid ID",
				"data":   nil,
			})
		}

		mov := findMovie(id, movies)
		if mov == nil {
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  "Movie not found",
				"data":   nil,
			})
		}

		if data == "sub" {
			return c.SendFile(mov.SubPath)
		} else if data == "movie" {
			return c.SendFile(mov.MoviePath)
		} else if data == "poster" {
			return c.SendFile(mov.PosterPath)
		} else {
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid data requested",
				"data":   nil,
			})
		}
		return c.JSON(fiber.Map{
			"status": "ok",
			"error":  "",
			"data":   nil,
		})
	})

	app.Listen(strings.Join([]string{":", PORT}, ""))
}
