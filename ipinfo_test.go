package iptaobao

import (
	"fmt"
	. "github.com/itang/gotang/test"
	"testing"
)

func TestGetIpInfo(t *testing.T) {
	err, ipInfo := GetIpInfo("8.8.8.8")
	fmt.Printf("%v %s\n", ipInfo, ipInfo.Country)
	AssertTrue(t, err == nil)
	AssertTrue(t, ipInfo.CountryId == "US")

	err1, _ := GetIpInfo("444.44")
	fmt.Printf("%v \n", err1)
	AssertTrue(t, err1 != nil)
}
