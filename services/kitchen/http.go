package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
	"github.com/simplemaxr/kitchen/services/common/util"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerId: "001",
			ProductId:  "001",
			Quantity:   1,
		})
		if err != nil {
			util.WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err := c.GetOrder(ctx, &orders.GetOrderRequest{
			CustomerId: "001",
		})
		if err != nil {
			util.WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatal(err)
		}
	})

	log.Println("Starting http server on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderId}}</td>
            <td>{{.CustomerId}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
