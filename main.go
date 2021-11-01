package main

import (
	"fmt"
	"net/http"
	_"time"
	 "io/ioutil"
	 _"encoding/json"
)

type Users struct{
	Id int `json:"id"`
	Name  string `json:"name"`
	UserName  string `json:"username"`
	Email  string `json:"email"`
	Address  Address `json:"address"`
	Phone  string `json:"phone"`
	Website  string `json:"website"`
	Company  Company `json:"company"`
}

type Address struct{
	Street  string `json:"street"`
	Suite  string `json:"suite"`
	City  string `json:"city"`
	Zipcode  string `json:"zipcode"`
}

type Company struct{
	Name  string `json:"name"`
	CatchPhrase  string `json:"catchPhrase"`
	Bs  string `json:"bs"`
}

func main()  {
	
	//endpoint
	url := "https://jsonplaceholder.typicode.com/users"

	resp, err := http.Get(url)
	if err != nil {
     fmt.Println(err)
    }
  
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
	fmt.Println(err)
   }

   fmt.Println(string(body))
	/**
	//create a client
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	//prepera new request s
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if(err != nil){
		fmt.Println(err)
	}

	result , getErr:= client.Do(req)
	if(getErr != nil){
		fmt.Println(getErr)
	}

	if(result.Body != nil){
		defer result.Body.Close()
	}

	body, _ := ioutil.ReadAll(result.Body)
    
	users := []Users{}

	jsonErr := json.Unmarshal(body, &users)

	if(jsonErr != nil){
		fmt.Println(jsonErr)
	}

	for _, data := range users{
		fmt.Println(data)
	}*/
}	