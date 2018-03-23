package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	base     string = "http://100.100.100.200"
	ram      string = "latest/meta-data/ram"
	security string = "security-credentials"
)

//the AccessKeyId --> --access-key-id
//the AccessKeySecret --> --access-key-secret
//the SecurityToken --> --sts-token
type RAMToken struct {
	AccessKeyID     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	StsToken        string `json:"SecurityToken"`
}

// check connective with the 100.100.100.200
func check(url string) (string, bool) {
	//fmt.Println(url)
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)

	if err != nil {
		return err.Error(), false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(body), true
}

// recursive to the role name

func checkRAM() bool {
	urlRAM := fmt.Sprintf("%s/%s/", base, ram)

	_, err := check(urlRAM)

	if !err {
		return false
	}
	return true
}

func checkSecurity() (string, bool) {
	urlSecurity := fmt.Sprintf("%s/%s/%s/", base, ram, security)
	resp, err := check(urlSecurity)
	if !err {
		return resp, false
	}

	return resp, true
}

func checkToken(roleName string) (string, bool) {
	token := fmt.Sprintf("%s/%s/%s/%s", base, ram, security, roleName)
	resp, err := check(token)
	if err {
		return resp, false
	}
	fmt.Println(resp)
	return resp, true
}

//GetCredential return the credential of the instance
func GetCredential() (dat RAMToken) {

	ram := checkRAM()
	if ram {
		res, err := checkSecurity()
		//fmt.Println(res)
		if !err {
			return
		}
		token, _ := checkToken(res)
		//fmt.Println(token)

		if err := json.Unmarshal([]byte(token), &dat); err != nil {
			panic(err)
		}
		//fmt.Println(dat.AccessKeyId)
		//fmt.Println(dat.AccessKeySecret)
		//fmt.Println(dat.SecurityToken)
		return

	}
	return
}
