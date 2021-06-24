package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Order struct {
	OrderUID        string `json:"order_uid"`
	Entry           string `json:"entry"`
	TotalPrice      int    `json:"total_price"`
	CustomerID      string `json:"customer_id"`
	TrackNumber     string `json:"track_number"`
	DeliveryService string `json:"delivery_service"`
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	var data []Order
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	jsonValue, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonValue)
}

func main() {
	http.HandleFunc("/orders", getOrders)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
