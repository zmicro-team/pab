package util

import (
	"fmt"
	"testing"
)

func TestMap2Xml(t *testing.T) {
	m := make(map[string]string)
	m["Name"] = "test"
	m["Id"] = "1"

	xmlStr, _ := Map2Xml(m)
	fmt.Println(xmlStr)

	m2, _ := Xml2Map(xmlStr)
	fmt.Println(m2)
}
