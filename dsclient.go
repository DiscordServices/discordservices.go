// Copyrights mikeuwu 2020-21

package discordservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// const notoken := errors.New()

type DSClient struct {
	token string
	botid int
}

func (c DSClient) postShardCOunt(count int) {
	// Post stats to discordservices api
	// Required arguments: count int -> shard count
	stred := strconv.Itoa(c.botid)
	var url = "https://api.discordservices.net/bot/" + stred + "/stats"
	if count == 0 {
		fmt.Println("[ERROR] - No count provided")
	} else {
		var jsonBody, parserr = json.Marshal(map[string]int{"shards": count})
		if parserr != nil {
			log.Fatalln(parserr)
		}
		var resp, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Fatalln(err)
		}
		resp.Header.Set("Content-Type", "application/json")
		resp.Header.Set("Authorization", c.token)
		var client = &http.Client{}
		var req, httperr = client.Do(resp)
		if httperr != nil {
			log.Fatalln(httperr)
		}
		defer req.Body.Close()
		var body = req.Body
		log.Println("Responce status: %v\n", req.Status)
		log.Println("Responce body: %v\n", body)
	}

}

func (c DSClient) postGuildCount(count int) {
	// Post stats to discordservices api
	// Required arguments: count int -> Guild count
	stred := strconv.Itoa(c.botid)
	var url = "https://api.discordservices.net/bot/" + stred + "/stats"
	if count == 0 {
		fmt.Println("[ERROR] - No count provided")
	} else {
		var jsonBody, parserr = json.Marshal(map[string]int{"servers": count})
		if parserr != nil {
			log.Fatalln(parserr)
		}
		var resp, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Fatalln(err)
		}
		resp.Header.Set("Content-Type", "application/json")
		resp.Header.Set("Authorization", c.token)
		var client = &http.Client{}
		var req, httperr = client.Do(resp)
		if httperr != nil {
			log.Fatalln(httperr)
		}
		defer req.Body.Close()
		//var body = req.Body
		fmt.Println("Responce status: ", req.Status)
		//log.Println("Responce body: %v\n", body)
	}

}
