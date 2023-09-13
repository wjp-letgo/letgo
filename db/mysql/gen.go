package mysql

import (
	"fmt"
	"strings"

	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
)

//GenModelAndEntity 自动生成Model和Entity
func GenModelAndEntity(moduleName string,dbConfig SlaveDB,isCode bool){
	g:=genInfo{
		dbConfig:dbConfig,
		isCode: isCode,
		moduleName:moduleName,
	}
	g.run()
}
//GenModelAndEntity 自动生成Model和Entity
func GenModelAndEntityByTableName(moduleName,table string,dbConfig SlaveDB,isCode bool){
	g:=genInfo{
		dbConfig:dbConfig,
		isCode: isCode,
		moduleName:moduleName,
	}
	g.runByTable(table)
}
type genInfo struct{
	dbConfig SlaveDB
	moduleName string
	isCode bool
	db DB
}
func (g *genInfo)run(){
	g.db=g.newDb()
	tables:=g.db.ShowTables()
	for _,tb:=range tables{
		g.genModel(tb)
		g.genEntity(tb)
	}
}
//runByTable 生成单表的go源码文件
func (g *genInfo)runByTable(table string){
	g.db=g.newDb()
	g.genModel(table)
	g.genEntity(table)
}

//newDb 初始化数据库
func (g *genInfo) newDb() DB{
	var configs []MysqlConnect
	configs=append(configs, MysqlConnect{
		Master:g.dbConfig,
		Slave:nil,
	})
	//fmt.Println(configs)
	var db DB
	p:=&MysqlPool{}
	p.Init(configs)
	db.SetPool(p)
	db.SetDB(g.dbConfig.Name,g.dbConfig.DatabaseName)
	return db
}
//FieldName
func (g *genInfo)FieldName(feild string)string{
	arr:=strings.Split(feild,"_")
	m:=""
	for _,k:=range arr{
		m+=lib.FirstToUpper(k)
	}
	return m
}
func (g *genInfo)getTableName(table string)string{
	table=g.getSmallTableName(table)
	arr:=strings.Split(table,"_")
	m:=""
	for i,k:=range arr{
		if g.dbConfig.Prefix!=""&&i==0&&k==g.dbConfig.Prefix{
		}else if g.dbConfig.Suffix!=""&&i==len(arr)-1&&k==g.dbConfig.Suffix{
		}else{
			m+=lib.FirstToUpper(arr[i])
		}
	}
	return m
}
//getDbName
func (g *genInfo)getDbName(dbName string)string{
	arr:=strings.Split(dbName,"_")
	var arr2 []string
	for i,k:=range arr{
		if g.isCode{
			if i!=len(arr)-1{
				arr2=append(arr2, k)
			}
		}else{
			arr2=append(arr2, k)
		}
	}
	return strings.Join(arr2,"_")
}
//getSmallTableName
func (g *genInfo)getSmallTableName(table string)string{
	arr:=strings.Split(table,"_")
	var arr2 []string
	for i,k:=range arr{
		if g.isCode{
			if i!=len(arr)-1{
				arr2=append(arr2, k)
			}
		}else{
			arr2=append(arr2, k)
		}
	}
	return strings.Join(arr2,"_")
}
//getModelName
func (g *genInfo) getModelName(table string)string{
	return g.getTableName(table)+"Model"
}
//getEntityName
func (g *genInfo) getEntityName(table string)string{
	return g.getTableName(table)+"Entity"
}
//genModel
func (g *genInfo) genModel(table string){
	modelName:=g.getModelName(table)
	entityName:=g.getEntityName(table)
	//tableName:=g.getTableName(table)
	content:=fmt.Sprintf(
		g.getModelTemplate(),
		g.moduleName,
		modelName,
		modelName,
		modelName,
		modelName,
		modelName,
		modelName,
		g.getDbName(g.dbConfig.DatabaseName),
		g.getSmallTableName(table),
		modelName,
		entityName,
		modelName,
		modelName,
		entityName,
		entityName,
	)
	fileName:=fmt.Sprintf("model/%s.go",strings.ToLower(modelName))
	file.PutContent(fileName,content)
	//fmt.Println(content)
}

