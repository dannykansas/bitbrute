package main

import (
  "crypto/ecdsa"
  "crypto/elliptic"
  "crypto/rand"
  "strings"
  "testing"
)

func TestBase58Encode(t *testing.T) {
  testCases := []struct {
    input    []byte
    expected string
  }{
    {[]byte{0}, "1"},
    {[]byte{255}, "5Q"},
    {[]byte{0, 1, 2, 3, 4}, "11A7YP"},
    {[]byte{250, 251, 252, 253, 254}, "4zzVvYg"},
  }

  for _, tc := range testCases {
    result := base58Encode(tc.input)
    if result != tc.expected {
      t.Errorf("base58Encode(%v) = %v, expected %v", tc.input, result, tc.expected)
    }
  }
}

func TestAddy(t *testing.T) {
  secp256k1curve := elliptic.P256()
  pk, err := ecdsa.GenerateKey(secp256k1curve, rand.Reader)
  if err != nil {
    t.Fatal(err)
  }

  address := addy(pk)
  if len(address) < 25 || len(address) > 34 {
    t.Errorf("Invalid address length: %d", len(address))
  }

  if !strings.HasPrefix(address, "1") {
    t.Errorf("Invalid address prefix: %s", string(address[0]))
  }
}

