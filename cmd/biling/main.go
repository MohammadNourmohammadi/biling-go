package main

import (
	aggregator "biling/pkg"
	auth "biling/pkg"
	calculatePrice "biling/pkg"
	collector "biling/pkg"
	database "biling/pkg"
	endPoints "biling/pkg"
	webservice "biling/pkg"
	"sync"
)

const timeSleep = 5

var wg sync.WaitGroup

func main() {
	database.ConfigDatabase("db_biling.sqlite3")
	database.ReadUserAndAddToDb()

	endpoints := endPoints.ReadEndpoints()

	wg.Add(len(endpoints))

	for i := 0; i < len(endpoints); i++ {
		go collector.Collect(endpoints[i], &wg, timeSleep)
	}

	wg.Add(2)

	go aggregator.Receiver()
	go webservice.RunService()
	go database.ReceiverUseInfo()
	calculatePrice.ReadCoefficient()
	auth.ReadTokenFile()

	wg.Wait()
}
