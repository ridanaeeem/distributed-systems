package main

import (
	"time"
)

// Channel-based aggregator that reports the global average temperature periodically
//
// Report the averagage temperature across all `k` weatherstations every `averagePeriod`
// seconds by sending a `WeatherReport` struct to the `out` channel. The aggregator should
// terminate upon receiving a singnal on the `quit` channel.
//
// Note! To receive credit, channelAggregator must not use mutexes.

func channelAggregator(
	k int,
	averagePeriod float64,
	getWeatherData func(int, int) WeatherReport,
	out chan WeatherReport,
	quit chan struct{},
) {
	// for each batch
	for batch := 0; ; batch++ {
		// find current time and when we should send report and move on to next batch
		cur := time.Now()
		end := cur.Add(time.Millisecond * time.Duration(averagePeriod * 1000))
		// channel for goroutines collecting weather data from k stations to send info back to this func
		ch := make(chan WeatherReport)
		
		// for each of the k weather stations, goroutine to get weatherdata and send it back via channel
		for index := 0; index < k; index++ {
			go func (i int, ch chan WeatherReport, batch int) {
				ch <- getWeatherData(i, batch)
			}(index, ch, batch)
		}

		// running temp
		temp := 0.0
		// count for average
		reportsSeen := 0

		for {
			select {
				// if told to quit, terminate
				case <- quit:
					return;
				default:
			} 
			// if it is not time to report yet
			if time.Now().Before(end){
				select {
					// wait for data in channel and use it to compute average
					case d := <- ch:
						temp += d.Value
						reportsSeen ++
					default:
				} 
			// if it is time to report, send report to out channel
			} else {
				out <- WeatherReport{temp/float64(reportsSeen), -1, batch}
				break
			}
		}
	}
}