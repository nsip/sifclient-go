package sifclient

import (
	"github.com/nsip/sifdata-go/infrastructure/environment"
	"time"
)

// SIFClientData stores internal data, including the state of authentication
type SIFClientData struct {
	ApplicationKey       string
	Password             string
	UserToken            string
	SolutionId           string
	ConsumerName         string
	EnvironmentURL       string
	AuthenticationMethod string
	Timestamp            string
	AuthToken            string
	LastEnvironment      environment.EnvironmentResponse
	Debug                bool
}

// Create a new SIF client, which requires an environment URL
// application key and password. Returns the new object used for requets.
func New(url string, appkey string, password string) *SIFClientData {
	s := &SIFClientData{
		SolutionId:           "HITS",
		AuthenticationMethod: "Basic",
		EnvironmentURL:       url,
		ApplicationKey:       appkey,
		UserToken:            appkey,
		Password:             password,
		//  Generate timestamp in format RFC3339
		Timestamp: time.Now().Format(time.RFC3339),
	}
	return s
}

func (s *SIFClientData) GetURL() string {
	return s.EnvironmentURL
}

func (s *SIFClientData) SetAuthenticationMethod_SIF_HMACSHA256() {
	s.AuthenticationMethod = "SIF_HMACSHA256"
}

func (s *SIFClientData) SetDebug(in bool) {
	s.Debug = in
}
