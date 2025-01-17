package taiga

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Project - creates a project in taigo
type Project struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	BacklogAct  bool   `json:"is_backlog_activated"`
	IssuesAct   bool   `json:"is_issues_activated"`
	KanbanAct   bool   `json:"is_kanban_activated"`
	Wiki        bool   `json:"is_wiki_activated"`
	Private     bool   `json:"is_private"`
}

//CreateProject - Create a project
func (p *Project) CreateProject() (*[]byte, error) {

	bodySend, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error Marshelling Project JSON")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/projects", baseURL), bytes.NewBuffer(bodySend))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

//DeleteProject - remove project from taigo
func (p *Project) DeleteProject() {

}

//UpdateProject - updates a project in taigo
func (p *Project) UpdateProject() {

}

//GetProject - retrieves a project from taigo
func (p *Project) GetProject() {

}
