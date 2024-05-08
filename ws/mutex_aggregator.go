package main

import (
	"sync"
	"time"
	// "fmt"
)

// Mutex-based aggregator that reports the global average temperature periodically
//
// Report the averagage temperature across all `k` weatherstations every `averagePeriod`
// seconds by sending a `WeatherReport` struct to the `out` channel. The aggregator should
// terminate upon receiving a singnal on the `quit` channel.
//
// Note! To receive credit, mutexAggregator must implement a mutex based solution.
func mutexAggregator(
	k int,
	averagePeriod float64,
	getWeatherData func(int, int) WeatherReport,
	out chan WeatherReport,
	quit chan struct{},
) {
	// mutex so only one goroutine can change temp at a time
	mutex := sync.Mutex{}
	// for each batch
	for batch := 0; ; batch++ {
		// find current time and when we should send report and move on to next batch
		cur := time.Now()
		end := cur.Add(time.Millisecond * time.Duration(averagePeriod * 1000))
		
		// running temp
		temp := 0.0
		// count for average
		reportsSeen := 0
		
		// for each of the k weather stations, goroutine to get weatherdata and send it back via channel
		for index := 0; index < k; index++ {
			go func (i int, batch int) {
				// once we have data for this batch
				d := getWeatherData(i, batch)
				// we can lock the mutex and make adjustments
				mutex.Lock()
				temp += d.Value
				reportsSeen++
				// after updating running temp and # of reports, unlock mutex 
				mutex.Unlock()
			}(index, batch)
		}


		for {
			select {
				// if told to quit, terminate
				case <- quit:
					return;
				default:
			} 
			// if it is time to report 
			if ! time.Now().Before(end){
				out <- WeatherReport{temp/float64(reportsSeen), -1, batch}
				break
			}
		}
	}
}
