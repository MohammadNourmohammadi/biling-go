package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

var userTokens []UserToken

func ReadTokenFile() {
	jsonFile, err := os.Open(path.Join(os.Getenv("json_path"), "auth-file.json"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &userTokens)

}

func authorization(token string, w http.ResponseWriter) (uid int, err error) {

	hashToken := sha256.Sum256([]byte(token))
	hashString := hex.EncodeToString(hashToken[:])

	uid, message := dbFindUserUidByToken(hashString)
	if message == "not find" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 HTTP status code returned!"))
		return -1, errors.New("not found")
	}
	return uid, nil

}

func findUidByToken(hashToken string) (uid int, err string) {

	for i := 0; i < len(userTokens); i++ {
		if userTokens[i].HashedToken == hashToken {
			return userTokens[i].Uid, "find"
		}
	}

	return -1, "not find"

}
