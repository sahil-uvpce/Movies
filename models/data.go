package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/shiggins8/go-imdb"
)

//Movie struct for Movie
type Movie struct {
	Id           string  `json:"id"`
	Title        string  `json:"title"`
	Release_year string  `json:"release_year"`
	Rating       float64 `json:"rating"`
	Genres       string  `json:"genres"`
}

//GetMovieInfo = get movie data
func GetMovieInfo(name string) Movie {
	imdb.SetOmdbAPIKey("ce349cca")

	result, err := imdb.FetchMovie(name)

	if err != nil {
		log.Fatal(err)
	}

	rating, err := result.GetRatings("Internet Movie Database")
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	value := strings.Split(rating, "/")
	rate, _ := strconv.ParseFloat(value[0], 64)
	fmt.Println("GNREEEEEEE:", result.Genre)
	movie := Movie{
		Id:           result.ImdbID,
		Title:        result.Title,
		Release_year: result.Year,
		Rating:       rate,
		Genres:       result.Genre,
	}
	fmt.Println("rate:", rate, "Genres:", result.Genre, "Title:", result.Title, "year:", result.Year, "ID:", result.ImdbID)

	return movie
}

//InsertData = get data by title
func InsertData(db *sql.DB, movie Movie) (sql.Result, error) {
	genreArray := strings.Fields(movie.Genres)
	fmt.Println(genreArray)

	ins := "INSERT INTO record (id,title,release_year,rating,genres) VALUES ($1, $2, $3, $4, $5)"
	return db.Exec(ins, movie.Id, movie.Title, movie.Release_year, movie.Rating, pq.Array(genreArray))
}

//GetDatabyTitle = get data by title
func GetDatabyTitle(db *sql.DB, title string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select * from record where title in('%s')", title))
}

//GetDatabyYear = get by year
func GetDatabyYear(db *sql.DB, year string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select * from record where release_year='%s'", year))
}

//GetDatabyYearRange = get by year range
func GetDatabyYearRange(db *sql.DB, year1, year2 string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("SELECT * from record WHERE release_year BETWEEN '%s' AND '%s'", year1, year2))
}

//GetDatabyGenres = get by genres
func GetDatabyGenres(db *sql.DB, genres string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("SELECT * from record WHERE genres @>'{%s}'", genres))
}

//GetDatabyRating = get by year rating
func GetDatabyRating(db *sql.DB, low, high float64) (*sql.Rows, error) {
	if low == 0 {
		return db.Query(fmt.Sprintf("SELECT * from record where rating <'%f'", high))
	} else if high == 0 {
		return db.Query(fmt.Sprintf("SELECT * from record where rating > '%f'", low))
	}
	return db.Query(fmt.Sprintf("SELECT * from record where rating > '%f' AND rating <'%f'", low, high))
}

//UpdateData data update
func UpdateData(db *sql.DB, movie Movie) (sql.Result, error) {
	genreArray := strings.Fields(movie.Genres)
	updateQuery := "update record set rating=$1, genres = $2 where id= $3"
	return db.Exec(updateQuery, movie.Rating, pq.Array(genreArray), movie.Id)
}
