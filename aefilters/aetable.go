package main

import (
	"io/ioutil"
	"log"

	"github.com/agileera/aeoffice/aetable"

	pf "github.com/oltolm/go-pandocfilters"
)

// D:\work\goprojects\src\github.com\agileera\aeoffice\aefilters\test>pandoc
// aeinclude_sample.md  --to json | D:\work\goprojects\bin\aefilters.exe

func aetablefilter(key string, value interface{}, format string, meta interface{}) interface{} {
	if key == "CodeBlock" {
		// fmt.Println("------------")
		// fmt.Println(value)
		// fmt.Println(format)
		// fmt.Println(meta)
		// fmt.Println("------------")

		c := value.([]interface{})
		// attrs := c[0].([]interface{})
		code := c[1].(string)
		config, err := aetable.ParseConfig(code)
		if err != nil {
			log.Println(err)
			return nil
		}

		table, _ := aetable.ParseExcel(config)
		// log.Println(table)

		pantable := pf.Table(nil, table.ColAlign, table.RelColWidth, table.ColHeader, table.Rows)

		// 	caption []interface{}, colAlign []interface{}, relColWidth []interface{},
		// colHeader []interface{}, rows [][]interface{}
		// ident := attrs[0]
		// classes := attrs[1].([]interface{})
		// fmt.Println(ident)
		// fmt.Println(code)
		// fmt.Println(classes)

		// para := []interface{}{}
		// para = append(para, pantable)
		// if contains(classes, "aetable") {
		// 	for _, path := range strings.Split(code.(string), "\n") {
		// 		path = strings.TrimSpace(path)
		// 		if path != "" {
		// 			fmt.Println("Include:", path)
		// 			content := getFileContent(path)
		// 			if content != "" {
		// 				fmt.Println(content)
		// 				// rawBlock := pf.RawBlock("format", content)
		// 				// rawBlock := pf.Str("xxx")
		// 				// para = append(para, rawBlock)
		// 			}
		// 		}
		// 	}
		// }
		// table.

		return pantable
		// log.Println(pantable)
		// return pf.Str("abcd")
	}
	return nil
}

func contains(s []interface{}, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getFileContent(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		// fmt.Print(err)
		return ""
	}
	return string(dat)
}

func main() {
	pf.ToJSONFilter(aetablefilter)
}
