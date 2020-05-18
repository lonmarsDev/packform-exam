package utils

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func TrimObjectChar(i string) string {
	i = strings.ReplaceAll(i, `ObjectID(`, ``)
	i = strings.ReplaceAll(i, `/`, ``)
	i = strings.ReplaceAll(i, `\`, ``)
	i = strings.ReplaceAll(i, `\\`, ``)
	i = strings.ReplaceAll(i, "\\", ``)
	i = strings.ReplaceAll(i, `)`, ``)
	i = strings.ReplaceAll(i, `"`, ``)
	return i
}

func Int64ToStr(s int64) string {
	return strconv.FormatInt(s, 10)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func InMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func PrimitiveDateTime(t *time.Time) *primitive.DateTime {
	ct := time.Now()

	priTime := new(primitive.DateTime)
	if t != nil {
		*priTime = primitive.DateTime(InMillis(*t))
		return priTime
	}
	*priTime = primitive.DateTime(InMillis(ct))
	return priTime

}

//Convert string time to primative datetime format
func StringTimeToPrimitive(s string) (*primitive.DateTime, error) {
	datetime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, errors.New("invalid time format")
	}
	return PrimitiveDateTime(&datetime), nil

}

//Convert string datetime to primative datetime format
func StringToDateime(s string) (*time.Time, error) {
	datetime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, errors.New("invalid time format")
	}
	return &datetime, nil

}

func StringToFloat32(s string) float64 {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return float64(0)
	}
	return v
}

func TimeUserFormat(s time.Time) string {
	panic("Todo")
}

func IntToPointer(i int) *int {
	return &i
}
