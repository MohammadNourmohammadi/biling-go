package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getJsonResponseFilter(uid int) ([]byte, error) {
	usage := DatabaseGetUseInfo(uid)
	responseUsage := ResponseUsages{usage}
	return json.Marshal(responseUsage)
}

func handleUsage(w http.ResponseWriter, req *http.Request) {

	uid, err := authorization(req.Header.Get("token"), w)
	if err != nil {
		return
	}
	response, err := getJsonResponseFilter(uid)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))

}

func handleCosts(w http.ResponseWriter, req *http.Request) {
	uid, err := authorization(req.Header.Get("token"), w)
	if err != nil {
		return
	}
	response := CalculateUserUse(uid)
	jsonResponse, err := json.Marshal(*response)
	fmt.Fprintf(w, string(jsonResponse))

	if err != nil {
		fmt.Println(err)
	}
}

func RunService() {
	http.HandleFunc("/usages", handleUsage)
	http.HandleFunc("/costs", handleCosts)
	listenAddress := "0.0.0.0:" + os.Getenv("BILING_PORT")
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		panic(err)
	}
}
