package main

import (
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {

    //creating the proxyURL
    proxyStr := "https://223.171.51.167:443"
    proxyURL, err := url.Parse(proxyStr)
    if err != nil {
        log.Println(err)
    }

    //creating the URL to be loaded through the proxy
    urlStr := "https://223.171.51.167:443"
    url, err := url.Parse(urlStr)
    if err != nil {
        log.Println(err)
    }

    //adding the proxy settings to the Transport object
    transport := &http.Transport{
        Proxy: http.ProxyURL(proxyURL),
    }

    //adding the Transport object to the http Client
    client := &http.Client{
        Transport: transport,
    }

    //generating the HTTP GET request
    request, err := http.NewRequest("GET", url.String(), nil)
    if err != nil {
        log.Println(err)
    }

    //adding proxy authentication
    auth := "user:Eclypse1234"
    basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
    request.Header.Add("Proxy-Authorization", basicAuth)

    //printing the request to the console
    dump, _ := httputil.DumpRequest(request, false)
    fmt.Println(string(dump))

    //calling the URL
    response, err := client.Do(request)
    if err != nil {
        log.Println(err)
    }

    log.Println(response.StatusCode)
    log.Println(response.Status)
    //getting the response
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println(err)
    }
    //printing the response
    log.Println(string(data))
}