// Calculate various request/response metrics.
package stats

import (
	"fmt"
	"sort"
	"time"
)

type Stats map[string][]time.Duration

func (s Stats) Collect(stop <-chan bool, times <-chan map[string]time.Duration, merge chan map[string][]time.Duration) {
	for {
		select {
		case <-stop:
			merge <- s
			return
		case t := <-times:
			for k, v := range t {
				s[k] = append(s[k], v)
			}

		}
	}
}

type Merger struct {
	data map[string][]time.Duration
}

func (m Merger) Run(merge chan map[string][]time.Duration) {
	for {
		select {
		case m.data = <-merge:
			for _, val := range m.data {
				slice := timeSlice(val)
				sort.Sort(slice)
				n := len(slice)
				fmt.Println("Requests: ", n)
				fmt.Println("Min: ", slice[0])
				fmt.Println("Max: ", slice[n-1])
				var sum time.Duration
				for i := 0; i < n; i++ {
					sum += slice[i]
				}
				fmt.Println("Mean: ", (time.Duration(int64(sum/time.Nanosecond)/int64(n)) * time.Nanosecond))
				if n%2 == 0 {
					fmt.Println("Median: ", time.Duration(int64(slice[n/2-1]+slice[n/2+1])/int64(2))*time.Nanosecond)
				} else {
					fmt.Println("Median: ", slice[n/2])
				}
			}
			return
		default:
		}
	}
}

type timeSlice []time.Duration

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
