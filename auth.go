package sifclient
import (
    "crypto/hmac"
    "crypto/sha256"
    b64 "encoding/base64"
    "log"
)

func (s *SIFClientData) generateTokenBasic(user string) {
    s.AuthToken = b64.StdEncoding.EncodeToString([]byte(user  + ":" + s.Password))
}

func (s *SIFClientData) generateTokenHMAC256(user string) {

    // Make the hmac builder
    h := hmac.New(sha256.New, []byte(s.Password))

    // Clear text to send to HMAC
    proofText := user + ":" + s.Timestamp
    h.Write([]byte(proofText))

    // Base64 Encoded hmac out
    proofEnc := b64.StdEncoding.EncodeToString(h.Sum(nil))

    // Base64 the token
    s.AuthToken = b64.StdEncoding.EncodeToString([]byte(user + ":" + proofEnc))
}

func (s *SIFClientData) GetAuthToken() string {
    if s.AuthenticationMethod == "Basic" {
        s.generateTokenBasic(s.ApplicationKey)
        return "Basic " + s.AuthToken
    } else if s.AuthenticationMethod == "SIF_HMACSHA256" {
        s.generateTokenHMAC256(s.ApplicationKey)
        return "SIF_HMACSHA256 " + s.AuthToken
    } else {
        log.Fatal("Failed GetAuthorization, method must be Basic or SIF_HMACSHA256")
        return ""
    }
}

func (s *SIFClientData) GetAuthorization() string {
    if s.AuthenticationMethod == "Basic" {
        s.generateTokenBasic(s.LastEnvironment.SessionToken)
        return "Basic " + s.AuthToken
    } else if s.AuthenticationMethod == "SIF_HMACSHA256" {
        s.generateTokenHMAC256(s.LastEnvironment.SessionToken)
        return "SIF_HMACSHA256 " + s.AuthToken
    } else {
        log.Fatal("Failed GetAuthorization, method must be Basic or SIF_HMACSHA256")
        return ""
    }

}
