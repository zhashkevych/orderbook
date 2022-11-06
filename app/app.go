package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/orderbook/orderbook"
)

// Real Time Orderbook w/ Websockets

type App struct {
	server *http.Server

	orderBook *orderbook.OrderBook
}

func NewApp() *App {
	return &App{
		server: &http.Server{
			Addr: ":8080",
		},
		orderBook: orderbook.NewOrderBook(),
	}
}

func (a *App) Init() error {
	a.initRouters()

	return a.server.ListenAndServe()
}

func (a *App) initRouters() {
	r := gin.Default()

	orders := r.Group("/orders")
	{
		orders.GET("/", a.getAllOrders)
		orders.GET("/:id", nil)
		orders.POST("/ask", a.createAsk)
		orders.POST("/bid", a.createBid)
		orders.PUT("/:id", nil)
		orders.DELETE("/:id", nil)
	}

	a.server.Handler = r
}
