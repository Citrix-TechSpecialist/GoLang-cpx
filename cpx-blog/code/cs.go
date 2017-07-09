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
			Feature: []string{"LB", "CS"},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addServers(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/server"

	type Server struct {
		Name string `json:"name"`
		Ipaddress string `json:"ipaddress"`
		Comment string `json:"comment"`
	}
	type PostData struct {
		Server []Server `json:"server"`
	}

	data := PostData{
		[]Server{
			Server{
				Name: "rserver.com",
				Ipaddress: "10.21.0.95",
				Comment: "R Web Server",
				},
			Server{
				Name: "gserver.com",
				Ipaddress: "10.21.0.150",
				Comment: "G Web Server",
				},
			Server{
				Name: "bserver.com",
				Ipaddress: "10.21.0.31",
				Comment: "B Web Server",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addServices(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/service"

	type Service struct {
		Name string `json:"name"`
		Servername string `json:"servername"`
		Servicetype string `json:"servicetype"`
		Port int `json:"port"`
	}
	type PostData struct {
		Service []Service `json:"service"`
	}

	data := PostData{
		[]Service{
			Service{
				Name: "http-rserver.com",
				Servername: "rserver.com",
				Servicetype: "http",
				Port: 80,
				},
			Service{
				Name: "http-gserver.com",
				Servername: "gserver.com",
				Servicetype: "http",
				Port: 80,
				},
			Service{
				Name: "http-bserver.com",
				Servername: "bserver.com",
				Servicetype: "http",
				Port: 80,
				},
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
	}
	type PostData struct {
		Lbvserver []Lbvserver `json:"lbvserver"`
	}

	data := PostData{
		[]Lbvserver{
			Lbvserver{
				Name: "lbvs-http-rserver.com",
				Servicetype: "http",
				},
			Lbvserver{
				Name: "lbvs-http-gserver.com",
				Servicetype: "http",
				},
			Lbvserver{
				Name: "lbvs-http-bserver.com",
				Servicetype: "http",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func bindLBVserversToServices(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/lbvserver_service_binding"

	type Lbvserver_service_binding struct {
		Name string `json:"name"`
		Servicename string `json:"servicename"`
	}
	type PostData struct {
		Lbvserver_service_binding []Lbvserver_service_binding `json:"lbvserver_service_binding"`
	}

	data := PostData{
		[]Lbvserver_service_binding{
			Lbvserver_service_binding{
				Name: "lbvs-http-rserver.com",
				Servicename: "http-rserver.com",
				},
			Lbvserver_service_binding{
				Name: "lbvs-http-gserver.com",
				Servicename: "http-gserver.com",
				},
			Lbvserver_service_binding{
				Name: "lbvs-http-bserver.com",
				Servicename: "http-bserver.com",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addCSVserver(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/csvserver"

	type Csvserver struct {
		Name string `json:"name"`
		Ipv46 string `json:"ipv46"`
		Servicetype string `json:"servicetype"`
		Port int `json:"port"`
	}
	type PostData struct {
		Csvserver Csvserver `json:"csvserver"`
	}

	data := PostData{
		Csvserver{
			Name: "csvs_xserver.com",
			Ipv46: "192.168.20.13",
			Servicetype: "http",
			Port: 80,
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func addCSPolicy(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/cspolicy"

	type Cspolicy struct {
		Policyname string `json:"policyname"`
		Domain string `json:"domain"`
	}
	type PostData struct {
		Cspolicy []Cspolicy `json:"cspolicy"`
	}

	data := PostData{
		[]Cspolicy{
			Cspolicy{
				Policyname: "rserver.com",
				Domain: "rserver.com",
				},
			Cspolicy{
				Policyname: "gserver.com",
				Domain: "gserver.com",
				},
			Cspolicy{
				Policyname: "bserver.com",
				Domain: "bserver.com",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func bindCSPolicyToCSVservers(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/csvserver_cspolicy_binding"

	type Csvserver_cspolicy_binding struct {
		Name string `json:"name"`
		Policyname string `json:"policyname"`
		Targetlbvserver string `json:"targetlbvserver"`
	}
	type PostData struct {
		Csvserver_cspolicy_binding []Csvserver_cspolicy_binding `json:"csvserver_cspolicy_binding"`
	}

	data := PostData{
		[]Csvserver_cspolicy_binding{
			Csvserver_cspolicy_binding{
				Name: "csvs_xserver.com",
				Policyname: "rserver.com",
				Targetlbvserver: "lbvs-http-rserver.com",
				},
			Csvserver_cspolicy_binding{
				Name: "csvs_xserver.com",
				Policyname: "gserver.com",
				Targetlbvserver: "lbvs-http-gserver.com",
				},
			Csvserver_cspolicy_binding{
				Name: "csvs_xserver.com",
				Policyname: "bserver.com",
				Targetlbvserver: "lbvs-http-bserver.com",
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
	addServers(client)
	addServices(client)
	addLBVservers(client)
	bindLBVserversToServices(client)
	addCSVserver(client)
	addCSPolicy(client)
	bindCSPolicyToCSVservers(client)
	saveConfig(client)
	logout(client)

	fmt.Println("done")
}
