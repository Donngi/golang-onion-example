package presentation

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Donngi/golang-onion-example/infrastructure"
	"github.com/Donngi/golang-onion-example/usecase"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/book/register", registerBookHandler).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

type RegisterBookBody struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func registerBookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Receive POST/book/register")

	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	ddbClient := dynamodb.NewFromConfig(config)
	repository := infrastructure.NewBookRepositoryImpl(ddbClient)
	registerBookUseCase := usecase.NewRegisterBookUseCase(repository)

	var reqBody RegisterBookBody
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Printf("Failed to decode request body: %v\n", err)
		w.Write([]byte("Invalid request body"))
		return
	}

	book, err := registerBookUseCase.Run(reqBody.Name, reqBody.Author)
	if err != nil {
		log.Fatalf("Failed to decode request body: %v", err)
		w.Write([]byte("Failed to register book"))
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		log.Fatalf("Failed to marshal book: %v", err)
		w.Write([]byte("Successfully registered book"))
	}
	w.Write([]byte("Successfully registered book. Book: " + string(res)))
}
