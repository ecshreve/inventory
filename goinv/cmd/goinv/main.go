//go:build !test

package main

import (
	"goinv"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

var inventory goinv.Inventory

func main() {
	log.Info("Starting goinv")
	os.Setenv("ENV", "prod")

	var err error
	inventory, err = goinv.NewGormInventory()
	if err != nil {
		log.Fatal("Failed to initialize inventory:", err)
	}

	locs := []goinv.StorageLocation{
		{Description: "HalfCrate_White_1", Location: "Office"},
		{Description: "HalfCrate_White_2", Location: "Office"},
		{Description: "FullCrate_Black_1", Location: "Office"},
	}

	items := []goinv.Item{}

	if err := inventory.Populate(items, locs); err != nil {
		log.Fatal("Failed to populate storage locations:", err)
	}

	r := gin.Default()

	r.GET("/items", getItems)
	r.POST("/item", createItem)
	r.PUT("/item/:id", updateItem)
	r.DELETE("/item/:id", deleteItem)
	r.GET("/items/category/:category", getItemsByCategory)
	r.GET("/items/location/:location", getItemsByLocation)

	log.Info("Listening on :8080")
	r.Run()
}

func getItems(c *gin.Context) {
	items, err := inventory.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var newItem goinv.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		return
	}

	if err := inventory.CreateItem(newItem); err != nil { // Assuming CreateItem is a method in inventory
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "item created"})
}

func updateItem(c *gin.Context) {
	var newItem goinv.Item
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
		return
	}

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		return
	}

	if err := inventory.UpdateItem(uint(uintID), newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item updated"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
		return
	}

	if err := inventory.DeleteItem(uint(uintID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item deleted"})
}

func getItemsByCategory(c *gin.Context) {
	category := c.Param("category")
	items, err := inventory.GetItemsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func getItemsByLocation(c *gin.Context) {
	location := c.Param("location")
	locationID, err := strconv.ParseUint(location, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid location"})
		return
	}

	items, err := inventory.GetItemsByStorageLocation(uint(locationID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
