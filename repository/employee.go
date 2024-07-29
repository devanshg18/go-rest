package repository

import (
	"context"
	"fmt"

	"github.com/devanshg18/go-rest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

// InsertEmployee inserts a new employee into the MongoDB collection
func (r *EmployeeRepo) InsertEmployee(emp *models.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil // Return the inserted ID
}

// FindEmployeeByID retrieves an employee by their ID
func (r *EmployeeRepo) FindEmployeeByID(empID string) (*models.Employee, error) {
	var emp models.Employee

	// Use bson.D{{Key: "employeeID", Value: empID}} to find by employee ID
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "employeeID", Value: empID}}).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

// FindAllEmployees retrieves all employees from the MongoDB collection
func (r *EmployeeRepo) FindAllEmployees() ([]models.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var emps []models.Employee
	err = results.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("error decoding employees: %v", err)
	}
	return emps, nil // Return the slice of employees
}
func (r *EmployeeRepo) UpdateEmployeeByID(empID string, updatedEmp *models.Employee) (int64, error) {
	filter := bson.M{"employeeID": empID}
	update := bson.M{
		"$set": bson.M{
			"name":       updatedEmp.Name,
			"department": updatedEmp.Department,
		},
	}

	result, err := r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	filter := bson.M{"employeeID": empID}
	result, err := r.MongoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployees() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
