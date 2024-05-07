package main

import (
	"fmt"
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
	// Your code here.
	batch := 0
	cur := time.Now()
	end := cur.Add(time.Second * time.Duration(averagePeriod))
	ch := make(chan WeatherReport)
	sentProcess := false
	temp := 0.0
	reportsSeen := 0
	for {
		select {
			case message := <- quit:
				fmt.Println("quit message ", message)
				close(out)
			case d := <- ch:
				temp += d.Value
				reportsSeen ++
			default:
				if time.Now().Before(end){
					if (!sentProcess) {
						for index := 0; index < k; index++ {
							go func (i int, ch chan WeatherReport, batch int) {
								ch <- getWeatherData(i, batch)
							}(index, ch, batch)
						}
					}
					sentProcess = true
				} else {
					out <- WeatherReport{temp/float64(reportsSeen), -1, batch}
					batch ++
					sentProcess = false
					temp = 0.0
					reportsSeen = 0
					cur = time.Now()
					end = cur.Add(time.Second * time.Duration(averagePeriod))
				}
		}
	}
}