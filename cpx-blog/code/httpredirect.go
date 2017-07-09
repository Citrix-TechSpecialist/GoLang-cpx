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

func enableFeature(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/nsfeature?action=enable"

	type NsFeature struct {
		Feature []string `json:"feature"`
	}
	type PostData struct {
		NsFeature NsFeature `json:"nsfeature"`
	}

	data := PostData{
		NsFeature{
			Feature: []string{"LB"},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addLBVservers(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/lbvserver"

	type Lbvserver struct {
		Name string `json:"name"`
		Servicetype string `json:"servicetype"`
		Ipv46 string `json:"ipv46"`
		Port int `json:"port"`
		Redirurl string `json:"redirurl"`
	}
	type PostData struct {
		Lbvserver []Lbvserver `json:"lbvserver"`
	}

	data := PostData{
		[]Lbvserver{
			Lbvserver{
				Name: "lbvs-redir-http-rserver.com",
				Servicetype: "http",
				Ipv46: "192.168.20.16",
				Port: 80,
				Redirurl: "https://rserver.com",
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
	enableFeature(client)
	addLBVservers(client)
	saveConfig(client)
	logout(client)

	fmt.Println("done")
}
