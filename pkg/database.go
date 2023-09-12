package pkg

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path"
)

var DatabaseReceiverUseInfoCh = make(chan UseInfo, 100)
var db *gorm.DB

type Metrics struct {
	gorm.Model
	Service   string
	Tags      map[string]int `gorm:"serializer:json"`
	Uid       int
	Timestamp int
}
type Users struct {
	Uid         int `gorm:"primaryKey"`
	HashedToken string
}

func ConfigDatabase(dbName string) {
	DB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrateErr := DB.AutoMigrate(&Metrics{}, &Users{})
	if migrateErr != nil {
		panic(migrateErr)
	}
	db = DB
}

func ReceiverUseInfo() {
	for {
		receivedUseInfo := <-DatabaseReceiverUseInfoCh
		db.Create(&Metrics{
			Service:   receivedUseInfo.Service,
			Tags:      receivedUseInfo.Tags,
			Uid:       receivedUseInfo.Uid,
			Timestamp: receivedUseInfo.Timestamp,
		})
	}
}

func DatabaseGetUseInfo(uid int) []UseInfo {
	var metrics []Metrics
	db.Where("uid <> ?", uid).Find(&metrics)
	var useinfos []UseInfo

	for _, v := range metrics {
		useinfo := UseInfo{
			Service:   v.Service,
			Tags:      v.Tags,
			Uid:       v.Uid,
			Timestamp: v.Timestamp,
		}
		useinfos = append(useinfos, useinfo)
	}
	return useinfos
}

func ReadUserAndAddToDb() {
	var userTokens []UserToken
	jsonFile, err := os.Open(path.Join(os.Getenv("json_path"), "auth-file.json"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	unMarshalErr := json.Unmarshal(byteValue, &userTokens)
	if unMarshalErr != nil {
		panic(unMarshalErr)
	}
	for _, v := range userTokens {
		db.Create(&Users{
			Uid:         v.Uid,
			HashedToken: v.HashedToken,
		})
	}
}

func dbFindUserUidByToken(hashToken string) (uid int, err string) {

	var user Users
	db.First(&user, "hashed_token = ?", hashToken)
	if user != (Users{}) {
		return user.Uid, "find"
	}
	return -1, "not find"

}
