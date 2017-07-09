package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"net/http/cookiejar"
	"fmt"
	"encoding/json"
)

const cpxip string = "192.168.6.5"
const cpxport string = "80"
const cpxusername string = "nsroot"
const cpxpassword string = "nsroot"
const debug bool = true

func postReq(client http.Client, url string, data string) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Add("content-type", "application/json")

	res, _ := client.Do(req)
	defer res.Body.Close()

	if debug {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(url, string(body))
	}
}

func putReq(client http.Client, url string, data string) {
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
	req.Header.Add("content-type", "application/json")

	res, _ := client.Do(req)
	defer res.Body.Close()

	if debug {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(url, string(body))
	}
}

func login(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/login"
	
	type Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Timeout int `json:"timeout"`
	}
	type PostData struct {
		Login Login `json:"login"`
	}

	data := PostData{
		Login{
			Username: cpxusername,
			Password: cpxpassword,
			Timeout: 300,
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func enableDNSRecursion(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/dnsparameter"

	type Dnsparameter struct {
		Recursion string `json:"recursion"`
	}
	type PostData struct {
		Dnsparameter Dnsparameter `json:"dnsparameter"`
	}

	data := PostData{
		Dnsparameter{
			Recursion: "ENABLED",
		},
	}
	out, _ := json.Marshal(data)

	fmt.Println(string(out))
	putReq(client, url, string(out))
}

func addADNSServer(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/dnsnameserver"

	type Dnsnameserver struct {
		Ip string `json:"ip"`
		Type string `json:"type"`
		Local bool `json:"local"`
	}
	type PostData struct {
		Dnsnameserver []Dnsnameserver `json:"dnsnameserver"`
	}

	data := PostData{
		[]Dnsnameserver{
			Dnsnameserver{
				Ip: "192.168.20.20",
				Type: "UDP",
				Local: true,
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addARecords(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/dnsaddrec"

	type Dnsaddrec struct {
		Hostname string `json:"hostname"`
		Ipaddress string `json:"ipaddress"`
	}
	type PostData struct {
		Dnsaddrec []Dnsaddrec `json:"dnsaddrec"`
	}

	data := PostData{
		[]Dnsaddrec{
			Dnsaddrec{
				Hostname: "r.cluster.local",
				Ipaddress: "10.10.10.10",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func saveConfig(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/nsconfig?action=save"
	
	type Nsconfig struct {
	}
	type PostData struct {
		Nsconfig Nsconfig `json:"nsconfig"`
	}

	data := PostData{
		Nsconfig{
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func logout(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/logout"
	
	type Logout struct {
	}
	type PostData struct {
		Logout Logout `json:"logout"`
	}

	data := PostData{
		Logout{
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func main() {
	cookieJar, _ := cookiejar.New(nil)
	client := http.Client{
	    Jar: cookieJar,
	}

	login(client)
	enableDNSRecursion(client)
	addADNSServer(client)
	addARecords(client)
	saveConfig(client)
	logout(client)

	fmt.Println("done")
}
