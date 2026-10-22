package main

import "github.com/golang-jwt/jwt/v5"

/*
* Issuer: beid
* Subject: user id
* Audience: request app(can either be a single string or an array)
* ExpiresAt: now + time(access: 5min, refresh: 30day)
* NotBefore: now
* IssuedAt: now
* ID: random bytes
* DeviceID: device id
* Use: access or refresh
* Ver: v1
 */
type BeidJWTClaims struct {
	DeviceID string `json:"did,omitempty"`
	Use      string `json:"use,omitempty"`
	Ver      string `json:"ver,omitempty"`
	jwt.RegisteredClaims
}
