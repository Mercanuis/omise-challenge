// Package http contains code that calls the API for Omise
//
// The usage of this API is using test credentials: any requests are made on a test environment
// and no real transactions are created. Because these are test credentials, and this code is
// not meant to be used for commercial or real-world use, these credentials are hard-coded for
// the sake of the exercise. If in production, these credentials should be pulled from the environment
// for security's sake.
//
// omise contains the main API code, with interfaces and functions used to call the Omise API descirbed here
//
// https://www.omise.co/charges-api
//
// https://www.omise.co/cards-api
package http
