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

func GetIpInfo(ip string) (error, *IpInfo) {
	err, content := getRestCallResult(RestApiUrlPrefix + ip)
	if err != nil {
		return err, nil
	}

	if strings.Contains(string(content), `"code":1`) { // error
		ret := Failure{}
		if err := json.Unmarshal(content, &ret); err != nil {
			return err, nil
		}
		return errors.New(ret.Message), nil
	} else { // success
		ret := Success{}
		if err := json.Unmarshal(content, &ret); err != nil {
			return err, nil
		}

		if ret.Code == 0 {
			return nil, &ret.Data
		}
		return errors.New(fmt.Sprintf("unexpected return code %d", ret.Code)), nil
	}
}

func getRestCallResult(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return err1, nil
	}
	return nil, body
}
