// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package execsqlfile

import (
	"io/ioutil"
)

func LoadFromFile(filePath string) (sqlExpressions []string) {
	fb, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	sqlExpressions = GeneratorSqlExpress(FilterAnnotation(string(fb)))
	return
}

func LoadFromString(content string) (sqlExpressions []string) {
	sqlExpressions = GeneratorSqlExpress(FilterAnnotation(content))
	return
}