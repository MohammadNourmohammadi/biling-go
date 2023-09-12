package pkg

import (
	"testing"
)

func initCoefficient() {
	coefficient = map[string]map[string]int{
		"vm":  {"cpu": 1, "mem": 2, "net": 3},
		"s3":  {"reqs": 4, "down": 5, "up": 6},
		"vod": {"nodes": 7, "reqs": 8, "bytes": 9},
	}
}

type calculateUseInfoCostEntry struct {
	response      ResponsePrice
	useInfo       UseInfo
	expectedTotal int
}

func TestCalculateUseInfoCost(t *testing.T) {
	initCoefficient()
	var entry = []calculateUseInfoCostEntry{
		{ResponsePrice{map[string]int{}, 0},
			UseInfo{"vm", map[string]int{"cpu": 1, "mem": 2, "net": 3}, 1, 12}, 14},
		{ResponsePrice{map[string]int{}, 0},
			UseInfo{"s3", map[string]int{"reqs": 1, "down": 2, "up": 3}, 1, 12}, 32},
		{ResponsePrice{map[string]int{}, 0},
			UseInfo{"vod", map[string]int{"nodes": 1, "reqs": 2, "bytes": 3}, 1, 12}, 50},
	}
	for _, v := range entry {
		CalculateUseInfoCost(&v.response, v.useInfo)
		if v.response.Total != v.expectedTotal {
			t.Errorf("we expect %d but comput %d", v.expectedTotal, v.response.Total)
		} else {
			t.Logf("test is ok")
		}
	}
}
