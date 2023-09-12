package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func ReadEndpoints() []string {
	jsonFile, err := os.Open(path.Join(os.Getenv("json_path"), "endpoints.json"))
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	var endpoints []string
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &endpoints)

	return endpoints
}
