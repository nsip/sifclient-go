package sifclient
import (
    "github.com/nsip/sifdata-go/infrastructure/environment"
    "log"
    "fmt"
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
    if s.Debug {
        fmt.Println("Enviroment URL", s.EnvironmentURL)
    }
    resp, err := restclient.R().
    		// EnableTrace().
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

    s.LastEnvironment = environment.ParseEnvironmentResponse(resp.Body())
}

func (s *SIFClientData) GetSessionToken() string {
    // TODO - consider auto CreateEnvironment here if blank
    return s.LastEnvironment.SessionToken
}

func (s *SIFClientData) GetRequestURL() string {
    for i := range s.LastEnvironment.InfrastructureServices.InfrastructureService {
        if s.LastEnvironment.InfrastructureServices.InfrastructureService[i].Name == "requestsConnector" {
            return s.LastEnvironment.InfrastructureServices.InfrastructureService[i].Text
        }
    }
    return ""
}
