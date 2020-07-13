package nodeEngine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Node struct {
	Id      int           `json:"id"`
	Type    string        `json:"type"`
	Inputs  []interface{} `json:"inputs"`
	Outputs []interface{} `json:"outputs"`
}

type Output struct {
	Name string      `json:"name"`
	Type string      `json:"type"`
	Link interface{} `json:"link"`
}

type GraphX struct {
	Nodes []Node        `json:"nodes"`
	Links []interface{} `json:"links"`
}

func NewGraphByFile(filename string) *Graph {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	var dataObject GraphX
	err = json.Unmarshal(data, &dataObject)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	graph := NewGraph()

	for _, v := range dataObject.Nodes {

		var inputs []int
		for _, input := range v.Inputs {
			if mp, ok := input.(map[string]interface{}); ok {
				if link, ok := mp["link"]; ok {
					if link == nil {
						continue
					}
					inputs = append(inputs, int(link.(float64)))
				}
			}
		}
		//for _, link := range links {
		//	if link == nil {
		//		continue
		//	}
		//	inputs = append(inputs, link.(int))
		//}
		var outputs [][]int
		for _, output := range v.Outputs {
			if mp, ok := output.(map[string]interface{}); ok {
				if links, ok := mp["links"]; ok {
					if links == nil {
						continue
					}
					var tmp []int
					for _, link := range links.([]interface{}) {
						if link == nil {
							continue
						}
						tmp = append(tmp, int(link.(float64)))
					}
					outputs = append(outputs, tmp)
				}
			}
		}

		graph.AddBlock(v.Id, v.Type, inputs, outputs)
	}

	for _, v := range dataObject.Links {
		if arr, ok := v.([]interface{}); ok {
			graph.AddLink(
				int(arr[0].(float64)),
				int(arr[1].(float64)),
				int(arr[2].(float64)),
				int(arr[3].(float64)),
				int(arr[4].(float64)),
			)
		}
	}

	return graph
}
