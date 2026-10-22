package main

import "github.com/golang-jwt/jwt/v5"

type BeidJWTClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}
