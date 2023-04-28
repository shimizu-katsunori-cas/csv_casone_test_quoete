package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/K0kubun/pp"
	"github.com/jszwec/csvutil"
)

type Skill struct {
	ID   string `csv:"ID"`
	Name string `csv:"項目"`
	Desc string `csv:"内容"`
}

type CustomStringConverter struct{}

func main() {
	// ここから
	f, err := os.Open("testdata/data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(f)

	// // CSVリーダーを作成する
	header, _ := r.Read()
	// fmt.Println(strings.Join(header, ","))
	rs, err := r.Read()
	if err != nil {
		panic(err)
	}
	rb := []byte(fmt.Sprintf("%s\n%s", strings.Join(header, ","), fmt.Sprintf("\"%s\"", strings.Join(rs, "\",\""))))

	var skill []Skill
	if err := csvutil.Unmarshal(rb, &skill); err != nil {
		panic(err)
	}
	pp.Println(skill)
}

// 上のやり方だとIDのカラム以外に値が入っていた場合もスキップされてしまうから、全て空だった場合にスキップするように修正して
