// main.go
package main

import (
    "fmt"
//	"os"
    "io/ioutil"
)

type Entity struct {
    Name       string
    Attributes []Attribute
}

type Attribute struct {
    Name string
    Type string
}

func main() {
    entity := Entity{
        Name: "Person",
        Attributes: []Attribute{
            {Name: "ID", Type: "int"},
            {Name: "Name", Type: "string"},
            {Name: "Age", Type: "int"},
        },
    }

    if err := generateEntityFile(entity); err != nil {
        fmt.Println("Error creating entity file:", err)
        return
    }

    fmt.Printf("Entity file '%s.go' created successfully.\n", entity.Name)
}

func generateEntityFile(entity Entity) error {
    content := fmt.Sprintf(`package main

type %s struct {
`, entity.Name)

    for _, attr := range entity.Attributes {
        content += fmt.Sprintf("    %s %s\n", attr.Name, attr.Type)
    }

    content += "}\n"


	/*embeddedcontent := fmt.Sprintf(`
	type %s struct {
	`, entity.Name)
	
		for _, attr := range entity.Attributes {
			embeddedcontent += fmt.Sprintf("    %s %s\n", attr.Name, attr.Type)
		}

		embeddedcontent += "}\n" 
        
        	content = content+embeddedcontent
        */

	
	

    filename := entity.Name + ".go"
    return ioutil.WriteFile(filename, []byte(content), 0644)
}


