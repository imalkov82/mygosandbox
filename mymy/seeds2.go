package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	json_data, _ := json.Marshal(map[string]string{"vendor": "webint"})
	resp, _ := http.Post("http://localhost:8000", "application/json", bytes.NewBuffer(json_data))

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
