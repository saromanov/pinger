// Package core defines main logic of the pinger
// it contains worker for checking site availability
package core

import (
	"fmt"
	"sync"

	"github.com/robfig/cron"
	"github.com/saromanov/pinger/internal/handler"
	pb "github.com/saromanov/pinger/proto"
)

const batchSize = 20

// Core defines struct for main logic
type Core struct {
	hand *handler.Handler
}

// checker gets records from db and checks availability
func (c *Core) checker() {
	sites, err := c.hand.GetSites(&pb.GetSitesRequest{})
	if err != nil {
		panic(fmt.Sprintf("unable to get list of sites: %v", err))
	}

	batches := len(sites) / batchSize
	it := 0
	for i := 0; i < batches; i++ {
		var wg sync.WaitGroup
		wg.Add(batches)
		for _, site := range sites[it : it+20] {
			go func(s *pb.Site) {
				ping(s.URL)
			}(site)
		}
		it += 20
	}
}

// startCron provides starting of the cron worker
func (c *Core) startCron() {
	cr := cron.New()
	cr.AddFunc("@every 1s", func() { fmt.Println("Every 1 second") })
	cr.Start()
}

// New provides initialization of the core
func New(h *handler.Handler) {
	c := &Core{
		hand: h,
	}
	c.startCron()
}
