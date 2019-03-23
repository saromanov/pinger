// Package core defines main logic of the pinger
// it contains worker for checking site availability
package core

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/saromanov/pinger/internal/handler"
)

// Core defines struct for main logic
type Core struct {
	hand *handler.Handler
}

// checker gets records from db and checks availability
func (c *Core) checker() {

}

func (core *Core) startCron() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { fmt.Println("Every 1 second") })
	c.Start()
}

// New provides initialization of the core
func New(h *handler.Handler) {
	c := &Core{
		hand: h,
	}
	c.startCron()
}
