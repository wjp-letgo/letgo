package lib

//时间相关操作
import (
	"fmt"
	"time"
)

//获得当前时间戳单位s
func Time() int {
	cur := time.Now()
	timestamp := cur.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//TimeByTime 获得时间戳
func TimeByTime(t time.Time) int {
	timestamp := t.UnixNano() / 1000000
	return int(timestamp / 1000)
}
//获得当前时间戳到毫秒
func TimeLong() int64 {
	cur := time.Now()
	return cur.UnixNano() / 1000000
}

//将时间戳转字符串
func TimeToStr(t int) string {
	tm := time.Unix(int64(t), 0)
	return tm.Format("2006-01-02 15:04:05")
}

//将时间戳转字符串
func TimeLongToStr(t int64) string {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//将字符串转时间戳
func StrToTime(date string) int {
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//将字符串转时间戳按格式
func StrToTimeFormat(date string, format string) int {
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation(format, date, loc)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//将字符串转时间戳按格式
func StrToTimeFormatByLocal(date string, format,localName string) int {
	loc, _ := time.LoadLocation(localName)
	tm, _ := time.ParseInLocation(format, date, loc)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//获得当前时间
func Now() string {
	tm := time.Unix(int64(Time()), 0)
	return tm.Format("2006-01-02 15:04:05")
}


//将时间戳转字符串并格式化
func TimeToFormat(t int, format string) string {
	tm := time.Unix(int64(t), 0)
	return tm.Format(format)
}

//将当前时间转字符串
func TimeFormat(format string) string {
	tm := time.Unix(TimeLong(), 0)
	return tm.Format(format)
}

//UTCTime 获得utc时间字符串
func UTCTime() string{
	t,_:=time.ParseInLocation("2006-01-02 15:04:05", Now(), time.Local)
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

//将时间戳转字符串并格式化
func TimeLongToFormat(t int64, format string) string {
	tm := time.Unix(t, 0)
	return tm.Format(format)
}

//获得凌晨零点时间戳
func ZeroTime() int {
	return ZeroTimeByTime(Time())
}

//ZeroTimeByLocal 获得凌晨零点时间戳
func ZeroTimeByLocal(localName string) int {
	return ZeroTimeByTimeByLocal(Time(),localName)
}

//获得时间戳对应的凌晨0点时间戳
func ZeroTimeByTimeByLocal(ti int,localName string) int {
	loc, _ := time.LoadLocation(localName)
	timeStr := TimeToFormat(ti, "2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	timeNumber := t.Unix()
	return int(timeNumber)
}

//获得时间戳对应的凌晨0点时间戳
func ZeroTimeByTime(ti int) int {
	timeStr := TimeToFormat(ti, "2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeNumber := t.Unix()
	return int(timeNumber)
}

//获得今天结束时间戳
func EndTime() int {
	return EndTimeByTime(Time())
}

//获得时间戳对应的23点59分59秒时间戳
func EndTimeByTime(ti int) int {
	timeStr := TimeToFormat(ti, "2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	timeNumber := t.Unix()
	return int(timeNumber)
}

//获得时间戳对应的起始时间和结束时间
func StartAndEndDayTime(ti int) []int {
	return []int{
		ZeroTimeByTime(ti),
		EndTimeByTime(ti),
	}
}

//获得当前时间的年份
func GetYear() int {
	return time.Now().Year()
}

//获得当前时间的年份
func GetYearByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return tm.Year()
}

//获得当前时间的年份
func GetYearByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return tm.Year()
}

//获得当前月份的第一天日期
func GetMonthFirstDay() int {
	year := GetYear()
	month := GetMonth()
	date := fmt.Sprintf("%d-%.2d-%.2d 00:00:00", year, month, 1)
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//获得指定月份的第一天日期
func GetMonthFirstDayByMonth(year, month int) int {
	date := fmt.Sprintf("%d-%.2d-%.2d 00:00:00", year, month, 1)
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//获得当前月份的最后一天日期
func GetMonthLastDay() int {
	year := GetYear()
	month := GetMonth()
	date := fmt.Sprintf("%d-%.2d-%.2d 23:59:59", year, month, 1)
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	tm = tm.AddDate(0, 1, -1)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//获得指定月份的最后一天日期
func GetMonthLastDayByMonth(year, month int) int {
	date := fmt.Sprintf("%d-%.2d-%.2d 23:59:59", year, month, 1)
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	tm = tm.AddDate(0, 1, -1)
	timestamp := tm.UnixNano() / 1000000
	return int(timestamp / 1000)
}

//获得当前时间的月份
func GetMonth() int {
	return int(time.Now().Month())
}

//获得当前时间的月份
func GetMonthByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return int(tm.Month())
}

//获得当前时间的月份
func GetMonthByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return int(tm.Month())
}

//获得当前时间的天数
func GetDay() int {
	return time.Now().Day()
}

//获得当前时间的天数
func GetDayByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return tm.Day()
}
//GetDayByTimeInt
func GetDayByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return tm.Day()
}
//GetTimeByDay 获得当月指定日期的时间戳
func GetTimeByDay(day int)int{
	t:=time.Date(GetYear(),time.Month(GetMonth()) , day, GetHour(), GetMinute(), 0 , 0, time.Local)
	return int(t.Unix())
}

//获得当前时间的小时数
func GetHour() int {
	return time.Now().Hour()
}


//获得当前时间的小时数
func GetHourByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return tm.Hour()
}

//获得当前时间的小时数
func GetHourByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return tm.Hour()
}

//获得当前时间的分钟数
func GetMinute() int {
	return time.Now().Minute()
}


//获得当前时间的分钟数
func GetMinuteByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return tm.Minute()
}

//获得当前时间的分钟数
func GetMinuteByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return tm.Minute()
}

//获得当前时间的星期几 从0-6  其中0表示星期日
func GetWeek() int {
	return int(time.Now().Weekday())
}

//获得当前时间的星期几
func GetWeekByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return int(tm.Weekday())
}

//获得当前时间的星期几
func GetWeekByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return int(tm.Weekday())
}

