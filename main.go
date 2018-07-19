package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var key = ""
var secret = ""

const fatsecretURL = "http://platform.fatsecret.com/rest/server.api"

func main() {
	loadSecrets()
	fmt.Println("key:", key, "secret:", secret)
	fmt.Println("Fat Secret API")
	res, err := http.Get(fatsecretURL)
	no(err)
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
}

func loadSecrets() {
	f, err := os.Open("ignore/key.txt")
	no(err)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	key = scanner.Text()
	scanner.Scan()
	secret = scanner.Text()
}

func no(err error) {
	if err != nil {
		panic(err)
	}
}
