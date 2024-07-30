package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/devanshg18/go-rest/usecase"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	//loading of the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env Load Error", err)
	}
	log.Println("Env File is uploaded")

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Connection Error", err)
	}
	mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping Failed", err)
	}

	//error check
	log.Println("Mongo Connection is successfull")
}
func main() {
	//Disconnect the mongodb connection
	defer mongoClient.Disconnect(context.Background())
	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	employeeService := usecase.EmployeeService{MongoCollection: coll}
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	r.HandleFunc("/employee", employeeService.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", employeeService.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee", employeeService.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", employeeService.UpdateEmployeeByID).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", employeeService.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee/{id}", employeeService.DeleteAllEmployees).Methods(http.MethodDelete)

	log.Println("Running Server")
	http.ListenAndServe(":4444", r)
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("runnning"))
}
