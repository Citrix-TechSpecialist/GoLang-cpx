package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"net/http/cookiejar"
	"fmt"
	"encoding/json"
	"encoding/base64"
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
			Feature: []string{"LB", "SSL"},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func uploadCertAndKeyFiles(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/systemfile"

	type Systemfile struct {
		Filename string `json:"filename"`
		Filelocation string `json:"filelocation"`
		Filecontent string `json:"filecontent"`
		Fileencoding string `json:"fileencoding"`
	}
	type PostData struct {
		Systemfile []Systemfile `json:"systemfile"`
	}

	rscrt := base64.StdEncoding.EncodeToString([]byte(`-----BEGIN CERTIFICATE-----
MIIESzCCAzOgAwIBAgIBAjANBgkqhkiG9w0BAQsFADB5MQswCQYDVQQGEwJVUzET
MBEGA1UECBMKQ2FsaWZvcm5pYTERMA8GA1UEBxMIU2FuIEpvc2UxEzARBgNVBAoT
CkNpdHJpeCBBTkcxFDASBgNVBAsTC05TIEludGVybmFsMRcwFQYDVQQDEw5kZWZh
dWx0IEVPRUxOTDAeFw0xNjA5MjgxNjA1MjBaFw0xNzA5MjgxNjA1MjBaMD4xCzAJ
BgNVBAYTAlVTMQswCQYDVQQIEwJUWDEMMAoGA1UEChMDQ1RYMRQwEgYDVQQDEwty
c2VydmVyLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOGealU5
gNLYY0c2WWWwulsW4SHavEx8cxUegLIZU9cf2+J8Dr/fsty7Arj7AMMsilimaIQw
K14bjHuTaM5fDNR2/ScdkED4F/NwCM9ruvgzOjUAUST983PxRNi/5CIWs4GJXv+1
41xqSB9XXcHiLYAd/b+bi1YQsXygBfT50TiIiw3Oh9yo3JSwp/48eqEqGf5iN8HM
vc/8dDf6usu75wfNTsqMIKQwt54Pz5H98VOdor7y8PXMFLEmtLFtwnwR7CGJr1wV
FcAoxQHh9CeGZ6tKtfkNsnI3fAB2LIuxbiCGOqbcE9I08BE8prqMNTljGtTeVHLT
QTGLxk2qZDXBZo8CAwEAAaOCARcwggETMAkGA1UdEwQCMAAwEQYJYIZIAYb4QgEB
BAQDAgZAMC4GCWCGSAGG+EIBDQQhFh9OZXRTY2FsZXIgR2VuZXJhdGVkIENlcnRp
ZmljYXRlMB0GA1UdDgQWBBTAARu/rsZLEWmpxwaEcbJ9y+X6fTCBowYDVR0jBIGb
MIGYgBQFbFHJY8DB/HeY7d4b3/cr8kqOs6F9pHsweTELMAkGA1UEBhMCVVMxEzAR
BgNVBAgTCkNhbGlmb3JuaWExETAPBgNVBAcTCFNhbiBKb3NlMRMwEQYDVQQKEwpD
aXRyaXggQU5HMRQwEgYDVQQLEwtOUyBJbnRlcm5hbDEXMBUGA1UEAxMOZGVmYXVs
dCBFT0VMTkyCAQAwDQYJKoZIhvcNAQELBQADggEBAEWCZzLCLeIeFHj9gXkD5jMN
KuArOowPvnflRDyW9lRrnWBm9KSrCVn8PBuu8F+4/z83SmVG1cnlwPjmppo1nTD+
B3isayV3oVS2DLkVNzmIgOHgg9s3k5gXYmNXJ1E9BF2TJUAU5wZQIqiUNiZIQbxU
m2Nj/2BjLI1O9cv7bmxvHcc/D+dn3/hg4+d7370ZlAQ+n+bkjCrlumMVZUWySHcf
VyQoY7MKrdBTakgFrpvQQETUTOet1+5JDwi55MDeA3gtos8N2/rlum8GQ9Dn0sM+
Vmv37ajQYylmzB66WYSlWoKBjb/Ua3Sf4dbJPjxYUSAwhD7BMiIIUYTS9Bx8dPo=
-----END CERTIFICATE-----
`))
	rskey := base64.StdEncoding.EncodeToString([]byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4Z5qVTmA0thjRzZZZbC6WxbhIdq8THxzFR6AshlT1x/b4nwO
v9+y3LsCuPsAwyyKWKZohDArXhuMe5Nozl8M1Hb9Jx2QQPgX83AIz2u6+DM6NQBR
JP3zc/FE2L/kIhazgYle/7XjXGpIH1ddweItgB39v5uLVhCxfKAF9PnROIiLDc6H
3KjclLCn/jx6oSoZ/mI3wcy9z/x0N/q6y7vnB81OyowgpDC3ng/Pkf3xU52ivvLw
9cwUsSa0sW3CfBHsIYmvXBUVwCjFAeH0J4Znq0q1+Q2ycjd8AHYsi7FuIIY6ptwT
0jTwETymuow1OWMa1N5UctNBMYvGTapkNcFmjwIDAQABAoIBAH3ZK17WcHErmlUC
j+MVLR3aKUIFDLttP5QsK4Usc4OvlatDn8aPNOnCtsYP3GEB2zmPuQTjCY24uCfG
FdPnWPS6WoMTDn/u4w07FO7+HJCNoo4l2x1TOhUWI1zzzIDnQMGkqoTgJC5MamZx
CS84xkCMehoC3TnondfyOuBm6LkrBJwsSAQD/lPnnwdlqpNyzoIDnUKL8YMPyyGC
kvA/K/YpSsAqdEhW68+ma26hfEiL4PeoJgIc7uzTbO2cKwJQWIio9/7UGTD/KE2c
msnh5xmiaOpFyoFQMb+4U2Fxpq1Ds5Eew5kOv9FN1ev31jqGsWWLzWvIu93BDGl8
BVxtK4ECgYEA9Pn6DF2V80O93TSiIgXTqHLOrIGNCznc9Ei1cimL8yik3E7qb6tG
5+vz3i3WPD/d2qSZVZ35sVb/Tr2ULJkWUfIEjQykIRco/kSkTlxVLvsqQiyfBNYS
dmhsW0NP82ritQPuEWkA3930fcRDk/Sg9S+YKyoGmrhoZzRoDipv378CgYEA68Vy
Ww5+i12DFqZ/zhvKSWYGkMg83jUxaXXFfZy10dP49fXPiyMFSjOQbCjgLDml41yX
FMIpf2EySFY4QtedUpSTf98HdosUfvex8yxGwXHm5fEjCnKHNbgDaLrxH4K5gWah
yMONlTpiovgJQHihO9RMKbxG7KAYoLZuLsS9LTECgYEA2/Jon5uS2yPyHt53x2ZF
39KcXuO+F9su16FET6ifv4S5aBfugq5b7jS58rxiwhtxfDIWfXllyuRaO38Yv2X/
VTme/mjgH9mkc457muNpk9Hr4hgf+f6d+vPMfbAU428O9wj9QWJuZ5DnR9fj+L5F
mX3O+Mo1vcpd6nNyDW3qng8CgYAWIcrCUXH/kx+jGK4WovUyPqmPHbzY/xVMWQnY
6MUIlWVhcVmyLe9pL7326T9h52dzGFX2VOOgWXdm4vEVFThncBsIfd8teZDK+mVx
9k4OCqsqGqC3cljO6h8nzaSk2JihVQkK15CK2Zg4xB/aNXitLRiZMltWCxFExNtC
+KTpgQKBgQCHMghURj5lnPfvT3EdU0JAoumWTH1yjcxDODJCTEBgW76QiQ6aIXI3
Z3A18JayK7oM2rowg97ed6M4Z518TGGkDa4ehZq6fh2hdKR0Y3tPWxUcjiaW7QVz
7YMFYAe5g5pdcx2ypYcgEVRyDORaHuVadI1HAxw/5/xESMCSgCI8NA==
-----END RSA PRIVATE KEY-----
`))

	data := PostData{
		[]Systemfile{
			Systemfile{
				Filename: "rs.crt",
				Filelocation: "/nsconfig/ssl/",
				Filecontent: rscrt,
				Fileencoding: "BASE64",
				},
			Systemfile{
				Filename: "rs.key",
				Filelocation: "/nsconfig/ssl/",
				Filecontent: rskey,
				Fileencoding: "BASE64",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func installCrtKey(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/sslcertkey"
	
	type Sslcertkey struct {
		Certkey string `json:"certkey"`
		Cert string `json:"cert"`
		Key string `json:"key"`
	}
	type PostData struct {
		Sslcertkey Sslcertkey `json:"sslcertkey"`
	}

	data := PostData{
		Sslcertkey{
			Certkey: "rs",
			Cert: "/nsconfig/ssl/rs.crt",
			Key: "/nsconfig/ssl/rs.key",
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
				Name: "lbvs-ssl-rserver.com",
				Servicetype: "ssl",
				Ipv46: "192.168.20.15",
				Port: 443,
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
				Name: "lbvs-ssl-rserver.com",
				Servicename: "http-rserver.com",
				},
		},
	}
	out, _ := json.Marshal(data)

	postReq(client, url, string(out))
}

func bindLBVserversToCertKey(client http.Client) {
	url := "http://" + cpxip + ":" + cpxport + "/nitro/v1/config/sslvserver_sslcertkey_binding"

	type Sslvserver_sslcertkey_binding struct {
		Vservername string `json:"vservername"`
		Certkeyname string `json:"certkeyname"`
	}
	type PostData struct {
		Sslvserver_sslcertkey_binding []Sslvserver_sslcertkey_binding `json:"sslvserver_sslcertkey_binding"`
	}

	data := PostData{
		[]Sslvserver_sslcertkey_binding{
			Sslvserver_sslcertkey_binding{
				Vservername: "lbvs-ssl-rserver.com",
				Certkeyname: "rs",
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
	uploadCertAndKeyFiles(client)
	installCrtKey(client)
	addServers(client)
	addServices(client)
	addLBVservers(client)
	bindLBVserversToServices(client)
	bindLBVserversToCertKey(client)
	saveConfig(client)
	logout(client)

	fmt.Println("done")
}
