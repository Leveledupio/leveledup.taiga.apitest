package main

import (

	"fmt"
	"github.com/strongjz/taigo-test/taigo"

)



func main() {


/*	t := taigo.Project{
		Name: "Test2",
		Description:"words are hard",
		KanbanAct:false,
		IssuesAct:true,
		BacklogAct:true,
		Wiki:false,
		Private:true,
	}

	project,err := t.CreateProject()
	if err != nil {
		fmt.Printf("Error creating project %s", err)
	}


	fmt.Printf("Project Body: %v", string(*project))
*/

	s := taigo.Story{
		Description:"Story 2",
		Subject:"Subject for story 2",
		ProjectID:231729,
	}

	err := s.CreateStory()
	if err != nil {
		fmt.Printf("Error creating story %s", err)
	}


	fmt.Printf("Story ID: %v Subject%s", s.ID, s.Subject)
}