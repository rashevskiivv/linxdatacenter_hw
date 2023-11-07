package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

type Product struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

func main() {
	fileType := os.Args[1]
	file := os.Args[2]

	var products []Product
	var err error
	if fileType == "csv" {
		products, err = readDataFromCSVFile(file)
	} else {
		err = readDataFromJSONFile(file, &products)
	}
	if err != nil {
		log.Println(err)
	}

	log.Println("самый дорогой продукт", findOutMostExpensiveProduct(products))
	log.Println("с самым высоким рейтингом", findOutHighRatingProduct(products))
}

func readDataFromJSONFile(file string, v interface{}) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			log.Println(err)
		}
	}(in)

	byteValue, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteValue, v)
	return err
}

func readDataFromCSVFile(file string) ([]Product, error) {
	// open file for reading
	in, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			log.Panic(err)
		}
	}(in)

	reader := csv.NewReader(in)

	data := make([]Product, 0)
	for i := 0; ; i++ {
		record, err := reader.Read()
		// skip headers
		if i == 0 {
			continue
		}
		if err != nil {
			// we have read whole file
			if errors.Is(err, io.EOF) {
				return data, nil
			}
			return nil, err
		}

		price, err := strconv.ParseInt(record[1], 10, 0)
		if err != nil {
			return nil, err
		}
		rating, err := strconv.ParseInt(record[2], 10, 0)
		if err != nil {
			return nil, err
		}

		product := Product{
			Product: record[0],
			Price:   int(price),
			Rating:  int(rating),
		}
		data = append(data, product)
	}
}

func findOutMostExpensiveProduct(products []Product) Product {
	maxPrice := -1
	var output Product
	for _, product := range products {
		if product.Price > maxPrice {
			maxPrice = product.Price
			output = product
		}
	}
	return output
}

func findOutHighRatingProduct(products []Product) Product {
	maxRating := -1
	var output Product
	for _, product := range products {
		if product.Rating > maxRating {
			maxRating = product.Rating
			output = product
		}
	}
	return output
}
