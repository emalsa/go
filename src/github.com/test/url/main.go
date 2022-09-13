package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//resp, err := http.Get("https://pinning-test.badssl.com/")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	////We Read the response body on the line below.
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	////Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)
	cmd := exec.Command("/usr/bin/curl", "--location", "--include", "--request", "GET", "--compressed", "https://stackoverflow.com/", "-H \"Accept-Encoding: gzip, deflate\"", "-H \"Accept: */*\"")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(err.Error()))
		fmt.Println("error")
		fmt.Printf("combined out:\n%s\n", string(stdout))
		return
	}
	fmt.Println("success")

	fmt.Printf("combined out:\n%s\n", string(stdout))
}
