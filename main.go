package cra_env_go

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

import "encoding/json"

func main() {
	var sb strings.Builder
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.HasPrefix(pair[0], "REACT_APP_") {
			value, _ := json.Marshal(pair[1])
			sb.WriteString(fmt.Sprintf("process.env.%s = %s;\n", pair[0], value))
		}
	}
	vars := sb.String()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	replace := fmt.Sprintf(`
	<head>
      <script>
		process = process || {};
		process.env = process.env || {};
		%s
      </script>
	`, vars)

	var replaced = strings.Replace(string(data), "<head>", replace, -1)
	fmt.Print(replaced)
}
