package pkg

import "fmt"

var AggregatorReceiverUseInfoCh = make(chan UseInfo, 100)

var usersInfoMap = make(map[int][]UseInfo)

func Receiver() {

	for {
		receivedUseInfo := <-AggregatorReceiverUseInfoCh

		usersInfoMap[receivedUseInfo.Uid] = append(usersInfoMap[receivedUseInfo.Uid], receivedUseInfo)

	}

}

func printTimestampUser(uid int) {

	fmt.Println("************************************************************")
	fmt.Printf("User: %d\n", uid)
	fmt.Println("Timestamp:")
	for _, v := range usersInfoMap[uid] {
		fmt.Println(v.Timestamp)
	}
	fmt.Println("************************************************************")
}

func AggregatorGetUseInfoUser(uid int) []UseInfo {
	return usersInfoMap[uid]
}
