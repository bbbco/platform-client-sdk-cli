package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AiscoringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AiscoringDud struct { }

// Aiscoring
type Aiscoring struct { }

// String returns a JSON representation of the model
func (o *Aiscoring) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aiscoring) MarshalJSON() ([]byte, error) {
    type Alias Aiscoring

    if AiscoringMarshalled {
        return []byte("{}"), nil
    }
    AiscoringMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

