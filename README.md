# guard

Dead simple json marshal guard for potential data type mismatch (eg. marshal array into struct, string into object).

## Usage

[See Doc](https://pkg.go.dev/github.com/Sora233/guard)

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sora233/guard"
)

func main() {
	var response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data guard.Guard[struct {
			Content string `json:"content"`
		}] `json:"data"`
	}
	var body = `{"code":10000, "msg":"xx error", "data": "wrong format"}`

	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println("can not unmarshal body", err.Error())
		return
	}
	if response.Data.IsSuccess() {
		fmt.Println("content is", response.Data.Get().Content)
	} else {
		fmt.Println("can not get content because of wrong body data", string(response.Data.GetRaw()))
	}
}

```