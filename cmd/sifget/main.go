package main

// Use api and Test
import (
	"fmt"
	// sifclient sifclient-go
	"encoding/xml"
	"github.com/nsip/sifclient-go"
	"os"
)

func main() {
	printXML := false

	fmt.Println("Testing")
	sifclient.Info()

	// minimum = url environment, applicationKey, password/secret
	client := sifclient.New(
		"https://hits.nsip.edu.au/SIF3InfraREST/hits/environments/environment",

		// SCOTT TEST
		// "330d73072e1848cc95a90f5942ad73d9",
		// "330d73072e1848cc95a90f5942ad73d9",

		// EXAMPLE - HMAC 2 school, lots of students
		"9d35716d628a4cdabcb37f61ae7cad4e",
		"9d35716d628a4cdabcb37f61ae7cad4e",

		// EXAMPLE - BASIC 2 school, lots of students
		// "50372a7c82a84391a1300857167324a8",
		// "50372a7c82a84391a1300857167324a8",
	)
	client.SetAuthenticationMethod_SIF_HMACSHA256()

	fmt.Println(client.GetAuthToken())

	e := client.GetEnviromentRequest()
	if printXML {
		output, err := xml.MarshalIndent(e, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		os.Stdout.Write(output)
	}

	// XXX make sure this works with HMAC too
	fmt.Println("CREATE ENVIRONMENT")
	client.CreateEnviroment()
	if printXML {
		output, err := xml.MarshalIndent(client.LastEnvironment, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		os.Stdout.Write(output)
	}
	fmt.Println("SessionToken", client.GetSessionToken())

	// get many vs get one, pagination etc etc
	fmt.Println("GET SCHOOLINFO")
	school := client.Get("SchoolInfos")
	schoolStr := string(school)
	fmt.Println("SchoolInfos get", schoolStr)

	fmt.Println("GET STUDENTPERSONALS")
	students := client.Get("StudentPersonals")
	studentStr := string(students)
	fmt.Println("StudentPersonals get", len(studentStr))
}
