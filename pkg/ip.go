package ipinfo

import (
	"fmt"
	"net/http"
	"io"
)



/*
IPINFO
» Aceita dois tipos de dados
 - ip   string | opcional
 - token string | opcional
*/
type Ipinfo struct {
	ip string 
	token string
}

/*
Retorna todas a informações disponiveis sobre o IP.
 - Caso nenhum IP tenha sido informado, o endereço IP a ser consultado será de onde partiu a requisição.
 - Suporta requisição com token.
*/
func (i Ipinfo) Getall() (string) {
	if i.token != "" {
		if i.ip != "" {
			var host = fmt.Sprintf("https://ipinfo.io/%s?token=%s", i.ip, i.token)
			res, _ := http.Get(host)
			ip, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(ip)
		} else {
			var host = fmt.Sprintf("https://ipinfo.io/?token=%s", i.token)
			res, _ := http.Get(host)
			ip, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(ip)
		}

	} else {
		if i.ip != "" {
			var host = fmt.Sprintf("https://ipinfo.io/?token=%s", i.token)
			res, _ := http.Get(host)
			ip, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(ip)

		} else {
			res, _ := http.Get("https://ipinfo.io")
			ip, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(ip)

		}
	}
}


func (i Ipinfo) City() string {
	if i.token != "" {
		if i.ip != "" {
			domain := fmt.Sprintf("https://ipinfo.io/%s/city?token=%s", i.ip, i.token)
			res, _ := http.Get(domain)
			city, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(city)

		}
		domain := fmt.Sprintf("https://ipinfo.io/city?token=%s", i.token)
		res, _ := http.Get(domain)
		city, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		return string(city)

	} else {
		if i.ip != "" {
			domain := fmt.Sprintf("https://ipinfo.io/%s/city", i.ip)
			res, _ := http.Get(domain)
			city, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(city)
		} else {
			res, _ := http.Get("https://ipinfo.io/city")
			city, _ := io.ReadAll(res.Body)
			defer res.Body.Close()
			return string(city)
		}
	}
}