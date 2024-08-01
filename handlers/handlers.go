package handlers

import (
	"encoding/csv"
	"net/http"
	"os"

	"go-backend-crm/config"
	"go-backend-crm/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var customers []models.Customer

func LoadCustomers() error {
	file, err := os.Open(config.GetConfig().CSVFile)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	customers = []models.Customer{}
	for _, record := range records[1:] {
		customer := models.Customer{
			ID:      record[0],
			Name:    record[1],
			Email:   record[2],
			Phone:   record[3],
			Address: record[4],
		}
		customers = append(customers, customer)
	}
	return nil
}

func SaveCustomers() error {
	file, err := os.Create(config.GetConfig().CSVFile)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write([]string{"ID", "Name", "Email", "Phone", "Address"})
	for _, customer := range customers {
		w.Write([]string{customer.ID, customer.Name, customer.Email, customer.Phone, customer.Address})
	}
	return nil
}

func GetCustomers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, customers)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	for _, customer := range customers {
		if customer.ID == id {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, customer)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = uuid.New().String()
	customers = append(customers, customer)
	if err := SaveCustomers(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var updatedCustomer models.Customer
	if err := c.BindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			if err := SaveCustomers(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, updatedCustomer)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			if err := SaveCustomers(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
}
