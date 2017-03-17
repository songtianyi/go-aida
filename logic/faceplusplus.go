package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const (
	apiKey    = "7qKNKrhL3wPTjaL5frh4CjYgZ0DjtH1q"
	apiSecret = "DJmEVbYsEX-vrgyn_xAKJ9yxTxelFsBV"
	apiUrl    = "https://api-cn.faceplusplus.com/facepp/v3/"
)

func Detect(filename string, content []byte) ([]byte, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, _ := w.CreateFormFile("image_file", filename)
	if _, err := io.Copy(fw, bytes.NewReader(content)); err != nil {
		return nil, err
	}
	fw, _ = w.CreateFormField("api_key")
	_, _ = fw.Write([]byte(apiKey))
	fw, _ = w.CreateFormField("api_secret")
	_, _ = fw.Write([]byte(apiSecret))
	fw, _ = w.CreateFormField("return_landmark")
	_, _ = fw.Write([]byte("0"))
	fw, _ = w.CreateFormField("return_attributes")
	_, _ = fw.Write([]byte("gender,age"))

	w.Close()
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl+"detect", &b)
	req.Header.Add("Content-Type", w.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
