package template

import (
	"encoding/json"
	"fmt"

	"github.com/leekchan/accounting"
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig/escape"
)

// jsonEncode converts the passed in value into a JSON string.
func jsonEncode(_ stick.Context, val stick.Value, _ ...stick.Value) stick.Value {
	v, err := json.Marshal(val)
	if err != nil {
		// TODO: Do something useful
		return ""
	}
	return string(v)
}

func urlEncode(_ stick.Context, val stick.Value, _ ...stick.Value) stick.Value {
	return escape.URLQueryParam(stick.CoerceString(val))
}

func format(_ stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	if len(args) != 1 {
		return ""
	}
	if v, ok := args[0].(string); ok {
		return fmt.Sprintf(v, val)
	}
	return ""
}

func money(_ stick.Context, val stick.Value, _ ...stick.Value) stick.Value {
	ac := accounting.Accounting{Symbol: "", Precision: 2}
	return "<span class=\"isk\">Ƶ</span> " + ac.FormatMoney(val)
}
