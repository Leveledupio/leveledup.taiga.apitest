package taiga

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
	)


//Story - Creates a story
//https://taigaio.github.io/taiga-doc/dist/api.html#user-stories-create
type Story struct {

	ID int `json:"id"`
	//Required
	ProjectID int `json:"project"`
	Subject string `json:"subject"`
	Description string `json:"description,omitempty"`

	//optional
	AssignTo string `json:"assigned_to,omitempty"`
	BacklogOrder int `json:"backlog_order,omitempty"`
	BlockedNote string `json:"blocked_note,omitempty"`
	ClientRequirement bool `json:"client_requirement,omitempty"`
	Blocked bool `json:"is_blocked,omitempty"`
	Closed bool `json:"is_closed,omitempty"`
	KanbanOrder int `json:"kanban_order,omitempty"`
	MileStoneID int `json:"milestone,omitempty"`
	Points map[string]int `json:"points,omitempty"`
	SprintOrder int `json:"sprint_order,omitempty"`
	Status int `json:"status,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TeamRequirement bool `json:"team_requirement,omitempty"`
	Watchers []int `json:"watchers,omitempty"`

}

func (s *Story) CreateStory() ( error){

	bodySend, err := json.Marshal(s)
	if err != nil{
		fmt.Printf("Error Marshelling Project JSON")
	}

	req, err := http.NewRequest("POST", baseURL + "/api/v1/userstories", bytes.NewBuffer(bodySend))
	if err != nil{
		return err
	}


	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return err

	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		return err
	}

	err = json.Unmarshal(respBody,s)
	if err != nil{
		fmt.Printf("Response body: %v", respBody)
		return err

	}

	return nil
}
