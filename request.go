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

var id_estafador string
var chat_id string

func main() {
	fmt.Println("Ingresa el id del estafador")
	fmt.Scanln(&id_estafador)
	fmt.Println("Ingresa el chat id")
	fmt.Scanln(&chat_id)

	for i := 0; i < 100000; i++ {
		enviados := envios()
		fmt.Println("Ciclo: " + strconv.Itoa(i+1) + "," + strconv.Itoa(enviados))
		time.Sleep(30 * time.Second)
	}
}

func envios() int {
	var enviados = 0
	for i := 0; i < 20; i++ {
		estate := enviar()
		if estate == true {
			enviados++
		}
	}
	return enviados
}

func enviar() bool {
	var est = false
	url := "https://api.telegram.org/bot" + id_estafador + "/sendMessage"
	method := "POST"
	ip_fake := rand_data(195, 190) + "." + rand_data(255, 0) + "." + rand_data(255, 0)
	fake_clave := rand_data(9999, 1000)
	fake_cedula := rand_data(1050000000, 1000000000)
	msg := "DAVIPLATA DATOS\nDOC: " + fake_cedula + "\nClave: " + fake_clave + "\nIP: " + ip_fake + "\n" + "Bogota D.C"

	payload := strings.NewReader(`{` +
		`"chat_id": "` + chat_id + `",` +
		`"text": "` + msg + `",` +
		`}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return est
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return est
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return est
	} else {
		est = true
	}

	//fmt.Println(string(body))
	return est
}

func rand_data(maxa int, mine int) string {
	valor := rand.Intn(maxa-mine) + mine
	return strconv.Itoa(valor)
}
