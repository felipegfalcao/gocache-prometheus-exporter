package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const Namespace = "gocache"

var (
	url   string
	token string
)

func main() {
	urlBuild()
	handleRequests()
}

func urlBuild() {
	tUnix := time.Now().Unix()
	tHour := tUnix - 3600

	flag.StringVar(&token, "token", "", "Ex.: -token <token> | Token GoCache. (Required)")

	var (
		domain   = flag.String("domain", "", "Ex.: -domain google.com.br | Domain. (Required)")
		interval = flag.String("interval", "1min", "Ex.: -interval 1min")
		host     = flag.String("host", "", "Ex.: -host www.yourdomain.com.br | Permite filtro por subdomínio. O valor deve ser o subdomínio completo")
	)

	flag.Parse()

	if *host != "" {
		*host = fmt.Sprintf("&host=%+v", *host)
	}

	if *domain == "" {
		os.Exit(1)
	}

	token = fmt.Sprintf("%v", token)

	url = fmt.Sprintf("https://api.gocache.com.br/v1/analytics/%v?graph=custom&interval=%v&from=%v&to=%v%v", *domain, *interval, tHour, tUnix, *host)

}

func handleRequests() {
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func metrics(w http.ResponseWriter, r *http.Request) {
	statuscode := metricsJson{}
	err := json.Unmarshal([]byte(connector()), &statuscode)
	if err != nil {
		fmt.Println(err)
		return
	}
	v := "gocache_metric_response"
	fmt.Fprintf(w, "# HELP gocache_metric_statuscode Number of scrapes by HTTP status code.")
	fmt.Fprintf(w, "\n# TYPE gocache_metric_statuscode counter")
	fmt.Fprintf(w, "\nrequests_status_gocache_metric_statuscode: %+v", statuscode.StatusCode)
	fmt.Fprintf(w, "\n# TYPE gocache_metric_status_group 1xx, 2xx, 3xx, 4xx, 5xx total counter")
	fmt.Fprintf(w, "\n%+v_requests_status_group_1xx_total{status_group=\"1xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusGroup.OneXx.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_group_2xx_total{status_group=\"2xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusGroup.TwoXx.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_group_3xx_total{status_group=\"3xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusGroup.ThreeXx.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_group_4xx_total{status_group=\"4xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusGroup.FourXx.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_group_5xx_total{status_group=\"5xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusGroup.FiveXx.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_301_total{status_code=\"301\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num301.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_302_total{status_code=\"302\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num302.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_304_total{status_code=\"304\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num304.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_400_total{status_code=\"400\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num400.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_401_total{status_code=\"401\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num401.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_403_total{status_code=\"403\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num403.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_404_total{status_code=\"404\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num404.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_500_total{status_code=\"500\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num500.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_502_total{status_code=\"502\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num502.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_503_total{status_code=\"503\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num503.Total[0])
	fmt.Fprintf(w, "\n%+v_requests_status_code_504_total{status_code=\"504\",namespace=\"gocache\"} %+v", v, statuscode.Response.Requests.StatusCode.Num504.Total[0])
	fmt.Fprintf(w, "\n%+v_stats_status_group_1xx{status_group=\"1xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.StatusGroup.OneXx)
	fmt.Fprintf(w, "\n%+v_stats_status_group_2xx{status_group=\"2xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.StatusGroup.TwoXx)
	fmt.Fprintf(w, "\n%+v_stats_status_group_3xx{status_group=\"3xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.StatusGroup.ThreeXx)
	fmt.Fprintf(w, "\n%+v_stats_status_group_4xx{status_group=\"4xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.StatusGroup.FourXx)
	fmt.Fprintf(w, "\n%+v_stats_status_group_5xx{status_group=\"5xx\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.StatusGroup.FiveXx)
	fmt.Fprintf(w, "\n%+v_stats_status_group_others{status_group=\"others\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.StatusGroup.Others)
	fmt.Fprintf(w, "\n%+v_stats_reqpersec{reqpersec=\"avg\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.ReqPerSec.Total.Avg)
	fmt.Fprintf(w, "\n%+v_stats_reqpersec{status_group=\"max\",namespace=\"gocache\"} %+v", v, statuscode.Response.Stats.ReqPerSec.Total.Max)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/metrics")
	fmt.Println("Endpoint Hit: homepage")
}

func connector() string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

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

func registerSignals() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("Received SIGTERM, exiting...")
		os.Exit(1)
	}()
}
