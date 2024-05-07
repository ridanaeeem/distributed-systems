package main

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
	// Your code here.
	// batch := 0
	// cur := time.Now()
	// end := cur.Add(time.Second * time.Duration(averagePeriod))

	// for {
	// 	temp := 0.0
	// 	select {
	// 		case message := <- quit:
	// 			fmt.Println("quit message ", message)
	// 			close(out)
	// 		default:
	// 			if time.Now().Before(end) {
	// 				for index := 0; index < k; index++ {
	// 					go func (int) {
	// 						d := getWeatherData(index, batch)
	// 						temp += d.Value
	// 						fmt.Println(d)
	// 					}(index)
	// 				}
	// 			} else {
	// 				out <- WeatherReport{temp/float64(k), -1, batch}
	// 				batch ++
	// 			}



	// 		// 	for index := 0; index < k; index++ {
	// 		// 		go func (int) {
	// 		// 		}(index)
	// 		// 		if time.Now().Before(end) {
	// 		// 			d := getWeatherData(index, batch)
	// 		// 			temp += d.Value
	// 		// 			fmt.Println(d)
	// 		// 			// fmt.Println("BEFORE", averagePeriod)
	// 		// 		} else {
	// 		// 			// fmt.Println("AFTER", averagePeriod)
	// 		// 			// out <- WeatherReport{temp/float64(k), index, batch}
	// 		// 			batch ++
	// 		// 		}
	// 		// }
	// 	}
	// }
}
