package main

// Use api and Test
import (
    "fmt"
    "github.com/nsip/sifclient-go"
    // "encoding/xml"
    "net/http"

"github.com/labstack/echo"
"github.com/labstack/echo/middleware"
)

func main() {
    debug := true

    fmt.Println("Starting HITS Proxy")
    sifclient.Info()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

                e.Use(middleware.Gzip())
                e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
                        AllowOrigins: []string{"*"},
                        AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
                        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
                })) // allow cors requests during testing



    e.GET("/hits/:id/:provider", func(c echo.Context) error {
        id := c.Param("id")
        provider := c.Param("provider")
        // minimum = url environment, applicationKey, password/secret
        client := sifclient.New(
            "https://hits.nsip.edu.au/SIF3InfraREST/hits/environments/environment",
            id,
            id,
            // "9d35716d628a4cdabcb37f61ae7cad4e",
        )
        client.SetDebug(debug)
        client.SetAuthenticationMethod_SIF_HMACSHA256()
        client.CreateEnviroment()
        fmt.Println("SessionToken", client.GetSessionToken())
        data := client.Get(provider)
        return c.XMLBlob(http.StatusOK, data)
    })

	e.Logger.Fatal(e.Start(":8089"))
}
