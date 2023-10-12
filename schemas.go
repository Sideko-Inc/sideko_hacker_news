
package sideko_hacker_news

import "encoding/json"

func UnmarshalGetUpdatesJSONResponse(data []byte) (GetUpdatesJSONResponse, error) {
	var r GetUpdatesJSONResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetUpdatesJSONResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalItem(data []byte) (Item, error) {
	var r Item
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Item) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GetUpdatesJSONResponse struct {
	// Changed items            
	Items              []int64  `json:"items,omitempty"`
	// Changed profiles         
	Profiles           []string `json:"profiles,omitempty"`
}

type Item struct {
	By          string  `json:"by"`
	Dead        *bool   `json:"dead,omitempty"`
	Deleted     *bool   `json:"deleted,omitempty"`
	Descendants *int64  `json:"descendants,omitempty"`
	ID          int64   `json:"id"`
	Kids        []int64 `json:"kids,omitempty"`
	Parent      *int64  `json:"parent,omitempty"`
	Parts       []int64 `json:"parts,omitempty"`
	Poll        *int64  `json:"poll,omitempty"`
	Score       *int64  `json:"score,omitempty"`
	Text        *string `json:"text,omitempty"`
	Time        int64   `json:"time"`
	Title       *string `json:"title,omitempty"`
	Type        string  `json:"type"`
	URL         *string `json:"url,omitempty"`
}

type User struct {
	// The user's optional self-description. HTML            
	About                                        *string     `json:"about,omitempty"`
	// Creation date of the user, in Unix Time               
	Created                                      *int64      `json:"created,omitempty"`
	ID                                           *string     `json:"id,omitempty"`
	// The user's karma                                      
	Karma                                        *int64      `json:"karma,omitempty"`
	Submitted                                    interface{} `json:"submitted"`
}