//获得当前时间的秒数
func GetSecond() int {
	return time.Now().Second()
}


//获得当前时间的秒数
func GetSecondByTimeInt(t int) int {
	var t1 int64 = int64(t)
	tm := time.Unix(t1, 0)
	return tm.Second()
}

//获得当前时间的秒数
func GetSecondByTimeLong(t int64) int {
	var t1 int64 = t / 1000
	tm := time.Unix(t1, 0)
	return tm.Second()
}

//获得周几字符串
func GetWeekStr() string {
	w := int(time.Now().Weekday())
	return GetWeekStrByWeekInt(w)
}
//获得当前时间戳下的周一的时间戳
func GetMondayTime(ti int) int{
	tm := time.Unix(int64(ti), 0)
	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekStartDateInt := int(weekStartDate.Unix())
	return weekStartDateInt
}
//获得指定时间戳的周一到周日的时间戳
func StartAndEndWeekTime(ti int) []int {
	tm := time.Unix(int64(ti), 0)
	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekStartDateInt := int(weekStartDate.Unix())
	return []int{
		weekStartDateInt,
		weekStartDateInt + 24*3600*6 + 23*3600 + 59*60 + 59,
	}
}

//获得周几字符串
func GetWeekStrByWeekInt(w int) string {
	if w == 1 {
		return "周一"
	} else if w == 2 {
		return "周二"
	} else if w == 3 {
		return "周三"
	} else if w == 4 {
		return "周四"
	} else if w == 5 {
		return "周五"
	} else if w == 6 {
		return "周六"
	} else if w == 7 {
		return "周日"
	} else {
		return "周一"
	}
}

//获得指定日期的前一天日期
func GetBeforeTime(year int, month int, day int) time.Time {
	date := fmt.Sprintf("%d-%.2d-%.2d 00:00:00", year, month, day)
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	return tm.AddDate(0, 0, -1)
}

//获得指定日期的前一天的年
func GetBeforeYear(year int, month int, day int) int {
	return GetBeforeTime(year, month, day).Year()
}

//获得指定日期的前一天的月
func GetBeforeMonth(year int, month int, day int) int {
	return int(GetBeforeTime(year, month, day).Month())
}

//获得指定时间戳的那个月的起始时间和结束时间
func StartAndEndMonthTime(ti int) []int {
	tm := time.Unix(int64(ti), 0)
	monthStartDate := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, time.Local)
	monthStartDateInt := int(monthStartDate.Unix())
	lastOfMonth := monthStartDate.AddDate(0, 1, -1)
	return []int{
		monthStartDateInt,
		int(lastOfMonth.Unix()) + 23*3600 + 59*60 + 59,
	}
}

//获得指定日期的前一天的年月
func GetPreviousYearMonthBy(year int, month int) (int, int) {
	return GetBeforeYear(year, month, 1), GetBeforeMonth(year, month, 1)
}

//获得指定日期的上一个年月
func GetPreviousYearMonth() (int, int) {
	year := GetYear()
	month := GetMonth()
	return GetBeforeYear(year, month, 1), GetBeforeMonth(year, month, 1)
}

//获得指定时间戳的那年的起始时间和结束时间
func StartAndEndYearTime(ti int) []int {
	tm := time.Unix(int64(ti), 0)
	yearStartDate := time.Date(tm.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	yearStartDateInt := int(yearStartDate.Unix())
	lastOfYear := yearStartDate.AddDate(1, 0, -1)
	return []int{
		yearStartDateInt,
		int(lastOfYear.Unix()) + 23*3600 + 59*60 + 59,
	}
}

//获得指定日期的前一天的日
func GetBeforeDay(year int, month int, day int) int {
	return GetBeforeTime(year, month, day).Day()
}

//获得指定日期前一天的星期几
func GetBeforeWeek(year int, month int, day int) int {
	return int(GetBeforeTime(year, month, day).Weekday())
}

//获取昨天日
func GetYesterday() int {
	return GetBeforeDay(GetYear(), GetMonth(), GetDay())
}

//获取昨日时间戳
func GetYesterdayInt() int {
	timestamp := GetBeforeTime(GetYear(), GetMonth(), GetDay()).UnixNano() / 1000000
	return int(timestamp / 1000)
}

//获取每个季度的起始月份
func GetStartQuarter() int {
	m := GetMonth()
	if m < 7 {
		return 3
	} else if m < 10 {
		return 6
	} else if m < 12 {
		return 9
	} else {
		return 12
	}
}
