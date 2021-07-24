package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

type Data struct {
	Id       int
	UserName string
}

func Select(db *sql.DB) (Data, error) {
	var sqlStr = `select id,user_name from tab where id=?`
	var arg = 101
	var errCtx = fmt.Sprintf("sqlStr=%v|arg=%v", sqlStr, arg)

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return Data{}, errors.Wrap(err, errCtx)
	}
	row := stmt.QueryRow(arg)
	var data Data
	err = row.Scan(&data.Id, &data.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return Data{}, errors.Wrap(sql.ErrNoRows, errCtx)
		} else {
			return Data{}, errors.Wrap(err, errCtx)
		}
	}

	return data, nil
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db")
	if err != nil {
		logrus.WithError(err).Error("sql.Open db failed")
		return
	}
	_, err = Select(db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logrus.WithError(err).Error("Select failed, no record")
			return
		} else {
			logrus.WithError(err).Error("Select failed, other error")
			return
		}
	}

}
