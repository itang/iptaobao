package iptaobao

import (
	"fmt"
	"testing"

	. "github.com/itang/gotang/test"
)

func TestGetIpInfo(t *testing.T) {
	ip1 := "8.8.8.8"
	ipInfo, err := GetIpInfo(ip1)
	fmt.Printf("%s: %v %s\n", ip1, ipInfo, ipInfo.Country)
	AssertTrue(t, err == nil)
	AssertTrue(t, ipInfo.CountryId == "US")

	ip2 := "444.44"
	_, err1 := GetIpInfo(ip2)
	fmt.Printf("%s: %v \n", ip2, err1)
	AssertTrue(t, err1 != nil)
}
