package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
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

		if err := generateEntityFile(entityMetaData); err != nil {
			fmt.Println("Error creating entity file:", err)
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


func generateEntityFile(entityMetaData EntityMetaData) error {
	content := fmt.Sprintf(`package main
	
type %s struct {
	`, entityMetaData.EntityName)

	content += fmt.Sprintf(`package main
	
	type %s struct {
		`, entityMetaData.EntityName)
	
	for _, attr := range entityMetaData.Attributes {

		if(attr.Type=="text"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "string")
		}
		if(attr.Type=="number"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "int")
		}
		if(attr.Type=="date"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "time.Time")
		}
		if(attr.Type=="dateTime"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "time.Time")
		}
		if(attr.Type=="boolean"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "bool")
		}
		if(attr.Type=="longtext"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "string")
		}
		if(attr.Type=="file"){
			content += fmt.Sprintf("    %s %s\n", attr.Name, "string")
		}
		//attrType="string"
       
    }
	
	content += "}\n"
	

    filename := entityMetaData.EntityName + ".go"
    return ioutil.WriteFile(filename, []byte(content), 0644)
}


func main() {
	http.HandleFunc("/postendpoint", handlePostRequest)
	http.ListenAndServe(":8080", nil)
}
