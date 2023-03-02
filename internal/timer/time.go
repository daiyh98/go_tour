package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d) //从字符串中解析出持续时间duration，有效单位ns、us/μs、ms、s、m、h
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil //时间位移
}
