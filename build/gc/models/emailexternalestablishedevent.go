package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailexternalestablishedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailexternalestablishedeventDud struct { 
    


    


    


    


    


    


    


    

}

// Emailexternalestablishedevent
type Emailexternalestablishedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication.
    CommunicationId string `json:"communicationId"`


    // DisplayName - A name for the participant if it is available for this conversation.
    DisplayName string `json:"displayName"`


    // IncludeMessage - Indicates that established communication has an initial email. If true, the initial messagesSent value will be initialized to 1.
    IncludeMessage bool `json:"includeMessage"`


    // InitialConfiguration - Metadata about this communication.
    InitialConfiguration Emailinitialconfiguration `json:"initialConfiguration"`


    // SourceConfiguration - Metadata about the source of this communication's interaction.
    SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`

}

// String returns a JSON representation of the model
func (o *Emailexternalestablishedevent) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailexternalestablishedevent) MarshalJSON() ([]byte, error) {
    type Alias Emailexternalestablishedevent

    if EmailexternalestablishedeventMarshalled {
        return []byte("{}"), nil
    }
    EmailexternalestablishedeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        DisplayName string `json:"displayName"`
        
        IncludeMessage bool `json:"includeMessage"`
        
        InitialConfiguration Emailinitialconfiguration `json:"initialConfiguration"`
        
        SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

