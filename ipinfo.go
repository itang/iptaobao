package iptaobao

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Code int    `json:"code"`
	Data IpInfo `json:"data"`
}

func GetIpInfo(ip string) (error, *IpInfo) {
	err, content := getCallResult(RestApiUrlPrefix + ip)
	if err != nil {
		return err, nil
	}
	ret := Ret{}
	err1 := json.Unmarshal(content, &ret)
	if err1 != nil {
		return err1, nil
	}

	if ret.Code == 0 {
		return nil, &ret.Data
	}
	return errors.New(fmt.Sprintf("return code %d", ret.Code)), nil
}

func getCallResult(url string) (error, []byte) {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			return nil, body
		}
	}
	return errors.New("error"), nil
}
