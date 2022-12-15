package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := func() int { return 100 }
	fmt.Printf("%v\n", f())
	fmt.Printf("%v\n", &f)

	map1 := map[string]interface{}{}
	map1["foo"] = "bar"
	map1["hoge"] = "fuga"
	map1["int"] = 10
	fmt.Printf("%v\n", map1)

	map2 := map[string]interface{}{}
	map2["foo"] = "bar"
	map2["hoge"] = "fuga"
	map2["int"] = 10
	fmt.Printf("%v\n", map2)

	var comp bool

	comp = reflect.DeepEqual(map1, map2)
	fmt.Printf("%v\n", comp)

	map1["func"] = f
	map2["func"] = f
	comp = reflect.DeepEqual(map1, map2)
	fmt.Printf("%v\n", comp)

	f1 := map1["func"]
	f2 := map2["func"]
	fmt.Printf("f1 = %t, f2 = %t\n", f1, f2)
	fmt.Printf("f1() = %v, f2() = %v\n", f1.(func() int)(), f2.(func() int)())

	map1["func"] = &f
	map2["func"] = &f
	comp = reflect.DeepEqual(map1, map2)
	fmt.Printf("%v\n", comp)

	f3 := map1["func"]
	f4 := map2["func"]
	fmt.Printf("f3 = %t, f4 = %t\n", *f3.(*func() int), *f4.(*func() int))

	f5 := *f3.(*func() int)
	f6 := *f4.(*func() int)
	fmt.Printf("f5() = %v, f6() = %v\n", f5(), f6())

	fmt.Println("=== toml Mershal() ===")

	var valueSection = &ValueSection{
		IntVal:  999,
		StrVal:  "foo",
		BoolVal: false,
		listVal: []string{"foo", "bar", "hoge", "fuga"},
	}

	root := &Root{Value: *valueSection}

	result, err := Marshal(root)
	if nil != err {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("%v", result)

		defer func() {
			e := recover()
			if nil != e {
				fmt.Printf("%v\n", e)
			}
		}()

		fmt.Println("=== toml UnMershal() ===")

		root := UnMarshal(result)
		fmt.Printf("%v\n", root)
	}
}

type Root struct {
	Value ValueSection `toml:"STRUCT_VALUES"`
}

type ValueSection struct {
	IntVal  int      `toml:"intVal"`  // 数値
	StrVal  string   `toml:"strVal"`  // 文字列
	BoolVal bool     `toml:"bool"`    // 論理値
	listVal []string `toml:"listVal"` // リスト
}
