package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func HelpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error2!: ", err)
	}
	fmt.Println(string(body))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}


