package timer

import "time"

// 函数/方法/结构体/常量：私有小驼峰 导出大驼峰
// 变量：小驼峰
// p.s. 可惜常量不是全大写蛇形

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculatedTime(currentTime time.Time, duration string) (time.Time, error) {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(d), nil
}
