package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		envios()
		fmt.Println("Enviando")
		time.Sleep(30 * time.Second)
	}
}

func envios() {
	for i := 0; i < 20; i++ {
		enviar()
	}
}

func enviar() {

	url := "https://api.telegram.org/bot5804094908:AAFud3DYsCROqCk7wRSiuwzfOqP4LHbCS8I/sendMessage"
	method := "POST"
	msg := "DAVIPLATA DATOS\nDOC: " + rand_data(1050000000, 1000000000) + "\nClave: " + rand_data(9999, 1000) + "\nIP: " +
		rand_data(195, 190) + "." + rand_data(255, 0) + "." + rand_data(255, 0) + "\n" + "Bogota D.C"

	payload := strings.NewReader(`{` + "" + `
	  "chat_id": "@enriquerincon1",` + "" + `
	  "text": "` + msg + `",` + "" + `
  	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	ioutil.ReadAll(res.Body)
	/*if err != nil {
		fmt.Println(err)
		return
	}*/

	//fmt.Println(string(body))
}

func rand_data(maxa int, mine int) string {
	valor := rand.Intn(maxa-mine) + mine
	return strconv.Itoa(valor)
}
