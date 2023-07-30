package slowreader

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestArray(t *testing.T) {

	ar := GiveAllarmArray(100)

	reader := NewSlowReader(1024, strings.NewReader(ar))

	decoder := json.NewDecoder(&reader)

	token, _ := decoder.Token()
	a := fmt.Sprintf("%v", token)

	if a == "[" {
		var allarm Allarm

		for decoder.More() {
			err := decoder.Decode(&allarm)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(allarm)
		}
	}
}
