package main

import (
	"fmt"
	"strings"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


//数据库配置
const (
    userName = "root"
    password = "root"
    ip = "172.31.7.23"
    port = "3306"
    dbName = "shorturl"
)
//Db数据库连接池
var DB *sql.DB

func InitDB()  {
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
    DB, _ = sql.Open("mysql", path)
    //设置数据库最大连接数
    DB.SetConnMaxLifetime(100)
    //设置上数据库最大闲置连接数
    DB.SetMaxIdleConns(10)
	
	err := DB.Ping();
	if (err != nil) {
		fmt.Println("opon database fail")
		// time.Sleep(1000 * time.Millisecond)
		// DB, _ = sql.Open("mysql", path)
		// err = DB.Ping();
	}
    fmt.Println("connnect success")
}

func CheckUrl(url string) bool {
	var shortlink string
	_ = DB.QueryRow(`SELECT shortlink From urls where shortlink = ?`, url).Scan(&shortlink)
	// if err != nil {
	// 	panic("query error")
	// }
	fmt.Println(shortlink)
	if shortlink == "" {
		fmt.Println("false")
		return false
	}
	fmt.Println("true")
	return true
}

func InsertUrl(url string, path string) bool {
	if (CheckUrl(url)) {
		return false
	} else {
		var shortlink string
    	var expiration_length_in_minutes int
    	var created_at time.Time
		var paste_path string
		
		shortlink = url
		expiration_length_in_minutes = 60
		created_at = time.Now()
		paste_path = path

		ret, err := DB.Exec(`INSERT INTO urls (shortlink, expiration_length_in_minutes, created_at, paste_path) VALUES (?, ?, ?, ?)`, shortlink, expiration_length_in_minutes, created_at, paste_path)
		if err != nil {
			fmt.Println("insert error")
			return false
		}
		fmt.Println(ret)
		return true
	}
}

func GetPath(shortlink string) string {
	var paste_path string
	_ = DB.QueryRow(`SELECT paste_path From urls where shortlink = ?`, shortlink).Scan(&paste_path)
	fmt.Println("paste_path ", paste_path)
	return paste_path
}