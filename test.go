package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a password with the following criteria:
// - Atleast one uppercase
// - Atleast one lowercase
// - Atleast one special character
// - Atleast one number/digit.
// - Password length should be equal to 8chars.

// visited = {
// 	1
// }

//

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const specialChars = "*@#$%^&!"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomLowercase(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = lowercase[seededRand.Intn(len(lowercase))]
	}
	return string(b)
}

func generateRandomUppercase(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = uppercase[seededRand.Intn(len(uppercase))]
	}
	return string(b)
}

func generateRandomDigits(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = digits[seededRand.Intn(len(digits))]
	}
	return string(b)
}
func generateRandomSpecialChars(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = specialChars[seededRand.Intn(len(specialChars))]
	}
	return string(b)
}

func generateRandomPass(length int, randomString string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randomString[seededRand.Intn(len(randomString))]
	}
	return string(b)
}

func passwordGenerator() string {
	lowercaseChars := generateRandomLowercase(2)
	uppercaseChars := generateRandomUppercase(2)
	digits := generateRandomDigits(2)
	specialChars := generateRandomSpecialChars(2)
	randomChars := lowercaseChars + uppercaseChars + digits + specialChars
	pass := generateRandomPass(8, randomChars)
	return pass
}

func isPrime(num int) bool {
	for j := 2; j < num/2; j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}

func printPrime(num int) {
	for i := 1; i <= num; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
