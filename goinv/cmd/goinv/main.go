// main.go

package main

import (
	"goinv"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var inventory goinv.Inventory

func main() {
	var err error
	inventory, err = goinv.NewGormInventory()
	if err != nil {
		log.Fatal("Failed to initialize inventory:", err)
	}

	r := gin.Default()

	r.POST("/item", func(c *gin.Context) {
		var item goinv.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			if err := inventory.CreateItem(item); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": "item created"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		}
	})

	r.GET("/items", func(c *gin.Context) {
		items, err := inventory.GetItems()
		if err == nil {
			c.JSON(http.StatusOK, items)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		}
	})

	r.PUT("/item/:id", func(c *gin.Context) {
		var newItem goinv.Item
		id := c.Param("id")
		uintID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		if err := c.ShouldBindJSON(&newItem); err == nil {
			if err := inventory.UpdateItem(uint(uintID), newItem); err == nil {
				c.JSON(http.StatusOK, gin.H{"status": "item updated"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		}
	})

	r.DELETE("/item/:id", func(c *gin.Context) {
		id := c.Param("id")
		uintID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "invalid id"})
			return
		}

		if err := inventory.DeleteItem(uint(uintID)); err == nil {
			c.JSON(http.StatusOK, gin.H{"status": "item deleted"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		}
	})

	r.GET("/items/category/:category", func(c *gin.Context) {
		category := c.Param("category")
		items, err := inventory.GetItemsByCategory(category)
		if err == nil {
			c.JSON(http.StatusOK, items)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
