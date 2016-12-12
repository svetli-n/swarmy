// Calculate various request/response metrics.
package stats

import (
	"fmt"
	"time"
)

type Stat struct {
	name  string
	times time.Duration
}

type Stats []Stat

func (s Stats) Collect(stop <-chan bool, times <-chan map[string]time.Duration) {
	for {
		select {
		case <-stop:
			fmt.Println(s)
			return
		case t := <-times:
			for k, v := range t {
				s = append(s, Stat{name: k, times: v})
			}

		}
	}
}
