package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

// LinkClient is basic struct
type LinkClient struct {
	BaseURL  string
	User     string
	Password string
}

// HTTPRequest is basic stuct
type HTTPRequest struct {
	ID     string `json:"id"`
	Method string `json:"method"`
	Params []any  `json:"params"`
}

func turnInterface(ina *[]any) string {
	message := "["
	for _, value := range *ina {
		if reflect.TypeOf(value) == reflect.TypeOf(1) {
			if message == "[" {
				message = message + strconv.Itoa(value.(int))
			} else {
				message = message + "," + strconv.Itoa(value.(int))
			}
		} else if value == "true" {
			if message == "[" {
				message = message + value.(string)
			} else {
				message = message + "," + value.(string)
			}
		} else if reflect.TypeOf(value) == reflect.TypeOf(make(map[string]string)) {
			bakStr := "{"
			for k, v := range value.(map[string]string) {
				bakStr += `"` + k + `": "` + v + `",`
			}
			bakStr = bakStr[:len(bakStr)-1]
			bakStr += "}"
			if message == "[" {

				message = message + bakStr
			} else {
				message = message + "," + bakStr
			}
		} else {
			if message == "[" {
				message = message + `"` + value.(string) + `"`
			} else {
				message = message + `,"` + value.(string) + `"`
			}
		}

	}
	message += "]"
	return message
}

// SafeLinkHTTPFunc is http request tool
func (client *LinkClient) SafeLinkHTTPFunc(function string, params *[]any) *simplejson.Json {
	sleepInterval := []int{5, 10, 20, 30, 40, 60, 120, 240}
	index := 0
	for {
		returnValue := client.LinkHTTPFunc(function, params)
		if returnValue != nil {
			_, exist := returnValue.CheckGet("result")
			if exist {
				return returnValue
			}
		}
		{

			if index >= 7 {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[7]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[7]))
				fmt.Println("unkown http not reached,restart")
				break
			} else {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[index]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[index]))
			}
		}
		index++
	}
	return nil
}

// UnSafeLinkHTTPFunc is UnSafe Link HTTP function
func (client *LinkClient) UnSafeLinkHTTPFunc(function string, params *[]any) *simplejson.Json {
	sleepInterval := []int{5, 10, 20, 30, 40, 60, 120, 240, 480, 960, 1920, 3840}
	index := 0
	for {
		returnValue := client.LinkHTTPFunc(function, params)
		if returnValue != nil {
			return returnValue
		}
		{

			if index >= 7 {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[7]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[7]))
				break
			} else {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[index]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[index]))
			}
		}
		index++
	}
	return nil
}

// LinkHTTPFunc is LinkHTTPFunc function
func (client *LinkClient) LinkHTTPFunc(function string, params *[]any) *simplejson.Json {
	strParams := turnInterface(params)
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	tmpClient := http.Client{Transport: &transport}

	message := "{ \"id\": 1, \"method\": \"" + function + "\", \"params\": " + strParams + "}"
	payload := strings.NewReader(message)
	//fmt.Println(payload)
	req, err := http.NewRequest("POST", client.BaseURL, payload)
	//fmt.Println(client.IP+":"+client.Port)
	if err != nil {
		slog.Error("http request fail", "err", err)
		return nil
	}
	req.Header.Add("content-type", "application/json")
	if client.User != "" && client.Password != "" {
		encodeUser := base64.StdEncoding.EncodeToString([]byte((*client).User + ":" + (*client).Password))
		req.Header.Add("authorization", "Basic "+encodeUser)
	}
	res, err := tmpClient.Do(req)
	if err != nil {
		slog.Error("http request fail", "err", err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		slog.Error("http status code fail", "code", res.StatusCode)
		return nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("ReadAll fail", "err", err)
		return nil
	}
	//fmt.Println(function,params,string(body))
	js, err := simplejson.NewJson(body)
	if err != nil {
		slog.Error("NewJson fail", "err", err)
		return nil
	}
	return js
}

// HTTPRpcFunction is basic http request function
func (client *LinkClient) HTTPRpcFunction(function string, param *[]any) string {
	a := HTTPRequest{"1", function, *param}
	fmt.Println(a)
	b, err := json.Marshal(a)
	if err != nil {
		slog.Error("Marshal fail", "err", err)
		return ""
	}
	payload := strings.NewReader(string(b))
	encodeUser := base64.StdEncoding.EncodeToString([]byte(client.User + ":" + client.Password))
	req, _ := http.NewRequest("POST", client.BaseURL, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Basic "+encodeUser)
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("http request fail", "err", err)
		return ""
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("ReadAll fail", "err", err)
		return ""
	}
	tempParam := make(map[string]any)
	err = json.Unmarshal(body, &tempParam)
	if err != nil {
		slog.Error("Unmarshal json fail")
		return ""
	}
	slog.Debug(string(body))
	return string(body)
}
