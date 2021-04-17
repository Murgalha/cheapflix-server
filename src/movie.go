package main

import (
	_ "fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Movie struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Year       uint64 `json:"year"`
	SubURL     string `json:"subtitle"`
	MovieURL   string `json:"url"`
	PosterURL  string `json:"poster"`
	SubPath    string `json:"-"`
	MoviePath  string `json:"-"`
	PosterPath string `json:"-"`
}

func isImage(Filename string) bool {
	exts := [...]string{".jpg", ".png"}
	for _, ext := range exts {
		if filepath.Ext(Filename) == ext {
			return true
		}
	}
	return false
}

func isSubtitle(Filename string) bool {
	exts := [...]string{".srt"}
	for _, ext := range exts {
		if filepath.Ext(Filename) == ext {
			return true
		}
	}
	return false
}

func isMovie(Filename string) bool {
	exts := [...]string{".mp4", ".avi", ".mkv"}
	for _, ext := range exts {
		if filepath.Ext(Filename) == ext {
			return true
		}
	}
	return false
}

func getMovie(Path, MovieDirName string) Movie {
	MoviePath := filepath.Join(Path, MovieDirName)

	dir, err := ioutil.ReadDir(MoviePath)
	var movie Movie

	if err != nil {
		log.Fatalf("Could not read movie '%s': %s", MovieDirName, err)
	}

	for _, element := range dir {
		if !element.IsDir() && isMovie(element.Name()) {
			movie.MoviePath = filepath.Join(MoviePath, element.Name())
		} else if !element.IsDir() && isImage(element.Name()) {
			movie.PosterPath = filepath.Join(MoviePath, element.Name())
		} else if !element.IsDir() && isSubtitle(element.Name()) {
			movie.SubPath = filepath.Join(MoviePath, element.Name())
		}
	}
	// I want a 4-digit ID
	// TODO: Check if ID is unique
	movie.ID = (rand.Uint64() % 9000) + 1000

	// extract movie name and year from directory name
	regex := regexp.MustCompile(`^(.+)\s*\((\d+)\)$`)
	matches := regex.FindStringSubmatch(MovieDirName)
	if len(matches) == 3 {
		movie.Name = strings.Trim(matches[1], " ")
		movie.Year, _ = strconv.ParseUint(matches[2], 10, 32)
	} else {
		movie.Name = ""
		movie.Year = 0
	}
	return movie
}

func GetAllMovies(Path string) []Movie {
	moviesDir, err := ioutil.ReadDir(Path)
	var movies []Movie

	if err != nil {
		log.Fatalf("Could not read '%s'", Path)
	}

	for _, element := range moviesDir {
		if element.IsDir() {
			movies = append(movies, getMovie(Path, element.Name()))
		}
	}
	return movies
}
