package core

import (
	"errors"
	"fmt"

	goping "github.com/sparrc/go-ping"
)

var errSiteNotAvailable = errors.New("site is not available")

// ping provides sending request to the site
func ping(url string) error {
	pinger, err := goping.NewPinger(url)
	if err != nil {
		return fmt.Errorf("ping: unable to ping site %s: %v", url, err)
	}

	pinger.Count = 3
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketsRecv == 0 {
		return errSiteNotAvailable
	}

	return nil
}
