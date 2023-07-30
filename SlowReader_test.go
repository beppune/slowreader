package slowreader

import (
	"encoding/json"
	"strings"
	"testing"
)

/*func TestJsonValue(t *testing.T) {

	var buffer string
	for i := 0; i < 20; i++ {
		buffer += GiveAllarm()
	}
	fmt.Println(buffer)

	reader := NewSlowReader(128, strings.NewReader(buffer))

	decoder := json.NewDecoder(&reader)

	var allarm Allarm
	for decoder.More() {

		e := decoder.Decode(&allarm)
		if e != nil {
			t.Fatal(e)
		}

		fmt.Println(allarm)
	}
}*/

func TestArray(t *testing.T) {

	ar := GiveAllarmArray(20)

	reader := NewSlowReader(128, strings.NewReader(ar))

	decoder := json.NewDecoder(&reader)

	decoder.Token()

}
