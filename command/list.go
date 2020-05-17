package command

// List Structure Command List
const (
	Lpush  = "lpush"
	Lrange = "lrange"
	Lpop   = "lpop"
	Lrem   = "lrem"
	Ltrim  = "ltrim"
	// and so on...
)

// ListCommandWhiteList is the white list for simple key value data
// 特に実装はしていません。
var ListCommandWhiteList = map[string]struct{}{
	Lpush:  {},
	Lrange: {},
	Lpop:   {},
	Lrem:   {},
	Ltrim:  {},
}
