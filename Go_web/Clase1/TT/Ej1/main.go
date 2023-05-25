package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var Products []Product

func main() {

	err := ReadDataBase("products.json")
	if err != nil {
		log.Fatal(err)
	}
	server := gin.Default()
	products := server.Group("/products")
	products.GET("/", ProductsList)

}

func Pong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func ReadDataBase(fileRoute string) (err error) {
	file, err := ioutil.ReadFile(fileRoute)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Products)
	if err != nil {
		return err
	}
	return nil
}

func ProductsList(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, Products)
}
