package repository

import (
	"context"
	"log"
	"testing"

	"github.com/devanshg18/go-rest/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// newMongoClient creates and returns a new MongoDB client
func newMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://main-user:devansh123@rest-api.gbsvkcx.mongodb.net/?retryWrites=true&w=majority&appName=Rest-API")

	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	log.Println("MongoDB Successfully connected")

	return mongoClient, nil
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient, err := newMongoClient()
	if err != nil {
		t.Fatalf("Failed to create MongoDB client: %v", err)
	}
	defer func() {
		if err := mongoTestClient.Disconnect(context.Background()); err != nil {
			t.Errorf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	emp1 := uuid.New().String()

	coll := mongoTestClient.Database("CompanyDB").Collection("employee_test")
	empRepo := EmployeeRepo{MongoCollection: coll}

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := models.Employee{
			Name:       "Abhinav Sharma",
			Department: "Sales",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(context.Background(), &emp)
		if err != nil {
			t.Fatalf("Insert operation failed: %v", err)
		}
		t.Logf("Insert operation successful: %v", result.InsertedID)
	})

	t.Run("Insert Employee 2", func(t *testing.T) {
		emp := models.Employee{
			Name:       "Dev Kumar",
			Department: "IT",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(context.Background(), &emp)
		if err != nil {
			t.Fatalf("Insert operation failed: %v", err)
		}
		t.Logf("Insert operation successful: %v", result.InsertedID)
	})
	t.Run("Get Employee 1 ", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatalf("Get Operation Failed : %v", err)
		}
		t.Logf("Emp 1 is:%v ", result.Name)
	})

	//Get All Employee
	t.Run("Get All Employee", func(t *testing.T) {
		results, err := empRepo.FindAllEmployees()
		if err != nil {
			t.Fatal("Get operation Failed", err)
		}
		t.Log("Employee", results)
	})

	t.Run("Update Employee 1 Name", func(t *testing.T) {
		emp := models.Employee{
			Name:       "Dev Kumar as Devashsih",
			Department: "IT",
			EmployeeID: emp1,
		}

		updatedEmp, err := empRepo.UpdateEmployeeByID(emp1, &emp)
		if err != nil {
			log.Fatal("Update operation failed", err)
		}
		t.Log("Update Count", updatedEmp)

	})
	//Delete Employee 1
	t.Run("Delete Employee 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			log.Fatal("Delete Operation failed", err)
		}
		t.Log("Delete Count ", result)

	})
	t.Run("Get All Employee", func(t *testing.T) {
		results, err := empRepo.FindAllEmployees()
		if err != nil {
			t.Fatal("Get operation Failed", err)
		}
		t.Log("Employee", results)
	})
	t.Run("Delete All Employees for Cleanup", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployees()
		if err != nil {
			t.Fatal("Delete Operation Failed", err)
		}
		t.Log("Deleted Count", result)
	})
}
