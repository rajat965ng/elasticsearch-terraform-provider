package es_ops

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Create(id int, name string, payload string) error {
	client := http.Client{}
	url := "http://localhost:9200/" + name + "/_doc/" + strconv.Itoa(id)
	fmt.Println("URL: ", url)
	fmt.Println("payload: ", payload)
	req, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buff := make([]byte, 0)
	n, err := resp.Body.Read(buff)
	if err != nil {
		panic(err)
	}
	fmt.Println("n: ", n)
	return nil
}
