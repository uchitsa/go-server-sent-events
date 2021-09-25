package go_server_sent_events

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Dashboard struct {
	User uint
}

type Client struct {
	name   string
	events chan *Dashboard
}

func main() {
	app := fiber.New()
	app.Get("/sse", adaptor.HTTPHandler(handler(dashboardHandler)))
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: write dashboard handler
	client := &Client{
		name:   r.RemoteAddr,
		events: make(chan *Dashboard, 10),
	}
	go updateDashboard(client)
}

func updateDashboard(client *Client) {

}
