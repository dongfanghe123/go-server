package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// LocalTime 自定义时间，序列化输出 yyyy-MM-dd HH:mm:ss 东八区
type LocalTime time.Time

// 序列化
func (t LocalTime) MarshalJSON() ([]byte, error) {
	// 东八区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	val := time.Time(t).In(loc).Format("2006-01-02 15:04:05")
	return []byte(`"` + val + `"`), nil
}

// 反序列化（如果接口接收时间字符串需要）
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	parseTime, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		return err
	}
	*t = LocalTime(parseTime)
	return nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(t)

	// 如果是零值，可以返回 nil（存为 NULL）或者 zeroTime
	if tTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	// 返回标准的 time.Time，MySQL 驱动原生认识这个类型
	return tTime, nil
}

// 2. 顺便实现 sql.Scanner 接口：数据库读取数据到 Go 结构体时调用
func (t *LocalTime) Scan(v interface{}) error {
	if val, ok := v.(time.Time); ok {
		*t = LocalTime(val)
		return nil
	}
	return nil
}

func (t LocalTime) ToTime() time.Time {
	return time.Time(t)
}
