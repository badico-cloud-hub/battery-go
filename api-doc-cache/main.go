package main

import (
	"fmt"
	"log"
	"math/rand"

	batterygo "github.com/badico-cloud-hub/battery-go/battery"
	"github.com/badico-cloud-hub/battery-go/storages"
	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func updateBatteryStorage() []batterygo.BatteryArgument {
	Key := "foo"
	Value := randSeq(10)
	fmt.Println("update storage: ", Value)

	b := batterygo.BatteryArgument{
		Key,
		Value,
	}

	return []batterygo.BatteryArgument{b}
}
func main() {
	app := fiber.New()

	storage := storages.New()
	battery := batterygo.NewBattery(storage, 3)
	go battery.Init(updateBatteryStorage)

	// GET /api/register
	// app.Get("/api/*", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("✋ %s", c.Params("*"))
	// 	return c.SendString(msg) // => ✋ register
	// })

	// // GET /flights/LAX-SFO
	// app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
	// 	return c.SendString(msg) // => 💸 From: LAX, To: SFO
	// })

	// // GET /dictionary.txt
	// app.Get("/:file.:ext", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
	// 	return c.SendString(msg) // => 📃 dictionary.txt
	// })

	// // GET /john/75
	// app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
	// 	return c.SendString(msg) // => 👴 john is 75 years old
	// })

	// GET /john
	app.Get("read/:name", func(c *fiber.Ctx) error {
		fmt.Println("batteryTestRead")
		value:= battery.Get(c.Params("name"))
		msg := fmt.Sprintf("Hello, %s 👋! Your value is %s", c.Params("name"), value)
		return c.SendString(msg) // => Hello john 👋!
	})
	app.Get("write/:name/:value", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("name"), c.Params("value"))
		arg := batterygo.BatteryArgument{
			Key:   c.Params("name"),
			Value: c.Params("value"),
		}
		fmt.Println("storage", storage)
		battery.Dispatch <- []batterygo.BatteryArgument{arg}
		msg := fmt.Sprintf("Hello, %s 👋! Your value is %s", c.Params("name"), c.Params("value"))
		return c.SendString(msg) // => Hello john 👋!
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("storage", storage)

		msg := fmt.Sprintf("Hello", c.Params("name"), c.Params("value"))
		return c.SendString(msg) // => Hello john 👋!
	})

	log.Fatal(app.Listen(":3001"))
}
