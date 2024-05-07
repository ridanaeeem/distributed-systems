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

// func oneBatch(
// 	batch int,
// 	ch chan WeatherReport,
// 	k int,
// 	getWeatherData func(int, int) WeatherReport) {

// 	for index := 0; index < k; index++ {
// 		go func (i int, ch chan WeatherReport, batch int) {
// 			ch <- getWeatherData(i, batch)
// 			fmt.Println("attempted ", i)
// 		}(index, ch, batch)
// 	}
// }

func channelAggregator(
	k int,
	averagePeriod float64,
	getWeatherData func(int, int) WeatherReport,
	out chan WeatherReport,
	quit chan struct{},
) {
	// Your code here.
	for batch := 0; ; batch++ {
		cur := time.Now()
		end := cur.Add(time.Millisecond * time.Duration(averagePeriod * 1000))
		
		ch := make(chan WeatherReport)
		
		for index := 0; index < k; index++ {
			go func (i int, ch chan WeatherReport, batch int) {
				ch <- getWeatherData(i, batch)
			}(index, ch, batch)
		}

		temp := 0.0
		reportsSeen := 0

		for {
			if time.Now().Before(end){
				select {
					case message := <- quit:
						fmt.Println("quit message ", message)
						close(out)
					case d := <- ch:
						temp += d.Value
						reportsSeen ++
					default:
				} 
			} else {
				fmt.Println("exiting")
				out <- WeatherReport{temp/float64(reportsSeen), -1, batch}
				break
			}
		}
	}
}