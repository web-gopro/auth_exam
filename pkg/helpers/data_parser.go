package helpers

import (
	"encoding/json"
	"fmt"
)

func DataParser[t any, t2 any](src t, dst t2) {
	byData, err := json.Marshal(src)

	if err != nil {

	fmt.Println(err)
		return
	}

	json.Unmarshal(byData, dst)
}