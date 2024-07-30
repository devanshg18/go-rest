package usecase

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/devanshg18/go-rest/models"
	"github.com/devanshg18/go-rest/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = "Invalid body request"
		return
	}
	//assign a new employee ID
	emp.EmployeeID = uuid.NewString()
	//creating repo instance
	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	//inserting starts from here
	insertID, err := repo.InsertEmployee(context.Background(), &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insert error", err)
		res.Error = "Failed to Insert Employee"
	}
	//response data
	res.Data = map[string]interface{}{
		"employeeId": emp.EmployeeID, "insertID": insertID,
	}
	w.WriteHeader(http.StatusCreated)

}

func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Finding error", err)
		res.Error = "Failed to Get Employee"
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)

}

func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Finding error", err)
		res.Error = "Failed to Get Employee"
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id")
		res.Error = "invalid employee id"
	}
	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = "Invalid body request"
		return
	}
	emp.EmployeeID = empID
	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = "Invalid body request"
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = "Invalid body request"
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) DeleteAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	count, err := repo.DeleteAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = "Invalid body request"
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
