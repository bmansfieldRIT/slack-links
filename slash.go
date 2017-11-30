package slash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// read the request parameter command
	command := r.FormValue("command")
	url := "https://slack.com/api/channels.list?token=xoxp-273143757027-272563939744-276827556560-fa6dda16a651751e1c8c466fb47c23a3"
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
	//fmt.Fprint(w, resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err1 := json.Unmarshal([]byte(body), &data)
	if err1 != nil {
		fmt.Fprint(w, err1)
	} else {
		// need to assert that the second indexing is on a map, otherwise, go will complain that a struct can't be indexed
		fmt.Fprint(w, data["channels"].(map[string]interface{})["is_member"])
	}

	/*
	       resp, err := http.Get(url)
	   	if err != nil {
	   		fmt.Fprint(w, err)
	   	}

	   	defer resp.Body.Close()
	   	body, _ := ioutil.ReadAll(resp.Body)
	   	//var respObj responseRtmStart
	   	//err = json.Unmarshal(body, &respObj)
	*/
	// ideally to other checks for tokens/usernames/etc
	if command == "/links" {
		//fmt.Fprint(w, "command recieved")
		//fmt.Fprint(w, body)
	} else {
		fmt.Fprint(w, "I don't understand your command")
	}
}
