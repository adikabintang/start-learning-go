package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json: "page"`
	Fruits []string `json: "fruits"`
}

func main() {
	/*
		"marshal" is the process of transforming memory representation of an object to a data format
		suitable for storage or transmission
	*/
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	strB, _ := json.Marshal("nyot")
	fmt.Println(string(strB))

	arr := []string{"oh", "my", "zsh"}
	arrB, _ := json.Marshal(arr)
	fmt.Println(string(arrB))

	aMap := map[string]int{"apple": 5, "mango": 4}
	mapB, _ := json.Marshal(aMap)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"men", "kancut"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{Page: 1, Fruits: []string{"oi", "et"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num": 6.13, "strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"duren": 4}
	enc.Encode(d)
}
