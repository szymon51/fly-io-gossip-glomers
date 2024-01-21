package main

import (
	"encoding/json"
	"log"

	mealstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// input:
// {
// 	"src": "c1",
// 	"dest": "n1",
// 	"body": {
// 	  "type": "echo",
// 	  "msg_id": 1,
// 	  "echo": "Please echo 35"
// 	}

func main() {
	n := mealstrom.NewNode()

	n.Handle("echo", func(msg mealstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// update the message type to return back "type": "echo_ok"
		body["type"] = "echo_ok"

		// Echo the original message back with the updated message type and "in_reply_to" attribute
		return n.Reply(msg, body)

	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
