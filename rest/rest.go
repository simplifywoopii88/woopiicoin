package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/simplifywoopii88/woopiicoin/blockchain"
	"github.com/simplifywoopii88/woopiicoin/utils"
)

var port string
type url string

func (u url) MarshalText() ([]byte, error){
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func (u urlDescription) String() string {
	return "Hello I'm the URL Description"
}

type urlDescription struct {
	URL url `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}
func documentation(rw http.ResponseWriter, r *http.Request){
	data := []urlDescription{
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: url("/blocks"),
			Method: "GET",
			Description: "See All Blocks",
		},
		{
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add a Block",
			Payload: "data:string",
		},
		{
			URL: url("/blocks/{height}"),
			Method: "GET",
			Description: "See a Block",
		},
	}
	
	
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
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}


func block(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	}else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int){
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(
		http.ListenAndServe(port, router),
	)
}