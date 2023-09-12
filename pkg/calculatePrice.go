package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var coefficient map[string]map[string]int

func ReadCoefficient() {

	jsonFile, err := os.Open(path.Join(os.Getenv("json_path"), "coeffiecents.json"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &coefficient)

}

func CalculateUserUse(uid int) *ResponsePrice {
	usage := DatabaseGetUseInfo(uid)
	responseP := new(ResponsePrice)
	responseP.Total = 0
	responseP.PerService = make(map[string]int)

	for i := 0; i < len(usage); i++ {
		CalculateUseInfoCost(responseP, usage[i])
	}
	return responseP
}

func CalculateUseInfoCost(responseP *ResponsePrice, useInfo UseInfo) {
	price := 0

	for k, v := range useInfo.Tags {
		price += coefficient[useInfo.Service][k] * v
	}

	responseP.PerService[useInfo.Service] = responseP.PerService[useInfo.Service] + price
	responseP.Total += price
}
