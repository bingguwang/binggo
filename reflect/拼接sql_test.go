package reflect

import (
    "fmt"
    "reflect"
    "strings"
    "testing"
)

/**
  reflect实现拼接SQL
*/

func Test(t *testing.T) {
    entity := &Stu{Name: "wb"}
    // 获取实体的反射对象
    areValue := reflect.Indirect(reflect.ValueOf(entity))
    areType := reflect.TypeOf(entity).Elem()

    // 通过反射类型获取到结构体的成员
    field, _ := areType.FieldByName("Name")
    fmt.Println("通过反射类型获取到结构体的成员:", field)

    // 通过成员的tag解析出数据库列名
    var db_column string
    if db_column = field.Tag.Get("db"); db_column == "" {
        db_column = field.Name
    }
    fmt.Println("通过成员的tag解析出数据库列名：", db_column)

    // 通过成员名，去反射对象value 获取成员的值
    fieldVal := areValue.FieldByName(field.Name)
    fmt.Println("通过成员名，去反射对象value 获取成员的值:", fieldVal)
    fmt.Println(reflect.Indirect(fieldVal).Interface())

    // 拼接insert SQL
    tableName := "t_user"
    insertSql := fmt.Sprintf("INSERT INTO `%s` (%%s) values (%%s)", tableName)
    fmt.Println(insertSql)

    // 分别是列名， ? , 参数值
    items, values, args := make([]string, 0), make([]string, 0), make([]interface{}, 0)
    items, values, args = append(items, db_column), append(values, "?"), append(args, reflect.Indirect(fieldVal).Interface())

    insertSql = fmt.Sprintf(insertSql, strings.Join(items, ","), strings.Join(values, ","))
    fmt.Println("拼接完后的SQL：", insertSql)
    fmt.Println("args：", args)

    // 然后执行db就行了
    //db := sql.DB{}
    //db.Exec(insertSql, args)
}

type Stu struct {
    Name string `db:"user_name"`
}
