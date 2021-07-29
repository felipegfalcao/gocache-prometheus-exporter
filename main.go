package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var url string
var token string



func metrics(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "%+v", connector())
	fmt.Println("Endpoint Hit: /metrics")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "/metrics")
	fmt.Println("Endpoint Hit: homepage")
}


func handleRequests() {
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


func urlOut() {
	var domain string
	var interval string
	var host string

	//tNow := time.Now().Add(-24*time.Hour)
	//handleRequests()
	tUnix := time.Now().Unix()
	//tYesterday := tUnix - 86400
	tHour := tUnix - 3600
	//tWeek := tUnix - 604800

	flag.StringVar(&token, "token", "", "Ex.: -token <token> | Token GoCache. (Required)")
	flag.StringVar(&domain, "domain", "", "Ex.: -domain google.com.br | Domain. (Required)")
	flag.StringVar(&interval, "interval", "1h", "Ex.: -interval 1h | Valores permitidos: 1min, 1h, 4h, 12h, 1d. (Required)")
	flag.StringVar(&host, "host", "", "Ex.: -host www.seudominio.com.br | Permite filtro por subdomínio. O valor deve ser o subdomínio completo")

	flag.Parse()

	if host != "" {
		host = fmt.Sprintf("&host=%+v", host)
	}

	token = fmt.Sprintf("%v", token)

	url = fmt.Sprintf("https://api.gocache.com.br/v1/analytics/%v?graph=custom&interval=%v&from=%v&to=%v%v", domain, interval, tHour, tUnix, host)

}

func main() {
	urlOut()
	handleRequests()
}

func connector() string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Add("GoCache-Token", fmt.Sprintf("%s", token))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	output := fmt.Sprint(string(body))
	return output
}