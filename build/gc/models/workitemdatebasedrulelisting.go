package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemdatebasedrulelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemdatebasedrulelistingDud struct { 
    


    


    


    


    

}

// Workitemdatebasedrulelisting
type Workitemdatebasedrulelisting struct { 
    // Entities
    Entities []Workitemdatebasedrule `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Workitemdatebasedrulelisting) String() string {
     o.Entities = []Workitemdatebasedrule{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemdatebasedrulelisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemdatebasedrulelisting

    if WorkitemdatebasedrulelistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemdatebasedrulelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemdatebasedrule `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemdatebasedrule{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

