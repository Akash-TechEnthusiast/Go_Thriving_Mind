package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	//"os"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "KGM@123$"
    dbname = "postgres"
	schema = "next_gen_app"
)


type EntityMetaData struct {
	ID         string `json:"_id"`
	EntityName string `json:"entityName"`
	Version    string `json:"version"`
	ResyncByProcess bool   `json:"resyncByProcess"`
	Class           string `json:"_class"`
	Attributes []struct {
		ID          string `json:"_id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		DisplayName string `json:"displayName,omitempty"`
		Validations []struct {
			ValidationType string `json:"validationType"`
		} `json:"validations"`
		Related      bool   `json:"related"`
		Modified     bool   `json:"modified"`
		MappedAsLOV  bool   `json:"mappedAsLOV"`
		MappedAsREF  bool   `json:"mappedAsREF"`
		Encrypted    bool   `json:"encrypted"`
		Audited      bool   `json:"audited"`
		Currency     bool   `json:"currency"`
		DefaultValue string `json:"defaultValue,omitempty"`
	} `json:"attributes"`
	Relationships struct {
		WeakRelationship []struct {
			ID               string `json:"_id"`
			EntityID         string `json:"entityId"`
			RelationshipName string `json:"relationshipName"`
			RelationshipType string `json:"relationshipType"`
			Attributes       []struct {
				Name        string `json:"name"`
				Type        string `json:"type"`
				DisplayName string `json:"displayName"`
			} `json:"attributes"`
			Mandatory bool `json:"mandatory"`
			JSON      bool `json:"json"`
		} `json:"weakRelationship"`
		StrongRelationship []struct {
			ID               string `json:"_id"`
			EntityID         string `json:"entityId"`
			RelationshipName string `json:"relationshipName"`
			RelationshipType string `json:"relationshipType"`
			Attributes       []struct {
				Name        string `json:"name"`
				Type        string `json:"type"`
				DisplayName string `json:"displayName"`
			} `json:"attributes"`
			Mandatory bool `json:"mandatory"`
			JSON      bool `json:"json"`
		} `json:"strongRelationship"`
	} `json:"relationships"`
	PersistenceRules struct {
		PrimaryKeys []struct {
			AttributeName string `json:"attributeName"`
		} `json:"primaryKeys"`
		ForeignKeys           []interface{} `json:"foreignKeys"`
		UniqueKeys            []interface{} `json:"uniqueKeys"`
		NotNullableKeys       []interface{} `json:"notNullableKeys"`
		CompositeKeys         []interface{} `json:"compositeKeys"`
		AutoCodeGenContraints []struct {
			KeyAttribute string `json:"keyAttribute"`
			Prefix       string `json:"prefix"`
		} `json:"autoCodeGenContraints"`
	} `json:"persistenceRules"`
	Type        string `json:"type"`
	ContainerID string `json:"containerId"`
	FolderID    string `json:"folderId"`
	Audited     bool   `json:"audited"`
	Translate   bool   `json:"translate"`
	JSON        bool   `json:"json"`
	PRules      struct {
		ID                    string        `json:"_id"`
		ParticipatingEntities []string      `json:"participatingEntities"`
		EntityID              string        `json:"entityId"`
		EntityPrefix          string        `json:"entityPrefix"`
		InheritDefaults       bool          `json:"inheritDefaults"`
		PresentationRules     []interface{} `json:"presentationRules"`
		SubscriptionEventMap  struct {
		} `json:"subscriptionEventMap"`
		ParentEventMap struct {
		} `json:"parentEventMap"`
		Mandatory                 bool          `json:"mandatory"`
		Visible                   bool          `json:"visible"`
		Order                     int           `json:"order"`
		OnLoadRequired            bool          `json:"onLoadRequired"`
		OnCancelRequired          bool          `json:"onCancelRequired"`
		FormComputations          []interface{} `json:"formComputations"`
		FormAttributeComputations []interface{} `json:"formAttributeComputations"`
		ParticipatingProcess      []interface{} `json:"participatingProcess"`
		DynamicForm               bool          `json:"dynamicForm"`
	} `json:"pRules"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var entityMetaData EntityMetaData
		//var person Person

		// Decode the JSON data from the request body into the Person struct
		err := json.NewDecoder(r.Body).Decode(&entityMetaData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		//fileName := entityMetaData.EntityName + ".go"

		// Attempt to remove the file
	/*	errcode := os.Remove(fileName)
		if errcode != nil {
			fmt.Println("Error:", errcode)
			return
		} */
		
		/*if err := generateEntityFile(entityMetaData); err != nil {
			fmt.Println("Error creating entity file:", err)
			return
		}*/
		tablerr := createtableForEntity(entityMetaData);

		if  tablerr != nil {
			fmt.Println("Error creating entity file:", tablerr)
			return
		}
		fmt.Printf("Entity file '%s.go' created successfully.\n", entityMetaData.EntityName)

		// You can now access the person's data and perform any desired actions
		fmt.Println("Received data:", entityMetaData.Relationships.WeakRelationship)

		fmt.Fprintf(w, "Received POST request successfully!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func createtableForEntity(entityMetaData EntityMetaData) error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    // Open a connection to the database.
    db, err := sql.Open("postgres", psqlInfo)
	
	
	if err != nil {
		panic(err)
	}
	defer db.Close()

     dropTable:=fmt.Sprintf(`DROP TABLE IF EXISTS next_gen_app.%s`,entityMetaData.ID)
	_, err = db.Exec(dropTable)
	if err != nil {
		panic(err)
	}

	// Replace "json_data" with your desired table name
	

	content := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS next_gen_app.%s (
id SERIAL PRIMARY KEY,`,entityMetaData.ID)
		content += "\n"

	// Define the SQL statement to create the table

		for _, attr := range entityMetaData.Attributes {

	
		    if(attr.Name!="id"){
				//content += fmt.Sprintf("    %s  %s\n", attr.Name, "INT,")
			
				if(attr.Type=="text"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "TEXT,")
				} else if(attr.Type=="number"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "INT,")
				} else if(attr.Type=="date"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "DATE,")
				} else if(attr.Type=="dateTime"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "TIMESTAMP,")
				} else if(attr.Type=="boolean"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "TEXT,")
				} else if(attr.Type=="longtext"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "TEXT,")
				} else if(attr.Type=="file"){
					content += fmt.Sprintf("    %s  %s\n", attr.Name, "TEXT,")
				}
			}
		   
		}
		for _, weakrel := range entityMetaData.Relationships.WeakRelationship {

		
			if(weakrel.RelationshipType=="ManyToOne"){
	
			//	content += fmt.Sprintf("    %s  %s\n", weakrel.Attributes[0].Name, "TEXT,")
	
			}
			if(weakrel.RelationshipType=="OneToMany"){
	
				// content += fmt.Sprintf("    %s  %s\n", attr.Name, "TEXT,")
				create_thirdtable_For_onetomany_Entity(weakrel.EntityID,entityMetaData.ID)
	
			}
		}
	
	
		for _, strongrel := range entityMetaData.Relationships.StrongRelationship {
	
	
			if(strongrel.RelationshipType=="OneToOne"){
	
				content += fmt.Sprintf("    %s  %s\n", strongrel.Attributes[0].Name, "TEXT,")
	
			}
			if(strongrel.RelationshipType=="OneToMany"){
	
				//content += fmt.Sprintf("    %s  %s\n", attr.Name, "TEXT,")
				create_thirdtable_For_onetomany_Entity(strongrel.EntityID,entityMetaData.ID)
	
			}
		}
		//fmt.Printf("table", content)
		content = content[:len(content)-2]
		content += ")"

		fmt.Printf("Table %s created successfully and content  ==> %s.\n ", entityMetaData.ID,content)

	_, err = db.Exec(content)

	if (err != nil) {
		panic(err)
	}

	fmt.Printf("Table %s created successfully.\n", entityMetaData.ID)
	return err
}

func create_thirdtable_For_onetomany_Entity(childId string ,parentId string) error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    // Open a connection to the database.
    db, err := sql.Open("postgres", psqlInfo)
	
	
	if err != nil {
		panic(err)
	}
	defer db.Close()

	thirdtableName:=parentId+"id__"+childId+"id"
	//parentTable=parentId+"id"
	//childTable=strongrel.EntityID+"id"

	// Replace "json_data" with your desired table name
	
		content := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS next_gen_app.%s (
		%s,%s`,thirdtableName,parentId+"id TEXT",childId+"id TEXT)")
	

		fmt.Printf("Table %s created successfully and content  ==> %s.\n ",content)

	_, err = db.Exec(content)

	if (err != nil) {
		panic(err)
	}

	//fmt.Printf("Table %s created successfully.\n", entityMetaData.ID)
	return err
}


func generateEntityFile(entityMetaData EntityMetaData) error {
	content := fmt.Sprintf(`package main`)

	content += "\n"
	content += "\n"


	content += fmt.Sprintf(`import (
		"time"
		"fmt"
)`)

	content += fmt.Sprintf(`

type %s struct {
	`, entityMetaData.EntityName)

	content += "\n"

	for _, attr := range entityMetaData.Attributes {

	
    	if(attr.Type=="text"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "string")
		} else if(attr.Type=="number"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "int")
		} else if(attr.Type=="date"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "time.Time")
		} else if(attr.Type=="dateTime"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "time.Time")
		} else if(attr.Type=="boolean"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "bool")
		} else if(attr.Type=="longtext"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "string")
		} else if(attr.Type=="file"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "string")
		}
		
       
    }


	for _, weakrel := range entityMetaData.Relationships.WeakRelationship {

		lowercase_entityId :=convertToLowerCase(weakrel.EntityID)
		if(weakrel.RelationshipType=="ManyToOne"){

			content += fmt.Sprintf("    %s %s\n", weakrel.EntityID,lowercase_entityId)

		}
	    if(weakrel.RelationshipType=="OneToMany"){

			content += fmt.Sprintf("    %s %s\n", weakrel.EntityID, "[]"+lowercase_entityId)

		}
	}


	for _, strongrel := range entityMetaData.Relationships.StrongRelationship {

		lowercase_entityId :=convertToLowerCase(strongrel.EntityID)
		if(strongrel.RelationshipType=="OneToOne"){

			content += fmt.Sprintf("    %s %s\n", strongrel.EntityID, lowercase_entityId)

		}
	    if(strongrel.RelationshipType=="OneToMany"){

			content += fmt.Sprintf("    %s %s\n", strongrel.EntityID, "[]"+lowercase_entityId)

		}
	}
	
	content += "}\n"
	

    filename := entityMetaData.EntityName + ".go"
    return ioutil.WriteFile(filename, []byte(content), 0644)
}


func main() {
	
	http.HandleFunc("/postendpoint", handlePostRequest)
	http.ListenAndServe(":8080", nil)
	
}

func convertToLowerCase(entityId string) string {
	lowercaseFirst := strings.ToLower(entityId[:1]) + entityId[1:]

	return lowercaseFirst
}


