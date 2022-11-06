package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zhashkevych/orderbook/orderbook"
)

func (a *App) createAsk(c *gin.Context) {
	order := orderbook.Order{
		ID:        uuid.New(),
		OrderType: orderbook.TYPE_ASK,
		CreatedAt: time.Now(),
	}

	if err := c.BindJSON(&order); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	a.orderBook.InsertOrder(&order)

	c.JSON(http.StatusOK, order)
}

func (a *App) createBid(c *gin.Context) {
	order := orderbook.Order{
		ID:        uuid.New(),
		OrderType: orderbook.TYPE_BID,
		CreatedAt: time.Now(),
	}

	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	a.orderBook.InsertOrder(&order)

	c.JSON(http.StatusOK, order)
}

func (a *App) getAllOrders(c *gin.Context) {
	resp := a.orderBook.GetResponse()

	c.JSON(http.StatusOK, resp)
}
