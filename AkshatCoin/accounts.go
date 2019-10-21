package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

//Username;Password;Coin;Level
func loadAccounts() map[string][3]interface{} {
	bytes, err := ioutil.ReadFile("accounts.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	acc := make(map[string][3]interface{})
	for _, line := range lines {
		items := strings.Split(line, ";")
		coins, _ := strconv.Atoi(items[2])
		acc[items[0]] = [3]interface{}{items[1], coins, items[3]}
	}
	return acc
}

func newAccount(username, password, level string, coins int) {
	accounts[username] = [3]interface{}{password, coins, level}
}

func saveAccounts() {
	lines := ""
	for username, data := range accounts {
		lines += strings.Join([]string{username, data[0].(string), strconv.Itoa(data[1].(int)), data[2].(string)}, ";") + "\n"
	}
	ioutil.WriteFile("accounts.txt", []byte(lines[:len(lines)-1]), 0644)
}
