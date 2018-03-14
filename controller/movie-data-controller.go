package controller

import (
	"Movies/dbs"
	"Movies/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	_ "github.com/lib/pq" //for postgres driver
)

//GetData by
func GetData(c echo.Context) (err error) {
	data := strings.Split(c.QueryString(), "=")
	key := data[0]

	if key == "year" {
		value := c.QueryParam("year")
		result, err := models.GetDatabyYear(dbs.DB, value)
		fmt.Println("return result", result)
		if err != nil {
			fmt.Println(err.Error())
		}
		Finalvalue := getvalues(result)
		if len(Finalvalue) > 0 {
			return c.JSON(http.StatusOK, Finalvalue)
		}
	} else if key == "genres" {
		value := c.QueryParam("genres")
		result, err := models.GetDatabyGenres(dbs.DB, value)
		fmt.Println("return result", result)
		if err != nil {
			fmt.Println(err.Error())
		}
		Finalvalue := getvalues(result)
		if len(Finalvalue) > 0 {
			return c.JSON(http.StatusOK, Finalvalue)
		}
	}
	return c.JSON(http.StatusNotFound, "Data Not Found")
}

//GetDataTitle = if found print otherwise add it
func GetDataTitle(c echo.Context) (err error) {

	data := strings.Split(c.QueryString(), "=")
	key := data[0]
	if key == "title" {
		value := c.QueryParam("title")
		result, err := models.GetDatabyTitle(dbs.DB, value)
		fmt.Println("return result", result)
		if err != nil {
			fmt.Println(err.Error())
		}
		Finalvalue := getvalues(result)
		if len(Finalvalue) == 0 {
			newresult := models.GetMovieInfo(value)
			fmt.Println("new result:", newresult)
			dataResult, err := models.InsertData(dbs.DB, newresult)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("data result:", dataResult)
			if result != nil {
				return c.JSON(http.StatusCreated, newresult)
			}
			return c.JSON(http.StatusNotFound, "data is not created")
		}
		return c.JSON(http.StatusOK, Finalvalue)
	}
	return c.JSON(http.StatusNotFound, "Data Not Found")
}

//GetDatainRange by
func GetDatainRange(c echo.Context) (err error) {
	data := strings.Split(c.QueryString(), "&")
	data1 := strings.Split(data[0], "=")
	data2 := strings.Split(data[1], "=")
	key1 := data1[0]
	key2 := data2[0]

	if key1 == "year1" && key2 == "year2" {
		value1 := c.QueryParam("year1")
		value2 := c.QueryParam("year2")
		result, err := models.GetDatabyYearRange(dbs.DB, value1, value2)
		fmt.Println("return result", result)
		if err != nil {
			fmt.Println(err.Error())
		}
		Finalvalue := getvalues(result)
		if len(Finalvalue) > 0 {
			return c.JSON(http.StatusOK, Finalvalue)
		}
	} else if key1 == "low" && key2 == "high" {
		value1 := stringtofloat(c.QueryParam("low"))
		value2 := stringtofloat(c.QueryParam("high"))
		result, err := models.GetDatabyRating(dbs.DB, value1, value2)
		fmt.Println("return result", result)
		if err != nil {
			fmt.Println(err.Error())
		}
		Finalvalue := getvalues(result)
		if len(Finalvalue) > 0 {
			return c.JSON(http.StatusOK, Finalvalue)
		}
	}
	return c.JSON(http.StatusNotFound, "Data Not Found")
}

//UpdateData update data
func UpdateData(c echo.Context) (err error) {
	data := &models.Movie{
		Id:     c.FormValue("id"),
		Rating: stringtofloat(c.FormValue("rating")),
		Genres: c.FormValue("genres"),
	}
	if err = c.Bind(data); err != nil {
		return err
	}
	result, err := models.UpdateData(dbs.DB, *data)
	if err != nil {
		fmt.Println("Error is", err)
	}
	fmt.Println("result:", result)
	if result != nil {
		return c.JSON(http.StatusOK, "updated")
	}
	return c.JSON(http.StatusNotFound, "Not updated")
}

func getvalues(sql *sql.Rows) []models.Movie {
	alldata := make([]models.Movie, 0.0)
	data := models.Movie{}
	for sql.Next() {
		sql.Scan(&data.Id, &data.Title, &data.Release_year, &data.Rating, &data.Genres)
		fmt.Println(data.Id, data.Title, data.Release_year, data.Rating, data.Genres)
		alldata = append(alldata, data)
		fmt.Println("alldata:", alldata)
	}
	return alldata
}

func stringtofloat(s string) float64 {
	str, _ := strconv.ParseFloat(s, 64)
	return str
}
