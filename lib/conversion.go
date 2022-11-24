package lib

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"strings"
)

//字符串转float32
func StrToFloat32(str string) float32 {
	vv, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return float32(vv)
	}
	return 0
}

//Round
func Round(x float32) int {
	return int(math.Floor(float64(x) + 0.5))
}

//字符串转float64
func StrToFloat64(str string) float64 {
	vv, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return vv
	}
	return 0
}

//字符串转int
func StrToInt(str string) int {
	vv, err := strconv.Atoi(str)
	if err == nil {
		return vv
	}
	return 0
}

//字符串转uint
func StrToUInt(str string) uint {
	vv, err := strconv.Atoi(str)
	if err == nil {
		return uint(vv)
	}
	return 0
}

//字符串转int64
func StrToInt64(str string) int64 {
	vv, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return vv
	}
	return 0
}

//将float64转int
func Float64ToInt(f float64) int {
	return int(f)
}

//将Interface转int
func InterfaceToInt(data interface{}) int {
	str := fmt.Sprintf("%v", data)
	return StrToInt(str)
}

//将Interface转String
func InterfaceToString(data interface{}) string {
	str := fmt.Sprintf("%v", data)
	return str
}

//将Interface转int
func InterfaceToInt64(data interface{}) int64 {
	str := fmt.Sprintf("%v", data)
	return StrToInt64(str)
}

//float64转int64
func Float64ToInt64(data float64) int64 {
	return int64(data)
}

//InRowToSqlRow 将InRow 转SqlRow
func InRowToSqlRow(row InRow) SqlRow {
	record := make(SqlRow)
	for k, v := range row {
		row := &Data{}
		row.Set(v)
		record[k] = row
	}
	return record
}

//RowsToSqlRows sql.Rows转 SqlRows
func RowsToSqlRows(rows *sql.Rows) SqlRows {
	cols, err := rows.Columns()
	if err != nil {
		return nil
	}
	scanArgs := make([]interface{}, len(cols))
	cs, _ := rows.ColumnTypes()
	for i, v := range cs {
		switch v.DatabaseTypeName() {
		case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
			scanArgs[i] = new(sql.NullString)
			break
		case "INT", "BIGINT", "BIT", "TINYINT", "INTEGER", "MEDIUMINT", "NUMERIC", "SMALLINT":
			scanArgs[i] = new(sql.NullInt64)
			break
		case "DECIMAL", "DOUBLE", "FLOAT":
			scanArgs[i] = new(sql.NullFloat64)
			break
		case "BOOL":
			scanArgs[i] = new(sql.NullBool)
			break
		default:
			scanArgs[i] = new(sql.NullString)
		}
	}
	var list SqlRows
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		record := make(SqlRow)
		for i, col := range scanArgs {
			row := &Data{}
			switch cs[i].DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				row.Set(col.(*sql.NullString).String)
				break
			case "INT", "BIGINT", "BIT", "TINYINT", "INTEGER", "MEDIUMINT", "NUMERIC", "SMALLINT":
				row.Set(col.(*sql.NullInt64).Int64)
				break
			case "DECIMAL", "DOUBLE", "FLOAT":
				row.Set(col.(*sql.NullFloat64).Float64)
				break
			case "BOOL":
				row.Set(col.(*sql.NullBool).Bool)
				break
			default:
				row.Set(col.(*sql.NullString).String)
				break
			}
			record[cols[i]] = row
		}
		list = append(list, record)
	}
	return list
}

//XmlObjectToString 将xml对象转字符串
func XmlObjectToString(data interface{}) string {
	js, err := xml.Marshal(data)
	if err != nil {
		return ""
	}
	return string(js)
}

//JsonObjectToString 将对象转成json字符串
func JsonObjectToString(data interface{}) string {
	return ObjectToString(data)
}

//ObjectToString 将对象转成json字符串
func ObjectToString(data interface{}) string {
	js, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(js)
}

//StringToJsonObject json字符串转json对象
func StringToJsonObject(str string, data interface{}) bool {
	return StringToObject(str, data)
}

