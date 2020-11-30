// Package cipher implements struct and functions with relation to the decryption/encryption of files
//
// cipher implements a simple Caesar Cipher, where the cipher key is 128 rotation
// the implementation is based on the code here
//
// https://play.golang.org/p/dCWYyWPHwj4
//
// decipher implements an interface to decipher files
//
// encipher implements an interface to encipher files
// This is meant to be a helper interface, used during testing and validation
//
// rot128 is the implementation of the 128 Cipher
package cipher
