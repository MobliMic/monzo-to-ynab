package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
)

var AuthToken string;

var AccountID string;

type Config struct {
	AuthToken string `json:"access_token"`
	AccountID string `json:"account_id"`
}

type Transaction struct {
	ID string `json:"id"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

func fetch_config() (config *Config, err error) {
	file, err := ioutil.ReadFile("./config.json")
	if (err != nil) {
		return nil, err
	}

	var configData *Config;
	err = json.Unmarshal(file, &configData)
	if (err != nil) {
		panic(err)
	}

	return configData, nil
}

func fetch_transactions() (transactions, err error) {
	client := http.Client{}
	//todo: error handling
	req, err := http.NewRequest("GET", "https://api.monzo.com//transactions?account_id="+AccountID, nil)
	req.Header.Set("Authorization", AuthToken)

	res, err := client.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var content *Transactions;
	err = json.Unmarshal(body, &content)

	//fmt.Println(string(body))
	fmt.Println(content)

	return nil, nil
}

func main() {
	// Fetch config and set auth token globally
	config, err := fetch_config()
	if (err != nil) {
		panic(err)
	}

	AuthToken = "Bearer " + config.AuthToken
	AccountID = config.AccountID

	fmt.Println(AuthToken)

	// go and fetch transaction data

	fetch_transactions();
}
