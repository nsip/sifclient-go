package sifclient

import (
    "log"
    "fmt"
)

func (s *SIFClientData) Get(what string) []byte {
    restclient := s.GetRequest()

    url := s.GetRequestURL() + "/" + what

    resp, err := restclient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

    // Explore response object
    fmt.Println("Response Info:")
    fmt.Println("Error      :", err)
    fmt.Println("Status Code:", resp.StatusCode())
    fmt.Println("Status     :", resp.Status())
    fmt.Println("Time       :", resp.Time())
    fmt.Println("Received At:", resp.ReceivedAt())
    // fmt.Println("Body       :\n", resp)
    // fmt.Println()
    // XXX check status and exception

    // s.LastEnvironment = environment.ParseEnvironmentResponse(resp.Body())
    // fmt.Println("SessionToken", s.LastEnvironment.SessionToken)
    // fmt.Println(resp.Header())

    fmt.Println("Navigation page/size:", resp.Header().Get("Navigationpage"), resp.Header().Get("Navigationpagesize"))
    return resp.Body()
}
