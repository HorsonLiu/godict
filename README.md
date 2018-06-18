# Dict for golang

## Inspired by Dict in Python

前一段时间使用 python，发现 python 的 Dict() 真是一个好东西，使用非常方便，go语言相比而言就显得有些死板，必须要固定 struct 形式才能与 json 字符串互相转换，所以写下来这个 godict 。

## Wellcome to contribute

目前写的比较简单，基本可以覆盖一般情况，欢迎指正、欢迎合作！

## Usage example

```golang
package main

import(
	"fmt"
	"godict"
	"encoding/json"
)

type TestStruct struct {
	Mi int `json:"mi"`
	Mf float64 `json:"mf"`
	Mb bool `json:"mb"`
	Ms string `json:"ms"`
}

func main() {
	t := godict.Dict{}
	u := godict.Dict{}

	vInt := 24
	vFloat := 0.33
	vBool := true
	vString := "haha"
	vSliceInt := []int{1,2,3}
	vSliceFloat := []float32{1.1,2.2,3.3}
	vSliceBool := []bool{true,false}
	vSliceString := []string{"wow","hello","bye"}
	vStruct := TestStruct{}
	vSliceStruct := [2]TestStruct{}

	t["A"] =  u
	t["B"] = &vInt
	t["C"] = &vFloat
	t["D"] = &vBool
	t["E"] = &vString
	t["F"] = &vSliceInt
	t["G"] = &vSliceFloat
	t["H"] = &vSliceBool
	t["I"] = &vSliceString
	t["J"] = &vStruct
	t["K"] = &vSliceStruct
	
	u["a"] = vInt
	u["b"] = vFloat
	u["c"] = vBool
	u["d"] = vString
	u["e"] = vSliceInt
	u["f"] = vSliceFloat
	u["g"] = vSliceBool
	u["h"] = vSliceString
	u["i"] = vStruct
	u["j"] = vSliceStruct

	// fmt.Println("t=", t)

	println("\n======== dict to string ========\n")

	js,_ := json.Marshal(t)
	fmt.Printf(string(js))

	fmt.Println("t=", t)

	println("\n======== string to dict ========\n")

	ot := godict.Dict{}
	err := json.Unmarshal(js, &ot)
	if err == nil {
		fmt.Print(ot)
	} else {
		fmt.Println(err.Error())
	}
	
	println("\n\n======== end ========\n")
}

```