//StringToObject json字符串转对象
func StringToObject(str string, data interface{}) bool {
	js := json.NewDecoder(bytes.NewReader([]byte(str)))
	js.UseNumber()
	err := js.Decode(data)
	if err == nil {
		return true
	}
	return false
}

//StringToXmlObject xml字符串转对象
func StringToXmlObject(str string, data interface{}) bool {
	err := xml.Unmarshal([]byte(str), data)
	if err == nil {
		return true
	}
	return false
}

//JSONToMap
func JSONToMap(str string) InRow {
	var tempMap InRow
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		return nil
	}
	return tempMap
}

//MapToJSON
func MapToJSON(list Rows) string {
	s := list.ToOutput()
	return ObjectToString(s)
}

//Int64ArrayToInterfaceArray int64转[]interface{}
func Int64ArrayToInterfaceArray(data []int64) []interface{} {
	var it []interface{}
	for _, v := range data {
		it = append(it, v)
	}
	return it
}

//StringArrayToInterfaceArray string转[]interface{}
func StringArrayToInterfaceArray(data []string) []interface{} {
	var it []interface{}
	for _, v := range data {
		it = append(it, v)
	}
	return it
}

//StringArrayToInt64Array string转[]int64
func StringArrayToInt64Array(data []string) []int64 {
	var it []int64
	for _, v := range data {
		it = append(it, StrToInt64(v))
	}
	return it
}

//interface转ArrayString
func InterfaceArrayToArrayString(list []interface{}) []string {
	var rp []string
	for _, v := range list {
		data := &Data{Value: v}
		rp = append(rp, data.String())
	}
	return rp
}

//interface转ArrayString
func InterfaceArrayToArrayInt64(list []interface{}) []int64 {
	var rp []int64
	for _, v := range list {
		data := &Data{Value: v}
		rp = append(rp, data.Int64())
	}
	return rp
}

//interface转ArrayRow
func InterfaceArrayToArrayRow(list []interface{}) []Row {
	var rp []Row
	for _, v := range list {
		data := &Data{Value: v}
		rp = append(rp, data.Row())
	}
	return rp
}

//interface转Rows
func InterfaceArrayToArrayRows(list []interface{}) Rows {
	var rp Rows
	for _, v := range list {
		data := &Data{Value: v}
		rp = append(rp, data.Row())
	}
	return rp
}

//map[string]interface {}转Row
func MapInterfaceArrayToRow(list map[string]interface{}) Row {
	var rp Row = make(Row)
	for k, v := range list {
		rp[k] = &Data{Value: v}
	}
	return rp
}

//interface转ArrayString
func Int64ArrayToArrayString(list []int64) []string {
	var rp []string
	for _, v := range list {
		rp = append(rp, fmt.Sprintf("%d", v))
	}
	return rp
}

//Utf8ToGb2312 UTF8转GBK2312
func Utf8ToGb2312(src string) string {
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GB18030.NewEncoder()))
	return string(data)
}

//Gb2312ToUtf8 utf8转gbk
func Gb2312ToUtf8(src string) string {
	data, _ := ioutil.ReadAll(simplifiedchinese.GB18030.NewDecoder().Reader(bytes.NewReader([]byte(src))))
	rts := string(data)
	return rts
}

//JsonArrayStringToStringArray json数组字符串转字符串数组
func JsonArrayStringToStringArray(in string) []string {
	var js []string
	err := json.Unmarshal([]byte(in), &js)
	if err == nil {
		return js
	}
	return nil
}

//CopyJSON 拷贝对象
func CopyJSON(input interface{}, out interface{}) {
	aj, _ := json.Marshal(input)
	_ = json.Unmarshal(aj, out)
}

//\u00转义
func U00(str string) string {
	str = strings.Replace(str, "\\u0026", "&", -1)
	str = strings.Replace(str, "\\u003c", "<", -1)
	str = strings.Replace(str, "\\u003e", ">", -1)
	return str
}

//U00Byte\u00转义
func U00Byte(str []byte) []byte {
	str = bytes.Replace(str, []byte("\\u0026"), []byte("&"), -1)
	str = bytes.Replace(str, []byte("\\u003c"), []byte("<"), -1)
	str = bytes.Replace(str, []byte("\\u003e"), []byte(">"), -1)
	return str
}
