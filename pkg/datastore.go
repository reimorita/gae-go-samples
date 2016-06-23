package pkg

import (
	"net/http"
	"time"
	"appengine"
	"appengine/datastore"
	"fmt"
	"encoding/json"
)

func init() {
	http.HandleFunc("/putEntity", putEntity)
	http.HandleFunc("/getEntity", getEntity)
}

type User struct {
	FirstName	string	`datastore:"firstName"`
	LastName	string	`datastore:"lastName"`
	Age	int	`datastore:"age"`
	RegisterDate	time.Time	`datastore:"registerDate"`
}


func putEntity(w http.ResponseWriter, r *http.Request) {
	// Context
	c := appengine.NewContext(r)
	// Create Entity
	user := User{
		FirstName: "Taro",
		LastName: "Yamada",
		RegisterDate:	time.Now(),
	}

	// Create Key By Using string key
	key := datastore.NewKey(c, "User", "stringKey", 0, nil)
	// Put entity to Datastore
	_, err := datastore.Put(c, key, &user)
	// If Error
	if err != nil {
		c.Errorf("Fail to put entity: %v", err)
		return
	}
}

func getEntity(w http.ResponseWriter, r *http.Request) {
	// Context
	c := appengine.NewContext(r)
	// Create Key
	key := datastore.NewKey(c, "User", "stringKey", 0, nil)
	var user User
	// Get Entity
	if err := datastore.Get(c, key, &user); err != nil {
		c.Errorf("Fail to get entity: %v", err)
		return
	}
	resultJson, err :=  json.Marshal(user)
	if err != nil {
		c.Errorf("Fail to create Json: %v", err)
		return
	}
	fmt.Fprintf(w, "%s", resultJson)
}


