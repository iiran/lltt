package helper

// 1秒(s)=1000000000纳秒(ns)
const ONE_SEC_NANO = 1000000000

// 1天(d)=86400秒(s)
const ONE_DAY_SEC = 86400

// 1时(h)=3600秒(s)
const ONE_HOUR_SEC = 3600

const ONE_HOUR_NANO = ONE_HOUR_SEC * ONE_SEC_NANO

func NanoToSec(nano int64) int64 {
	return nano / ONE_SEC_NANO
}

func SecToNano(sec int64) int64 {
	return sec * ONE_SEC_NANO
}

func HourToNano(hour int64) int64 {
	return hour * ONE_HOUR_NANO
}
