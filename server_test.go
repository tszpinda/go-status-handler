package main

import(  
    "net/http"
    "net/http/httptest"
    "testing"
    "io/ioutil"
    "log"
    "strings"
)

func TestOk(t *testing.T) {
    verify(request("application/json", "200"), TestResult{"[{'code':'HTTP 200','message':'ok'}]", 200}, t)
    verify(request("application/json", "404"), TestResult{"[{'code':'HTTP 404','message':'No service found for this url'}]", 200}, t)
    verify(request("application", "404"), TestResult{"Unsupported content type", 200}, t)
}

func verify(actual, expected TestResult, t *testing.T) {
    if actual.httpCode != expected.httpCode {
        t.Errorf("Expectecd httpcode: '%v' but was '%v'", expected.httpCode, actual.httpCode)
    }
    actualBody := strings.Replace(actual.body, "\"", "'", -1)
    actualBody = strings.TrimSpace(actualBody)
    if actualBody != expected.body{
        t.Errorf("Expectecd body: '%v' but was '%v'", expected.body, actualBody)
    }
}

type TestResult struct {
    body string
    httpCode int
}

func request(contentType, statusCode string) TestResult {
    ts := httptest.NewServer(http.HandlerFunc(handleAll))
    defer ts.Close()

    r, _ := http.NewRequest("GET", ts.URL, nil)
    r.Header.Set("Content-Type", contentType)
    r.Header.Set("x-status-code", statusCode)

    client := &http.Client{}
    res, err := client.Do(r)

    if err != nil {
        log.Fatal(err)
    }

    resBody, err := ioutil.ReadAll(res.Body)
    res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }
    result := TestResult{body: string(resBody), httpCode: 200}
    return result
}
