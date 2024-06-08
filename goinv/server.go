package goinv

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func NewServer(inventory Inventory) *gin.Engine {
	r := gin.Default()

	r.GET("/items", func(c *gin.Context) {
		getItems(c, inventory)
	})
	r.POST("/item", func(c *gin.Context) {
		createItem(c, inventory)
	})
	r.PUT("/item/:id", func(c *gin.Context) {
		updateItem(c, inventory)
	})
	r.DELETE("/item/:id", func(c *gin.Context) {
		deleteItem(c, inventory)
	})

	return r
}

func getItems(c *gin.Context, inventory Inventory) {
	category := c.Query("category")
	location := c.Query("location")

	var items []Item
	var err error

	filter := ItemFilter{
		Category: category,
		Location: location,
	}

	log.Info(fmt.Sprintf("Filtering items by category: %s, location: %s", category, location))

	if category != "" || location != "" {
		items, err = inventory.GetItemsWithFilter(filter)
	} else {
		items, err = inventory.GetItems()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context, inventory Inventory) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := inventory.CreateItem(item); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, item)
}

func updateItem(c *gin.Context, inventory Inventory) {
	id := c.Param("id")
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert id to uint
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	if err := inventory.UpdateItem(uint(uid), item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context, inventory Inventory) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64) // base 10, up to 64 bits
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := inventory.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
