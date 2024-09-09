package main

import (
	"fmt"
	"net/http"
	"io"
)

// Tipo Ipinfo
type Ipinfo struct {
	ip string 

}

/*
Retorna informações sobre o IP passado.
 - Caso nenhum IP tenha sido informado, o endereço IP a ser consultado será de onde partiu a requisição.
*/
func (i Ipinfo) Getall() (string) {
	if i.ip != "" {
		host := fmt.Sprintf("https://ipinfo.io/%s", i.ip)
		res, _ := http.Get(host)
		info, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		return string(info)
	}
	return ""
}