// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package execsqlfile

import (
	"database/sql"
	"io/ioutil"
)

func loadFromFile(filePath string) (sqlExpressions []string) {
	fb, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	sqlExpressions = generatorSqlExpress(filterAnnotation(string(fb)))
	return
}

func loadFromString(content string) (sqlExpressions []string) {
	sqlExpressions = generatorSqlExpress(filterAnnotation(content))
	return
}
func ExecSqlExpressionFromFile(filename string,db *sql.DB) (err error) {
	sqlExpressions := loadFromFile(filename)
	for _,expression :=range sqlExpressions{
		_,err = db.Exec(expression)
		if err !=nil{
			break
		}

	}
	return err
}
func ExecSqlExpressionFromString(content string,db *sql.DB) (err error) {
	sqlExpressions :=  loadFromString(content)
	for _,expression :=range sqlExpressions{
		_,err = db.Exec(expression)
		if err !=nil{
			break
		}

	}
	return err
}