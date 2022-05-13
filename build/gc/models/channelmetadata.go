package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChannelmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChannelmetadataDud struct { }

// Channelmetadata - Information about the channel.
type Channelmetadata struct { }

// String returns a JSON representation of the model
func (o *Channelmetadata) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Channelmetadata) MarshalJSON() ([]byte, error) {
    type Alias Channelmetadata

    if ChannelmetadataMarshalled {
        return []byte("{}"), nil
    }
    ChannelmetadataMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

