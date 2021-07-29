package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int `json:"status_code"`
	Response   struct {
		Requests struct {
			StatusGroup struct {
				ThreeXx struct {
					Total []int `json:"total"`
				} `json:"3xx"`
				FourXx struct {
					Total []int `json:"total"`
				} `json:"4xx"`
				FiveXx struct {
					Total []int `json:"total"`
				} `json:"5xx"`
				OneXx struct {
					Total []int `json:"total"`
				} `json:"1xx"`
				TwoXx struct {
					Total []int `json:"total"`
				} `json:"2xx"`
			} `json:"status_group"`
			Ratelimit struct {
				Challenge struct {
					Total []int `json:"total"`
				} `json:"challenge"`
				Count struct {
					Total []int `json:"total"`
				} `json:"count"`
				Simulate struct {
					Total []int `json:"total"`
				} `json:"simulate"`
				ChallengeSuccess struct {
					Total []int `json:"total"`
				} `json:"challenge_success"`
				ChallengeFailed struct {
					Total []int `json:"total"`
				} `json:"challenge_failed"`
				Block struct {
					Total []int `json:"total"`
				} `json:"block"`
				Whitelist struct {
					Total []int `json:"total"`
				} `json:"whitelist"`
			} `json:"ratelimit"`
			StatusCode struct {
				Num301 struct {
					Total []int `json:"total"`
				} `json:"301"`
				Num302 struct {
					Total []int `json:"total"`
				} `json:"302"`
				Num304 struct {
					Total []int `json:"total"`
				} `json:"304"`
				Num400 struct {
					Total []int `json:"total"`
				} `json:"400"`
				Num401 struct {
					Total []int `json:"total"`
				} `json:"401"`
				Num403 struct {
					Total []int `json:"total"`
				} `json:"403"`
				Num404 struct {
					Total []int `json:"total"`
				} `json:"404"`
				Num499 struct {
					Total []int `json:"total"`
				} `json:"499"`
				Num500 struct {
					Total []int `json:"total"`
				} `json:"500"`
				Num502 struct {
					Total []int `json:"total"`
				} `json:"502"`
				Num503 struct {
					Total []int `json:"total"`
				} `json:"503"`
				Num504 struct {
					Total []int `json:"total"`
				} `json:"504"`
			} `json:"status_code"`
			Smartrule struct {
				Challenge struct {
					Total []int `json:"total"`
				} `json:"challenge"`
				Accept struct {
					Total []int `json:"total"`
				} `json:"accept"`
				ChallengeFailed struct {
					Total []int `json:"total"`
				} `json:"challenge_failed"`
				ChallengeSuccess struct {
					Total []int `json:"total"`
				} `json:"challenge_success"`
				Block struct {
					Total []int `json:"total"`
				} `json:"block"`
			} `json:"smartrule"`
			Waf struct {
				Challenge struct {
					Total []int `json:"total"`
				} `json:"challenge"`
				ChallengeSuccess struct {
					Total []int `json:"total"`
				} `json:"challenge_success"`
				ChallengeFailed struct {
					Total []int `json:"total"`
				} `json:"challenge_failed"`
				Simulate struct {
					Total []int `json:"total"`
				} `json:"simulate"`
				Block struct {
					Total []int `json:"total"`
				} `json:"block"`
			} `json:"waf"`
			Firewall struct {
				Accept struct {
					Total []int `json:"total"`
				} `json:"accept"`
				Block struct {
					Total []int `json:"total"`
				} `json:"block"`
			} `json:"firewall"`
		} `json:"requests"`
		Stats struct {
			StatusGroup struct {
				ThreeXx int `json:"3xx"`
				FourXx  int `json:"4xx"`
				FiveXx  int `json:"5xx"`
				Others  int `json:"others"`
				OneXx   int `json:"1xx"`
				TwoXx   int `json:"2xx"`
			} `json:"status_group"`
			ReqPerSec struct {
				Total struct {
					Avg float64 `json:"avg"`
					Max float64 `json:"max"`
				} `json:"total"`
			} `json:"req_per_sec"`
			Requests struct {
				Total       int `json:"total"`
				User        int `json:"user"`
				ThreeLcloud int `json:"3lcloud"`
				Optimized   int `json:"optimized"`
			} `json:"requests"`
			Firewall struct {
				Accept int `json:"accept"`
				Block  int `json:"block"`
			} `json:"firewall"`
			Smartrule struct {
				Challenge        int `json:"challenge"`
				Accept           int `json:"accept"`
				ChallengeFailed  int `json:"challenge_failed"`
				ChallengeSuccess int `json:"challenge_success"`
				Block            int `json:"block"`
			} `json:"smartrule"`
			CacheCoverage struct {
				Uncacheable int `json:"uncacheable"`
				Cached      int `json:"cached"`
				Uncached    int `json:"uncached"`
			} `json:"cache_coverage"`
			AvgReqSize struct {
				Total float64 `json:"total"`
			} `json:"avg_req_size"`
			Ratelimit struct {
				Cdn              int `json:"cdn"`
				Count            int `json:"count"`
				Simulate         int `json:"simulate"`
				ChallengeSuccess int `json:"challenge_success"`
				ChallengeFailed  int `json:"challenge_failed"`
				Block            int `json:"block"`
				User             int `json:"user"`
				Challenge        int `json:"challenge"`
				Whitelist        int `json:"whitelist"`
			} `json:"ratelimit"`
			Bandwidth struct {
				Ratio struct {
					Optimized float64 `json:"optimized"`
				} `json:"ratio"`
				User        float64 `json:"user"`
				Optimized   float64 `json:"optimized"`
				Unoptimized float64 `json:"unoptimized"`
				ThreeLcloud float64 `json:"3lcloud"`
				Total       float64 `json:"total"`
			} `json:"bandwidth"`
			Latency struct {
				HTML    float64 `json:"html"`
				NonHTML float64 `json:"non_html"`
			} `json:"latency"`
			Waf struct {
				Challenge        int `json:"challenge"`
				ChallengeSuccess int `json:"challenge_success"`
				ChallengeFailed  int `json:"challenge_failed"`
				Simulate         int `json:"simulate"`
				Block            int `json:"block"`
			} `json:"waf"`
			Pageviews struct {
				User        int `json:"user"`
				ThreeLcloud int `json:"3lcloud"`
				Total       int `json:"total"`
			} `json:"pageviews"`
		} `json:"stats"`
		Metadata struct {
			StartPoint int     `json:"start_point"`
			EndPoint   int     `json:"end_point"`
			Interval   int     `json:"interval"`
			ServerTime float64 `json:"server_time"`
		} `json:"metadata"`
		Traffic struct {
			BandwidthPerSec struct {
				Total       []float64 `json:"total"`
				ThreeLcloud []float64 `json:"3lcloud"`
				User        []float64 `json:"user"`
			} `json:"bandwidth_per_sec"`
			ReqPerSec struct {
				Total       []float64 `json:"total"`
				ThreeLcloud []float64 `json:"3lcloud"`
				User        []float64 `json:"user"`
			} `json:"req_per_sec"`
			Requests struct {
				ThreeLcloud []int `json:"3lcloud"`
				Total       []int `json:"total"`
				Optimized   []int `json:"optimized"`
				User        []int `json:"user"`
			} `json:"requests"`
			Bandwidth struct {
				Total       []float64 `json:"total"`
				Optimized   []float64 `json:"optimized"`
				ThreeLcloud []float64 `json:"3lcloud"`
				Unoptimized []float64 `json:"unoptimized"`
				User        []float64 `json:"user"`
			} `json:"bandwidth"`
			Pageviews struct {
				Total       []int `json:"total"`
				ThreeLcloud []int `json:"3lcloud"`
				User        []int `json:"user"`
			} `json:"pageviews"`
			Latency struct {
				HTML    []float64 `json:"html"`
				NonHTML []float64 `json:"non_html"`
			} `json:"latency"`
			AvgReqSize struct {
				Total []float64 `json:"total"`
			} `json:"avg_req_size"`
		} `json:"traffic"`
	} `json:"response"`
}

func main() {
	//tNow := time.Now().Add(-24*time.Hour)
	tUnix := time.Now().Unix()
	tYesterday := tUnix - 86400
	tHour := tUnix - 3600
	tWeek := tUnix - 604800
	fmt.Printf("Now is = %+v\n", tUnix)
	fmt.Printf("Last hour = %+v\n", tHour)
	fmt.Printf("Last 24hrs = %+v\n", tYesterday)
	fmt.Printf("Last week = %+v\n", tWeek)

	var token string
	var domain string

	flag.StringVar(&token, "token", "", "Ex.: -token <token> | Token GoCache. (Required)")
	flag.StringVar(&domain, "domain", "", "Ex.: -domain google.com.br | Domain. (Required)")
	flag.Parse()

	url := fmt.Sprintf("https://api.gocache.com.br/v1/analytics/%v?graph=custom&interval=1min&from=%v&to=%v", domain, tHour, tUnix)

	fmt.Printf("URL Test = %+v\n", url)

	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("GoCache-Token", fmt.Sprintf("%s", token))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Response

	if err := json.Unmarshal(body, &responseObject); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(PrettyPrint(responseObject))
	fmt.Println("Teste:", token)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
