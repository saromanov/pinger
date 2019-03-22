// Package core defines main logic of the pinger
// it contains worker for checking site availability
package core

import (
	"fmt"

	"github.com/robfig/cron"
)

// Core defines struct for main logic
type Core struct {
}

// checker gets records from db and checks availability
func (c *Core) checker() {

}

// New provides initialization of the core
func New() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { fmt.Println("Every 1 second") })
	c.Start()
}
