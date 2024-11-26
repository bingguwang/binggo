package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
*
在go代码里实现与mysql递归查询同样效果的查询
*/
func findAllChildren(parentID int64) ([]int64, error) {
	var result []int64
	var queue []int64
	queue = append(queue, parentID) // 将起始ID加入队列

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]

		// 查询当前ID的所有直接子节点
		children := getChildOrgByParentId(currentID)

		// 将子节点添加到结果集
		result = append(result, children...)

		// 将所有子节点的ID添加到队列中，以便进一步查找它们的子节点
		for _, child := range children {
			queue = append(queue, child)
		}
	}

	return result, nil
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql",
		"root:123456@tcp(192.168.2.133:3306)/platform")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	defer db.Close()

	children, err := findAllChildren(2596)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(children)
}

func getChildOrgByParentId(parent int64) []int64 {
	var res []int64
	rows, err := db.Query("select id from organization_organization where parent_id=?", parent)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil
		}
		res = append(res, id)
	}

	return res
}
