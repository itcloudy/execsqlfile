// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package execsqlfile

import (
	"strings"
)

func FilterAnnotation(content string) (sqlLines []string) {
	// split content by \n
	lines := strings.Split(content, "\n")
	// trim space  and filter line annotation
	var trimLines []string
	for _, line := range lines {

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "--") {
			continue
		}
		if len(line) > 0 {
			trimLines = append(trimLines, line)
		}
	}
	var blockStart bool
	var blockEnd bool
	// filter block  annotation
	for _, line := range trimLines {
		// annotation start
		if strings.HasPrefix(line, "/*") && blockEnd == false {
			blockStart = true
			continue
		}
		// annotation end
		if strings.HasSuffix(line, "*/") && blockStart == true {
			blockEnd = true
			continue
		}
		// annotation body
		if blockStart == true && blockEnd == false {
			continue
		}
		// block annotation end
		if blockStart == true && blockEnd == true {
			blockStart = false
			blockEnd = false
		}
		// sql line
		if blockStart == false && blockEnd == false {
			sqlLines = append(sqlLines, line)
		}

	}

	return
}
func GeneratorSqlExpress(sqlLines []string) (sqlExpressions []string) {
	var sqlEnd bool
	var sqlExpression string
	for _, line := range sqlLines {
		if strings.HasSuffix(line, ";") {
			sqlEnd = true
		}
		sqlExpression += line
		sqlExpression += "  "
		if sqlEnd {
			sqlExpressions = append(sqlExpressions, sqlExpression)
			sqlEnd = false
			sqlExpression = ""
		}
	}
	return
}
