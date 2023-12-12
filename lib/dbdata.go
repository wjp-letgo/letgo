package lib

import (
	"encoding/xml"
	"reflect"
)

//SqlRows 查询多行
type SqlRows []SqlRow

//表的字段信息
type Columns []Column

//String
func (c Columns)String()string{
	return ObjectToString(c)
}
//ToOutput
func (s SqlRows)ToOutput()[]InRow{
	var list []InRow
	for _,data:=range s{
		d:=make(InRow)
		for k,v:=range data{
			if v.Value!=nil{
				typeOf :=reflect.TypeOf(v.Value)
				switch typeOf.String() {
					case "int64":
						d[k]=v.Int64()
					case "int":
						d[k]=v.Int()
					case "float64":
						d[k]=v.Float64()
					case "float32":
						d[k]=v.Float32()
					case "bool":
						d[k]=v.Value
					default:
						d[k]=v.String()
				}
			}else{
				d[k]=v.Value
			}
		}
		list=append(list, d)
	}
	return list
}

//Bind 绑定对象
func (s SqlRows)Bind(value interface{})bool{
	return StringToObject(ObjectToString(s.ToOutput()), value) 
}

//SqlRow 查询单行
type SqlRow Row

//ToOutput
func (s SqlRow)ToOutput()InRow{
	d:=make(InRow)
	for k,v:=range s{
		if v.Value!=nil{
			typeOf :=reflect.TypeOf(v.Value)
			switch typeOf.String() {
				case "int64":
					d[k]=v.Int64()
				case "int":
					d[k]=v.Int()
				case "float64":
					d[k]=v.Float64()
				case "float32":
					d[k]=v.Float32()
				case "bool":
					d[k]=v.Value
				default:
					d[k]=v.String()
			}
		}else{
			d[k]=v.Value
		}
		
	}
	return d
}

//Bind 绑定对象
func (s SqlRow)Bind(value interface{})bool{
	return StringToObject(ObjectToString(s.ToOutput()), value) 
}

//SqlIn sql插入更新数据格式
type SqlIn InRow
type SqlRaw string //原生sql语句字符串

type Rows []Row //Row 数组数据

//ToOutput
func (s Rows)ToOutput()[]InRow{
	var list []InRow
	for _,data:=range s{
		d:=make(InRow)
		for k,v:=range data{
			if v.Value!=nil{
				typeOf :=reflect.TypeOf(v.Value)
				switch typeOf.String() {
					case "int64":
						d[k]=v.Int64()
					case "int":
						d[k]=v.Int()
					case "float64":
						d[k]=v.Float64()
					case "float32":
						d[k]=v.Float32()
					case "bool":
						d[k]=v.Value
					default:
						d[k]=v.String()
				}
			}else{
				d[k]=v.Value
			}
		}
		list=append(list, d)
	}
	return list
}

//Row 数据
type Row map[string] *Data

//ToOutput
func (s Row)ToOutput()InRow{
	d:=make(InRow)
	for k,v:=range s{
		if v.Value!=nil{
			typeOf :=reflect.TypeOf(v.Value)
			switch typeOf.String() {
				case "int64":
					d[k]=v.Int64()
				case "int":
					d[k]=v.Int()
				case "float64":
					d[k]=v.Float64()
				case "float32":
					d[k]=v.Float32()
				case "bool":
					d[k]=v.Value
				default:
					d[k]=v.String()
			}
		}else{
			d[k]=v.Value
		}
		
	}
	return d
}

//InRow 数据
type InRow map[string]interface{}
//IntRow 整型数据
type IntRow map[int]interface{}
//IntStringMap 
type IntStringMap map[int]string
//StringMap 
type StringMap map[string]string

//MergeInRow 合并InRow
func MergeInRow(values ...InRow)InRow{
	result:=make(InRow)
	for _,row:=range values{
		for k,v:=range row{
			result[k]=v
		}
	}
	return result
}

// MergeInRow 合并SqlIn
func MergeSqlIn(values ...SqlIn) SqlIn {
	result := make(SqlIn)
	for _, row := range values {
		for k, v := range row {
			result[k] = v
		}
	}
	return result
}

//MergeInRow 合并InRow
func MergeIntRow(values ...IntRow)IntRow{
	result:=make(IntRow)
	for _,row:=range values{
		for k,v:=range row{
			result[k]=v
		}
	}
	return result
}
//MergeRow 合并Row
func MergeRow(values ...Row)Row{
	result:=make(Row)
	for _,row:=range values{
		for k,v:=range row{
			result[k]=v
		}
	}
	return result
}
//MarshalXML
func (i InRow)MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t:=xml.ProcInst{
		Target:"xml",
		Inst:[]byte(`version="1.0" encoding="UTF-8"`),
	}
	e.EncodeToken(t)
	start.Name=xml.Name{
		Space: "",
		Local: "map",
	}
	if err:=e.EncodeToken(start);err!=nil{
		return err
	}
	for key,value:=range i{
		elem:=xml.StartElement{
			Name: xml.Name{
				Space: "",
				Local: key,
			},
			Attr: []xml.Attr{},
		}
		if err:=e.EncodeElement(value,elem);err!=nil{
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
//Column 表字段信息
type Column struct{
	Name string `json:"name"`  //字段名
	Type string `json:"type"`  //字段类型
	DataType string `json:"dataType"` //字段类型
	Length int `json:"length"` //字段长度
	Scale int `json:"scale"`  //小数点
	Extra string `json:"extra"` //扩展
	Key string `json:"key"`  //主键为PRI
	IsNull string `json:"isNull"` //是否为空 NO不为空 YES为空
	Default string `json:"default"` //默认值
	Comment string `json:"comment"`  //备注
	CharacterSetName string `json:"characterSetName"` //字段字符编码
	CollationName string `json:"collationName"` //字段字符编码
}
//String
func (c Column)String()string{
	return ObjectToString(c)
}