package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var host string = ""

type CreatePayment struct {
	LinkName    string `json:"link"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Price       int    `json:"price"`
	Token       string `json:"token"`
	Description string `json:"description"`
}

type TransactionParameter struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Price         int    `json:"price"`
	Token         string `json:"token"`
	PaymentLinkID int    `json:"paymentlink"`
}

func createPaymentLinks(client http.Header, linkName string, name string, email string, price int, description string, err error) (string, error) {

	data := CreatePayment{
		LinkName:    linkName,
		Name:        name,
		Email:       email,
		Price:       price,
		Description: description,
	}

	// jsonData, _ := json.Marshal(data)

	// client := &http.Client{}

	// // var endpoint = host
	// response, err := http.Post(host)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	response.Header = client
	if err != nil {
		log.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func fetchUserLinks() {
	response, err := http.Post("")
	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func createPaymentRecord() {
	response, err := http.Post("")
	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func fecthPaymentRecords() {
	response, err := http.Get("")
	http.NewRequest
	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func runTransaction() {
	response, err := http.Post("localhost:5000/paymentslink/generate-paylink")
	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func main() {
}
