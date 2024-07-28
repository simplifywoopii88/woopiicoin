package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/simplifywoopii88/woopiicoin/blockchain"
	"github.com/simplifywoopii88/woopiicoin/utils"
)


const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error){
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func (u URLDescription) String() string {
	return "Hello I'm the URL Description"
}

type URLDescription struct {
	URL URL `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}
func documentation(rw http.ResponseWriter, r *http.Request){
	data := []URLDescription{
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
		{
			URL: URL("/blocks/{id}"),
			Method: "GET",
			Description: "See a Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	
	// struct to json 
	// case 1
	
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)

	// case 2
	json.NewEncoder(rw).Encode(data)

}

func blocks(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {

	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(
		http.ListenAndServe(port, nil),
	)
}