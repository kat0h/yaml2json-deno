package main

import (
	"syscall/js"
	"bytes"
	"encoding/json"
	"github.com/ghodss/yaml"
)

func main() {
	js.Global().Set("y2j", js.FuncOf(y2j_js))
	c := make(chan struct{}, 0)
	<-c
}

func y2j_js(this js.Value, yaml_s []js.Value) interface{} {
	return yaml2json(yaml_s[0].String())
}

func yaml2json(yml_s string) string {
	buf := []byte(yml_s)
	m := make(map[string]interface{})
	y_err := yaml.Unmarshal(buf, &m)
	if y_err != nil {
		panic(y_err)
	}
	d, j_err := json.Marshal(&m)
	if j_err != nil {
		panic(j_err)
	}
	var b bytes.Buffer
	i_err := json.Indent(&b, d, "", "  ")
	if i_err != nil {
		panic(i_err)
	}
	return b.String()
}
