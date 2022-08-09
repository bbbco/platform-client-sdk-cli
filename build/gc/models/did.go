package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DidDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Did
type Did struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // PhoneNumber
    PhoneNumber string `json:"phoneNumber"`


    // DidPool
    DidPool Domainentityref `json:"didPool"`


    // Owner - A Uri reference to the owner of this DID, which is either a User or an IVR
    Owner Domainentityref `json:"owner"`


    // OwnerType
    OwnerType string `json:"ownerType"`


    

}

// String returns a JSON representation of the model
func (o *Did) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Did) MarshalJSON() ([]byte, error) {
    type Alias Did

    if DidMarshalled {
        return []byte("{}"), nil
    }
    DidMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        DidPool Domainentityref `json:"didPool"`
        
        Owner Domainentityref `json:"owner"`
        
        OwnerType string `json:"ownerType"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

