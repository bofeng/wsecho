package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("simple websocket error server")
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		return c.Next()
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		clientID := strconv.FormatUint(uint64(time.Now().UnixMicro()), 36)
		log.Println("client connected:", clientID)
		defer func() {
			c.Close()
			log.Println("client disconnected:", clientID)
		}()

		var (
			mt  int
			msg []byte
			err error
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				break
			}
			if err = c.WriteMessage(mt, msg); err != nil {
				break
			}
		}

		if websocket.IsUnexpectedCloseError(
			err,
			websocket.CloseGoingAway,
			websocket.CloseNoStatusReceived) {
			log.Println("ws error:", err)
		}
	}))
	log.Fatal(app.Listen(":6001"))
}
