package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func genWord() {
	str := ""
	for i := 0; i < 10; i++ {
		str += string(rune(rand.Intn(26) + 97))
	}
	word = str
	clients.sendToAll(Info{str, 0})
}

func verify(word string, num int) bool {
	a := sha256.Sum256([]byte(word + strconv.Itoa(num)))
	return string(a[:3]) == "000"
}

func raceHash() {
	for {
		for name, conn := range clients.users {
			var data Info
			err := conn.ReadJSON(&data)
			if err != nil {
				delete(clients.users, name)
				fmt.Println(name + " has disconnected")
			}
			if verify(data.Word, data.Num) {
				clients.sendToClient(Info{name, accounts[name][2].(int) + 10}, name)
				clients.sendToAll(Info{"newWord", 0})
				time.Sleep(time.Second * 10)
				genWord()
			}
		}
	}
}
