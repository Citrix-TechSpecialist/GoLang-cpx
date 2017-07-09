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

func addServers(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/server"

	type Server struct {
		Name string `json:"name"`
		Domain string `json:"domain"`
		Comment string `json:"comment"`
	}
	type PostData struct {
		Server []Server `json:"server"`
	}

	data := PostData{
		[]Server{
			Server{
				Name: "qserver.com",
				Domain: "qserver.com",
				Comment: "Q Web Server",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addServiceGroup(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/servicegroup"

	type Servicegroup struct {
		Servicegroupname string `json:"servicegroupname"`
		Servicetype string `json:"servicetype"`
		Autoscale string `json:"autoscale"`
	}
	type PostData struct {
		Servicegroup Servicegroup `json:"servicegroup"`
	}

	data := PostData{
		Servicegroup{
			Servicegroupname: "svg-qserver.com",
			Servicetype: "HTTP",
			Autoscale: "DNS",
			},
	}

	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func bindServerToServiceGroup(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/servicegroup_servicegroupmember_binding"

	type Servicegroup_servicegroupmember_binding struct {
		Servicegroupname string `json:"servicegroupname"`
		Servername string `json:"servername"`
		Port int `json:"port"`
	}
	type PostData struct {
		Servicegroup_servicegroupmember_binding Servicegroup_servicegroupmember_binding `json:"servicegroup_servicegroupmember_binding"`
	}

	data := PostData{
		Servicegroup_servicegroupmember_binding{
			Servicegroupname: "svg-qserver.com",
			Servername: "qserver.com",
			Port: 80,
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
	}
	type PostData struct {
		Lbvserver []Lbvserver `json:"lbvserver"`
	}

	data := PostData{
		[]Lbvserver{
			Lbvserver{
				Name: "lbvs-http-qserver.com",
				Servicetype: "http",
				Ipv46: "192.168.20.17",
				Port: 80,
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func bindLBVserverToServiceGroup(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/lbvserver_servicegroup_binding"

	type Lbvserver_servicegroup_binding struct {
		Servicegroupname string `json:"servicegroupname"`
		Name string `json:"name"`
	}
	type PostData struct {
		Lbvserver_servicegroup_binding Lbvserver_servicegroup_binding `json:"lbvserver_servicegroup_binding"`
	}

	data := PostData{
		Lbvserver_servicegroup_binding{
			Servicegroupname: "svg-qserver.com",
			Name: "lbvs-http-qserver.com",
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
	addServers(client)
	addServiceGroup(client)
	bindServerToServiceGroup(client)
	addLBVservers(client)
	bindLBVserverToServiceGroup(client)
	saveConfig(client)
	logout(client)

	fmt.Println("done")
}
