package {{.Model}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
	"log"
	"fmt"
)
{{end}}

{{range .Tables}}

const(
	{{Mapper .Name}}_TableName = ""
)

type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}} {{Type $col}} {{if $col.IsPrimaryKey}}{{else}}`xorm:"{{$col.Name}} {{$col.SQLType.Name}}({{$col.Length}}){{if $col.Nullable}}{{else}} NOTNULL{{end}}{{if $col.IsAutoIncrement}} auroincr{{end}}"`{{end}}
{{end}}
}

func (self *{{Mapper .Name}}) TableName() string {
	return {{Mapper .Name}}_TableName
}

//插入数据
func (self *{{Mapper .Name}}) Create() *{{Mapper .Name}}{
	_, err := engine.InsertOne(self)
	if err != nil {
		log.Fatalln("Create *{{Mapper .Name}} Err:", err)
		return nil
	}

	return self
}

//创建数据表
func (self *{{Mapper .Name}}) CreateTable() {
	err := engine.CreateTables(self)
	if err != nil {
		log.Fatalln("Create {{Mapper .Name}} Table Error:", err)
	}
}

// 删除数据
func (u *{{Mapper .Name}}) Delete() {
	_, err := engine.Delete(u)
	if err != nil {
		log.Fatalln(" Delete {{Mapper .Name}} Error:", err)
	}
}

//更新{{Mapper .Name}}
func (self *{{Mapper .Name}} ) Update() {
	engine.Id(self.Id).Update(self)
}

//获取{{Mapper .Name}}列表
//Find: 获取多条数据
func Get{{Mapper .Name}}List(page, perPage int, wherestr, orderStr string) []*{{Mapper .Name}} {
	var list []*{{Mapper .Name}}
	var err error
	sql := "id is not null"
	if wherestr != "" {
		sql += fmt.Sprintf(" and %s", wherestr)
	}
	if orderStr != "" {
		err = engine.Where(sql).Limit(perPage, (page-1)*perPage).Find(&list)
	} else {
		err = engine.Where(sql).OrderBy(orderStr).Limit(perPage, (page-1)*perPage).Find(&list)
	}
	if err != nil {
		log.Fatalln("Get{{Mapper .Name}}List Error:", err)
		return nil
	}
	return list
}

//{{Mapper .Name}}总数量
func Get{{Mapper .Name}}Count(whereStr string) int64{
	sql := "id is not null"
	if wherestr != "" {
		sql += fmt.Sprintf(" and %s", wherestr)
	}
	total, err := engine.Where(sql).Count(new({{Mapper .Name}}))
	if err != nil {
		log.Fatalln("Get{{Mapper .Name}}Count Error:", err)
		return 0
	}
	return total
}

//通过ID获取{{Mapper .Name}}
func Get{{Mapper .Name}}ById(id int64) *{{Mapper .Name}}{
	item := new({{Mapper .Name}})
	_, err := engine.Where("id = ?", id).Get(item)
	// has, err := x.Id(id).Get(a)
	if err != nil {
		log.Fatalln("Get{{Mapper .Name}}ById Error:", err)
		return nil
	}
	return item
}

//通过原生sql获取
func Get{{Mapper .Name}}BySql(sql string) []*{{Mapper .Name}} {
	var list []*{{Mapper .Name}}
	err := engine.Sql(sql).Find(&list)
	if err != nil {
		log.Fatalln("Get{{Mapper .Name}}BySql Error:", err)
		return nil
	}
	return list
}
{{end}}
