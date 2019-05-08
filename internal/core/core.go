// Package core defines main logic of the pinger
// it contains worker for checking site availability
package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron"
	"github.com/saromanov/pinger/internal/handler"
	"github.com/saromanov/pinger/internal/log"
	"github.com/saromanov/pinger/internal/models"
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
	if len(sites) == 0 {
		log.Info("core: site's db is empty")
		return
	}

	batches := len(sites) / batchSize
	iters := batches
	it := 0
	if batches == 0 {
		batches = len(sites)
		iters = 1
	}
	for i := 0; i < iters; i++ {
		var wg sync.WaitGroup
		wg.Add(batches)
		for _, site := range sites[it : it+batches] {
			fmt.Println("SITE: ", site)
			go func(s *pb.Site) {
				start := time.Now()
				err := ping(s.Url)
				end := time.Since(start)
				defer func(delta time.Duration, id int64, e error) {
					err := c.writeStat(end, id, e)
					if err != nil {
						log.Error(err.Error())
					}
					wg.Done()
				}(end, site.Id, err)
			}(site)
			if it >= len(sites) {
				break
			}
		}

		wg.Wait()
		if it >= len(sites) {
			break
		}
		it += batchSize
	}
}

// writeStat provides writing of the stat ingo after ping
func (c *Core) writeStat(duration time.Duration, id int64, e error) error {
	res := &models.PingData{
		SiteID: id,
	}
	if e != nil {
		res.ErrorMessage = e.Error()
	} else {
		res.Available = true
		res.ResponseTime = duration.Nanoseconds()
	}
	_, err := c.hand.CreateStat(res)
	if err != nil {
		return fmt.Errorf("unable to write stat: %v", err)
	}
	log.Infof("Write statistics for site: %d", id)
	return nil
}

// startCron provides starting of the cron worker
func (c *Core) startCron() {
	cr := cron.New()
	cr.AddFunc("@every 1s", c.checker)
	cr.Start()
}

// New provides initialization of the core
func New(h *handler.Handler) {
	c := &Core{
		hand: h,
	}
	c.startCron()
}
