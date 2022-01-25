// vim:noet ts=4 sw=4 sts=4
package main

import (
	"os"
	"log"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func usage() {
	fmt.Printf("usage: %s slack_incoming_webhook_url\n", os.Args[0])
	fmt.Printf("\n")
	fmt.Printf("example:\n")
	fmt.Printf("    $ somecommand | %s https://hooks.slack.com/services/xxxxxxxxxxxx \n", os.Args[0])
	fmt.Printf("\n")

	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	stdin_buf, _ := ioutil.ReadAll(os.Stdin)

	if len(stdin_buf) == 0 {
		log.Printf("[*] nothing to do...\n")
		os.Exit(0)
	}

	url := os.Args[1]
	log.Printf("[+] url=%s\n", url)

	message := new(Message)
	message.Text = string(stdin_buf)

	json_str, _ := json.Marshal(message)
	log.Printf("[+] %s\n", string(json_str))

	// see also... https://api.slack.com/messaging/webhooks#posting_with_webhooks
	res, err := http.Post(url, "application/json", bytes.NewBuffer(json_str))
	defer res.Body.Close()

	if err != nil {
		log.Println("[!] " + err.Error())
		os.Exit(0)
	} 

	log.Println("[*] " + res.Status)
}
