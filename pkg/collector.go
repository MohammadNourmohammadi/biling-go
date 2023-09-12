package pkg

import (
	"bufio"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

func getUseInfoRequest(endpoint string) []UseInfo {

	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Scan()

	body := scanner.Bytes()

	var useInfos []UseInfo
	json.Unmarshal(body, &useInfos)

	return useInfos
}

func Collect(endpoint string, wg *sync.WaitGroup, timeSleep time.Duration) {
	defer wg.Done()

	for {
		time.Sleep(timeSleep * time.Second)

		useInfos := getUseInfoRequest(endpoint)

		sentToDatabaseUseInfo(useInfos)

	}

}

func sendToAggregatorUseInfo(useInfos []UseInfo) {

	for i := 0; i < len(useInfos); i++ {
		AggregatorReceiverUseInfoCh <- useInfos[i]
	}
}
func sentToDatabaseUseInfo(useInfos []UseInfo) {
	for i := 0; i < len(useInfos); i++ {
		DatabaseReceiverUseInfoCh <- useInfos[i]
	}
}
