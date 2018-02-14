package taiga

//Epic - Epic struct for taiga
type Epic struct {

	//Required
	ProjectID string `json:"project"`
	Subject   string `json:"subject"`

	//optional
	AssignedTo   string   `json:"assigned_to"`
	BlockedNotes string   `json:"blocked_note"`
	Description  string   `json:"description"`
	Blocked      bool     `json:"is_blocked"`
	Closed       bool     `json:"is_closed"`
	Color        string   `json:"color"`
	Tags         []string `json:"tags"`
	Watchers     []string `json:"watchers"`
}

//CreateEpic - Create an Epic
func (e *Epic) CreateEpic() {
}

//DeleteEpic - remove an Epic from taigo
func (e *Epic) DeleteEpic() {

}

//UpdateEpic - updates an Epic in taigo
func (e *Epic) UpdateEpic() {

}

//GetEpic - retrieves an Epic from taigo
func (e *Epic) GetEpic() {

}
