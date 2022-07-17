package restfulpractice

// unmarshal jsondan cevirir
// marshal jsona cevirir

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	Price       float64 `json:"productPrice"`
}



//show product
func ShowProducts() {
	response, err := http.Get("http://localhost:3000/vehicles")

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var vehicles []Product
	json.Unmarshal(bodyBytes, &vehicles)

	fmt.Println(vehicles)
}

// add product

func AddProduct() {

	product := Product{Id: 4, ProductName: "Truck", CategoryId: 1, Price: 17000}

	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/vehicles", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

}