var mysqlTypMap map[string]string=map[string]string{
	"bigint":"int64",
	"binary":"[]byte",
	"bit":"byte",
	"blob":"[]byte",
	"char":"string",
	"date":"string",
	"datetime":"string",
	"decimal":"float32",
	"double":"float64",
	"float":"float32",
	"int":"int",
	"integer":"int",
	"varchar":"string",
	"tinyint":"int",
	"time":"int",
	"tinytext":"string",
	"text":"string",
	"longtext":"string",
	"mediumtext":"string",
	"mediumint":"int",
	"smallint":"int",
}

//genEntity
func (g *genInfo) genEntity(table string){
	entityName:=g.getEntityName(table)
	c:=g.db.Desc(table)
	cn:=""
	for _,v:=range c{
		cn+=fmt.Sprintf("    %s    %s  `json:\"%s\"`\n",g.FieldName(v.Name),mysqlTypMap[v.DataType], v.Name)
	}
	tpl:=g.getEntityTemplate()
	content:=fmt.Sprintf(
		tpl,
		entityName,
		entityName,
		cn,
		entityName,
	)
	//fmt.Println(content)
	fileName:=fmt.Sprintf("model/entity/%s.go",strings.ToLower(entityName))
	file.PutContent(fileName,content)
}
//getModelTemplate 模型文件模板文件
func (g *genInfo) getModelTemplate()string{
	tmp:=`package model

import (
    "github.com/wjp-letgo/letgo/db/mysql"
    "github.com/wjp-letgo/letgo/lib"
	"%s/model/entity"
)

//%s
type %s struct{
    mysql.Model
}
`
	if g.isCode{
		tmp+=`//Get%s 获得操作模型
func Get%s(dbCode,tableCode string) *%s{
    model:=&%s{}
    model.Init("%s_"+dbCode,"%s_"+tableCode)
    //开启软删除
    model.SoftDelete=true
    return model
}`
	}else{
		tmp+=`//Get%s 获得操作模型
func Get%s() *%s{
    model:=&%s{}
    model.Init("%s","%s")
    //开启软删除
    model.SoftDelete=true
    return model
}`
	}
	tmp+=`
//SaveByEntity
func (m *%s)SaveByEntity(data entity.%s) int64{
    var inData lib.SqlIn
    lib.StringToObject(data.String(), &inData)
    inData["delete_time"]=-1
    if data.Id>0{
        inData["update_time"]=lib.Time()
        delete(inData,"create_time")
        m.Where("id", data.Id).Update(inData)
        return data.Id
    }else{
        inData["create_time"]=lib.Time()
        delete(inData,"id")
        delete(inData,"update_time")
        return m.Insert(inData)
    }
}
//SaveByInRow 保存
func (m *%s)SaveByInRow(id int64,data lib.SqlIn) int64{
    if id>0{
        m.Where("id",id).Update(data)
        return id
    }else{
        return m.Insert(data)
    }
}
//GetById 通过id获得数据
func (m *%s) GetById(id int64) lib.SqlRow{
    return m.Where("id", id).Find()
}
//GetEntityById 通过id获得数据
func (m *%s) GetEntityById(id int64) entity.%s{
    var out entity.%s
    data:= m.Where("id", id).Find()
    data.Bind(&out)
    return out
}`
	return tmp
}

//getEntityTemplate 获得实体模板文件
func (g *genInfo)getEntityTemplate()string{
	tmp:=`package entity

import (
    "github.com/wjp-letgo/letgo/lib"
)

//%s
type %s struct{
%s
}
//String
func (e *%s)String()string{
    return lib.ObjectToString(e)
}
`
	return tmp
}