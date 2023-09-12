package pkg

type UseInfo struct {
	Service   string         `json:"service"`
	Tags      map[string]int `json:"tags"`
	Uid       int            `json:"uid"`
	Timestamp int            `json:"timestamp"`
}

type ResponseUsages struct {
	Usages []UseInfo `json:"usages"`
}

type ResponsePrice struct {
	PerService map[string]int `json:"per_service"`
	Total      int            `json:"total"`
}

type UserToken struct {
	Uid         int    `json:"uid"`
	HashedToken string `json:"hashed_token"`
}
