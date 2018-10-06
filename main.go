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

type Merchant struct {
	Name string `json:"name"`
}

type Transaction struct {
	ID string `json:"id"`
	Description string `json:"description"`
	MerchantName Merchant `json:"merchant"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

func checkError(err error){
	if (err != nil){
		panic(err)
	}
}

func fetch_config() (config *Config, err error) {
	file, err := ioutil.ReadFile("./config.json")
	checkError(err)

	var configData *Config;
	err = json.Unmarshal(file, &configData)
	checkError(err)

	return configData, nil
}

func fetch_transactions() (transactions, err error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.monzo.com//transactions?expand[]=merchant&account_id="+AccountID, nil)
	checkError(err)

	req.Header.Set("Authorization", AuthToken)

	res, err := client.Do(req)
	checkError(err)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	checkError(err)

	var content *Transactions;
	err = json.Unmarshal(body, &content)

	//fmt.Println(string(body))
	fmt.Println(content)

	return nil, nil
}

func main() {
	// Fetch config and set auth token globally
	config, err := fetch_config()
	checkError(err)

	AuthToken = "Bearer " + config.AuthToken
	AccountID = config.AccountID

	fmt.Println(AuthToken)

	// go and fetch transaction data

	fetch_transactions();
}
