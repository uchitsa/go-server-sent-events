package go_server_sent_events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
	"time"
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
	app.Listen(":3000")
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type\"", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(1 * time.Second)
	select {
	case events := <-client.events:
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.Encode(events)
		fmt.Fprintf(w, "data: %v\n\n", buf.String())
		fmt.Printf("data: %v\n", buf.String())
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func updateDashboard(client *Client) {
	for {
		db := &Dashboard{
			User: uint(rand.Uint32()),
		}
		client.events <- db
	}
}
