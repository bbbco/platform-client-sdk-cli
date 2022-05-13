package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgelineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgelineDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Edgeline
type Edgeline struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


    // Schema
    Schema Domainentityref `json:"schema"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // Edge
    Edge Edge `json:"edge"`


    // EdgeGroup
    EdgeGroup Edgegroup `json:"edgeGroup"`


    // LineType
    LineType string `json:"lineType"`


    // Endpoint
    Endpoint Endpoint `json:"endpoint"`


    // IpAddress
    IpAddress string `json:"ipAddress"`


    // LogicalInterfaceId
    LogicalInterfaceId string `json:"logicalInterfaceId"`


    

}

// String returns a JSON representation of the model
func (o *Edgeline) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeline) MarshalJSON() ([]byte, error) {
    type Alias Edgeline

    if EdgelineMarshalled {
        return []byte("{}"), nil
    }
    EdgelineMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        Schema Domainentityref `json:"schema"`
        
        Properties map[string]interface{} `json:"properties"`
        
        Edge Edge `json:"edge"`
        
        EdgeGroup Edgegroup `json:"edgeGroup"`
        
        LineType string `json:"lineType"`
        
        Endpoint Endpoint `json:"endpoint"`
        
        IpAddress string `json:"ipAddress"`
        
        LogicalInterfaceId string `json:"logicalInterfaceId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

