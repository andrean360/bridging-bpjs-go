package Vclaim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Respon_MentahDTO struct {
	MetaData struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response string `json:"response"`
}

type Respon_DTO struct {
	MetaData struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response interface{} `json:"response"`
}

func GetRequest(endpoint string) interface{} {

	cons_id, User_key, tstamp, X_signature := SetHeader()

	req, _ := http.NewRequest("GET", endpoint, nil)

	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")
	req.Header.Add("X-cons-id", cons_id)
	req.Header.Add("X-timestamp", tstamp)
	req.Header.Add("X-signature", X_signature)
	req.Header.Add("user_key", User_key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var resp_mentah Respon_MentahDTO
	var resp Respon_DTO
	json.Unmarshal([]byte(body), &resp_mentah)

	resp_decrypt, _ := ResponseVclaim(string(resp_mentah.Response), string(cons_id+Secret_key+tstamp))
	resp.MetaData = resp_mentah.MetaData
	json.Unmarshal([]byte(resp_decrypt), &resp.Response)

	return &resp
}

func PostRequest(endpoint string, data interface{}) interface{} {

	cons_id, User_key, tstamp, X_signature := SetHeader()

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(data)

	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", endpoint, &buf)

	req.Header.Add("Content-Type", "Application/x-www-form-urlencoded")
	req.Header.Add("X-cons-id", cons_id)
	req.Header.Add("X-timestamp", tstamp)
	req.Header.Add("X-signature", X_signature)
	req.Header.Add("user_key", User_key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var resp_mentah Respon_MentahDTO
	var resp Respon_DTO
	json.Unmarshal([]byte(body), &resp_mentah)

	resp_decrypt, _ := ResponseVclaim(string(resp_mentah.Response), string(cons_id+Secret_key+tstamp))
	resp.MetaData = resp_mentah.MetaData
	json.Unmarshal([]byte(resp_decrypt), &resp.Response)

	return &resp

}
