// Package sifclient provides all the necessary authentication,
// paging and infrastructure handling to connect to a SIF 3 REST Service.
// It is agnostic of the data returned.
package sifclient

import (
	"fmt"
)

func Info() {
	fmt.Println("SIFCLIENT: Information")
}

// func Example() {
//     client := sifclient.New(
//         "https://hits.nsip.edu.au/SIF3InfraREST/hits/environments/environment",
//         "9d35716d628a4cdabcb37f61ae7cad4e",
//     )
//     client.SetAuthenticationMethod_SIF_HMACSHA256()
//     client.CreateEnviroment()
//     school := client.Get("SchoolInfos")
//     schoolStr := string(school)
//     fmt.Println("SchoolInfos get", schoolStr)
//}
