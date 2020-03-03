// Package sifclient provides all the necessary authentication,
// paging and infrastructure handling to connect to a SIF 3 REST Service.
// It is agnostic of the data returned.
package sifclient

import (
    "fmt"
)

//
//
// Notes:
//
// * Inputs - minimum
//     solutionId
//     applicationKey
//     userToken
//     authMethod
//
// * Create environment
//     * Send necessary payload
//         * Start with TEXT template
//         * Replace with Struct
//     * Deal with error on existing environment
//     * Read existing
//     * Delete
//     * Store return data
//         * Paths etc
//
// * GET
//     * Get single
//     * Get many
//     * Dealing with pagination
//
//
// * generateToken
//     * generateTokenStandard
//     * generateTokenHMAC256
//
// * generateXMLEnvironment
//

func Info() {
    fmt.Println("SIFCLIENT: Information")
}
