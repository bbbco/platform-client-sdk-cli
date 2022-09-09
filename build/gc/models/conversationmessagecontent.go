package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagecontentDud struct { 
    


    


    


    


    


    


    


    


    

}

// Conversationmessagecontent - Message content element. If contentType = \"Attachment\" only one item is allowed.
type Conversationmessagecontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // Location - Location content.
    Location Conversationcontentlocation `json:"location"`


    // Attachment - Attachment content.
    Attachment Conversationcontentattachment `json:"attachment"`


    // QuickReply - Quick reply content.
    QuickReply Conversationcontentquickreply `json:"quickReply"`


    // ButtonResponse - Button response content.
    ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`


    // Template - Template notification content.
    Template Conversationcontentnotificationtemplate `json:"template"`


    // Story - Ephemeral story content.
    Story Conversationcontentstory `json:"story"`


    // Card - Card content
    Card Conversationcontentcard `json:"card"`


    // Carousel - Carousel content
    Carousel Conversationcontentcarousel `json:"carousel"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagecontent

    if ConversationmessagecontentMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        Location Conversationcontentlocation `json:"location"`
        
        Attachment Conversationcontentattachment `json:"attachment"`
        
        QuickReply Conversationcontentquickreply `json:"quickReply"`
        
        ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`
        
        Template Conversationcontentnotificationtemplate `json:"template"`
        
        Story Conversationcontentstory `json:"story"`
        
        Card Conversationcontentcard `json:"card"`
        
        Carousel Conversationcontentcarousel `json:"carousel"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

