package lang

import "github.com/starter-go/base/lang"

////////////////////////////////////////////////////////////////////////////////
// basic

// TimeStamp 用 int64 表示一个 unix 时间戳，单位是毫秒，基于 UTC_1970-01-01_00:00:00
type TimeStamp = lang.Time

// TimeSpan 以毫秒为单位表示时间长度
type TimeSpan = lang.Milliseconds

////////////////////////////////////////////////////////////////////////////////
// alias

type Time = TimeStamp

type Milliseconds = TimeSpan

// Seconds 以秒为单位表示时间长度
type Seconds = lang.Seconds

////////////////////////////////////////////////////////////////////////////////
// functions

func Now() Time {
	return lang.Now()
}
