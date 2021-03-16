package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	VertifyEnv()

}

func VertifyEnv() {
	imageString := "registry-jinan-lab.inspurcloud.cn/library/cke/virt-operator-amd64:v0.32.0"
	imageRex :=regexp.MustCompile("^(.*)/(.*)virt-operator.*([@:].*)?$")
	matches := imageRex.FindAllStringSubmatch(imageString,2)
	//fmt.Println("matches", len(matches), len(matches[0]), matches[0][1])
	for i, v := range matches[0] {
		fmt.Println(i, " ", v)
	}
}

func test() {
	imagePrefix := "test"
	imageName := fmt.Sprintf("%s%s", imagePrefix, "virt-api")
	fmt.Println(imageName)
}

func AddVersionSeparatorPrefix(version string) string {
	// version can be a template, a tag or shasum
	// prefix tags with ":" and shasums with "@"
	// templates have to deal with the correct image/version separator themselves
	if strings.HasPrefix(version, "sha256:") {
		version = fmt.Sprintf("@%s", version)
	} else if !strings.HasPrefix(version, "{{if") {
		version = fmt.Sprintf(":%s", version)
	}
	return version
}