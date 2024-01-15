package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

// objectならそれぞれの要素を再帰ソートをかけ、その後そのまま返す
// スライスならそれぞれの要素に再帰ソートをかけ、その後yamlエンコードした文字列の昇順でソートし、返す
// bool, int, stringならそのまま返す
func recursiveSort(obj interface{}) interface{} {
	switch typed := obj.(type) {
	// literalならそのまま
	case string, int, bool, nil:
		return obj

	// map[string]interface{}ならfor回す
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, v := range typed {
			result[k] = recursiveSort(v)
		}
		return result

	// []interface{}ならソート
	case []interface{}:
		ss := []interface{}{}
		for _, v := range typed {
			ss = append(ss, recursiveSort(v))
		}
		sort.Slice(ss, func(i, j int) bool {
			yi, _ := yaml.Marshal(&ss[i])
			yj, _ := yaml.Marshal(&ss[j])
			return string(yi) < string(yj)
		})
		return ss

	default:
		return obj
	}
}

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var y interface{}
	err = yaml.Unmarshal(b, &y)
	if err != nil {
		panic(err)
	}

	sorted := recursiveSort(y)

	output, err := yaml.Marshal(&sorted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", output)
}
