package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessageheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessageheaderDud struct { }

// Messageheader
type Messageheader struct { }

// String returns a JSON representation of the model
func (o *Messageheader) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messageheader) MarshalJSON() ([]byte, error) {
    type Alias Messageheader

    if MessageheaderMarshalled {
        return []byte("{}"), nil
    }
    MessageheaderMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

