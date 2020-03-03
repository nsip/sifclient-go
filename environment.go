package sifclient
import (
    "github.com/nsip/sifdata-go/infrastructure/environment"
    "log"
)

func (s *SIFClientData) GetEnviromentRequest() environment.EnvironmentRequest {
    e:= environment.NewEnvironmentRequest()
    e.UserToken = s.UserToken
    e.ApplicationInfo.ApplicationKey = s.ApplicationKey
    e.UserToken = s.ApplicationKey  // NOTE: Duplicating for now
    return e
}

func (s *SIFClientData) CreateEnviroment() {
    restclient := s.GetClient()

    resp, err := restclient.R().
    		EnableTrace().
            SetHeader("Content-Type", "application/xml; charset=UTF-8").
            SetBody(s.GetEnviromentRequest()).
            // SetBasicAuth(s.ApplicationKey, s.Password).
            SetHeader("Authorization", s.GetAuthToken()).
            SetHeader("Timestamp", s.Timestamp).
            SetHeader("applicationKey", s.ApplicationKey).
    		Post(s.EnvironmentURL)

	if err != nil {
		log.Fatal(err)
	}

    // Explore response object
    // fmt.Println("Response Info:")
    // fmt.Println("Error      :", err)
    // fmt.Println("Status Code:", resp.StatusCode())
    // fmt.Println("Status     :", resp.Status())
    // fmt.Println("Time       :", resp.Time())
    // fmt.Println("Received At:", resp.ReceivedAt())
    // fmt.Println("Body       :\n", resp)
    // fmt.Println()
    // XXX check status and exception

    s.LastEnvironment = environment.ParseEnvironmentResponse(resp.Body())
    // fmt.Println("SessionToken", s.LastEnvironment.SessionToken)
}

func (s *SIFClientData) GetSessionToken() string {
    // TODO - consider auto CreateEnvironment here if blank
    return s.LastEnvironment.SessionToken
}
