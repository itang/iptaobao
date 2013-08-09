package iptaobao

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const RestApiUrlPrefix = "http://ip.taobao.com/service/getIpInfo.php?ip="

type IpInfo struct {
	CountryId string `json:"country_id"`
	Country   string `json:"country"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
	IspId     string `json:"isp_id"`
	Ip        string `json:"ip"`
}

type Ret struct {
	Code int `json:"code"`
}

type Success struct {
	Ret
	Data IpInfo `json:"data"`
}

type Failure struct {
	Ret
	Message string `json:"data"`
}

func GetIpInfo(ip string) (*IpInfo, error) {
	content, err := getRestCallResult(RestApiUrlPrefix + ip)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(content), `"code":1`) { // error
		ret := Failure{}
		if err := json.Unmarshal(content, &ret); err != nil {
			return nil, err
		}
		return nil, errors.New(ret.Message)
	} else { // success
		ret := Success{}
		if err := json.Unmarshal(content, &ret); err != nil {
			return nil, err
		}

		if ret.Code == 0 {
			return &ret.Data, nil
		}
		return nil, errors.New(fmt.Sprintf("unexpected return code %d", ret.Code))
	}
}

func getRestCallResult(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
