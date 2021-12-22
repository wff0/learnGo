package testDemo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"testing"
	"time"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

// 插入数据，sql预编译
func (dbw *DbWorker) insertData() {
	stmt, err := dbw.Db.Prepare(`INSERT INTO t_article_cate (cname, addtime, scope) VALUES (?, ?, ?)`)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	ret, err := stmt.Exec("栏目1", time.Now().Unix(), 10)

	// 通过返回的ret可以进一步查询本次插入数据影响的行数
	// RowsAffected和最后插入的Id(如果数据库支持查询最后插入Id)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

// 删除数据，预编译
func (dbw *DbWorker) deleteData() {
	stmt, err := dbw.Db.Prepare(`DELETE FROM t_article_cate WHERE cid=?`)
	ret, err := stmt.Exec(122)
	// 通过返回的ret可以进一步查询本次插入数据影响的行数RowsAffected和
	// 最后插入的Id(如果数据库支持查询最后插入Id).
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

// 修改数据，预编译
func (dbw *DbWorker) editData() {
	stmt, err := dbw.Db.Prepare(`UPDATE t_article_cate SET scope=? WHERE cid=?`)
	ret, err := stmt.Exec(111, 123)
	// 通过返回的ret可以进一步查询本次插入数据影响的行数RowsAffected和
	// 最后插入的Id(如果数据库支持查询最后插入Id).
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

// 查询数据，预编译
func (dbw *DbWorker) queryData() {
	// 如果方法包含Query，那么这个方法是用于查询并返回rows的。其他用Exec()
	// 另外一种写法
	// rows, err := db.Query("select id, name from users where id = ?", 1)
	stmt, _ := dbw.Db.Prepare(`SELECT cid, cname, addtime, scope From t_article_cate where status=?`)
	//err = db.QueryRow("select name from users where id = ?", 1).Scan(&name) // 单行查询，直接处理
	defer stmt.Close()

	rows, err := stmt.Query(0)
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}

	// 构造scanArgs、values两个slice，
	// scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	fmt.Println(columns)
	rowMaps := make([]map[string]string, 9)
	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
		scans[i] = &values[i]
	}
	i := 0
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scans...)

		each := make(map[string]string, 4)
		// 由于是map引用，放在上层for时，rowMaps最终返回值是最后一条。
		for i, col := range values {
			each[columns[i]] = string(col)
		}

		// 切片追加数据，索引位置有意思。不这样写就不是希望的样子。
		rowMaps = append(rowMaps[:i], each)
		fmt.Println(each)
		i++
	}
	fmt.Println(rowMaps)

	for i, col := range rowMaps {
		fmt.Println(i, col)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (dbw *DbWorker) transaction() {
	tx, err := dbw.Db.Begin()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`INSERT INTO t_article_cate (cname, addtime, scope) VALUES (?, ?, ?)`)
	if err != nil {

		fmt.Printf("insert data error: %v\n", err)
		return
	}

	for i := 100; i < 110; i++ {
		cname := strings.Join([]string{"栏目-", strconv.Itoa(i)}, "-")
		//cname := "栏目"+strconv.Itoa(i)
		_, err = stmt.Exec(cname, time.Now().Unix(), i+20)
		if err != nil {
			fmt.Printf("insert data error: %v\n", err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	stmt.Close()
}

func TestMysql(t *testing.T) {
	dbw := DbWorker{Dsn: "root:root@tcp(localhost:3306)/db?charset=utf8mb4"}
	// 支持下面几种DSN写法，具体看MySQL服务端配置，常见为第2种
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	db, err := sql.Open("mysql", dbw.Dsn)
	dbw.Db = db
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
		return
	}
	defer dbw.Db.Close()

	// 插入数据测试
	//dbw.insertData()

	// 删除数据测试
	dbw.deleteData()

	// 修改数据测试
	dbw.editData()

	// 查询数据测试
	dbw.queryData()

	// 事务操作测试
	dbw.transaction()
}
