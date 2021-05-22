package helpers

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mintance/go-uniqid"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateMd5(key string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}

func Ksort(c map[string]interface{}) (map[string]interface{}, []string) {
	to := make(map[string]interface{})
	var keys []string
	for s := range c {
		if s != "hash" {
			keys = append(keys, s)
		}

	}
	sort.Strings(keys)

	for _, v := range keys {
		str := fmt.Sprint(c[v])
		to[v] = str
	}

	return to, keys
}

func GetCurrentTime() time.Time {
	timezone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(timezone)

	return now
}

func VoucherActionConvert(i string) string {
	if i == "u" {
		return "sử dụng"
	} else if i == "b" {
		return "mua"
	}
	return i
}

func GetDayOfYear() int64 {
	dayOfYear := GetCurrentTime().YearDay()
	return int64(dayOfYear)
}

func CurrentHour() time.Time {
	timezone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	year, week := GetCurrentTime().ISOWeek()

	date := time.Date(year, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, timezone)

	isoYear, isoWeek := date.ISOWeek()

	for date.Day() < time.Now().Day() { // iterate back to Monday
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}

	for date.Hour() != GetCurrentTime().Hour() { // iterate back to Monday
		date = date.Add(time.Hour)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

func PrevHour() time.Time {
	timezone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	year, week := GetCurrentTime().ISOWeek()

	date := time.Date(year, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	for date.Hour() != GetCurrentTime().Hour()-1 { // iterate back to Monday
		date = date.Add(time.Hour)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

func NextHour() time.Time {
	timezone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	year, week := GetCurrentTime().ISOWeek()

	date := time.Date(year, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	next_hour := GetCurrentTime().Hour() + 1
	if GetCurrentTime().Hour() == 23 {
		next_hour = 0
	}
	for date.Hour() != next_hour { // iterate back to Monday
		date = date.Add(time.Hour)
		fmt.Println(2)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

func FirstDayOfISOWeek() time.Time {
	timezone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	year, week := GetCurrentTime().ISOWeek()

	date := time.Date(year, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

func EndDayOfISOWeek() time.Time {
	timezone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	year, week := GetCurrentTime().ISOWeek()

	date := time.Date(year, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	for date.Weekday() != time.Sunday { // iterate back to Monday
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

func BusinessID(prefix string) string {
	id := prefix + RandomInt(999, 1) + uniqid.New(uniqid.Params{prefix, false}) + RandomInt(99, 1)

	return strings.ToUpper(id)
}

func TransID(prefix string) string {
	id := prefix + RandomInt(999, 1) + uniqid.New(uniqid.Params{prefix, false}) + RandomInt(99, 1)

	return strings.ToUpper(id)
}

func RandomInt(max, min int64) string {
	return fmt.Sprint(rand.Int63n(max-min) + min)
}

func JsonToArrayValues(jsonstring string) (data map[string]interface{}) {
	c := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonstring), &c)

	if err != nil {
		return data
	}

	for k, _ := range c {
		typeof := reflect.TypeOf(c[k])
		if fmt.Sprint(typeof) == "float64" || fmt.Sprint(typeof) == "float" {
			c[k] = fmt.Sprintf("%.f", c[k])
		}
	}

	return c
}
func ContextWithTimeOut() (context_data context.Context) {
	context_data, _ = context.WithTimeout(context.Background(), time.Minute*10)
	return context_data
}
