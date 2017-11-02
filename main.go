package main

import (

	"fmt"
	"github.com/strongjz/taigo-test/taigo"

)



func main() {


	t := taigo.Project{
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
		fmt.Printf("Error creating project")
	}


	fmt.Printf("Project Body: %v", string(*project))

}