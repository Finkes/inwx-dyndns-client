package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	username := os.Getenv("username")
	password := os.Getenv("password")
	for {
		ip, err := getIPv4Address()
		if err != nil {
			continue
		}
		updateInwxDynDns(ip, username, password)
		time.Sleep(30 * time.Second)
	}

}

type IPv4Response struct {
	Ip string `json:"ip"`
}

func getIPv4Address() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	response := IPv4Response{}
	json.Unmarshal(body, &response)

	fmt.Println("router ip address: ", response.Ip)
	return response.Ip, nil
}

func updateInwxDynDns(ip string, username string, password string) error {
	url := "https://dyndns.inwx.com/nic/update?myip=" + ip
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("updating dyndns:", resp.StatusCode)
	return nil
}

