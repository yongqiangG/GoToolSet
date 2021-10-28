package json

import (
	"fmt"
	"github.com/tidwall/pretty"
)

func JSONPretty(s string, color bool) string {
	if color {
		return fmt.Sprintf("%s\n", pretty.Color(pretty.Pretty([]byte(s)), pretty.TerminalStyle))
	} else {
		return fmt.Sprintf("%s\n", pretty.Pretty([]byte(s)))
	}
}
