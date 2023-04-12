package main

import (
  "crypto/ecdsa"
  "crypto/elliptic"
  "crypto/rand"
  "crypto/sha256"
  "encoding/hex"
  "fmt"
  "math/big"

  "golang.org/x/crypto/ripemd160"
)

var secp256k1curve = elliptic.P256()
var one = new(big.Int).SetInt64(1)

func main() {
  pk, err := ecdsa.GenerateKey(secp256k1curve, rand.Reader)
  if err != nil {
    panic(err)
  }

  address := addy(pk)
  fmt.Println("Address:", address)
}

func addy(pk *ecdsa.PrivateKey) string {
  pubkey := pk.PublicKey
  pubkeyBytes := elliptic.Marshal(secp256k1curve, pubkey.X, pubkey.Y)

  hash := sha256.New()
  hash.Write(pubkeyBytes)
  pubkey2 := hash.Sum(nil)

  hash = ripemd160.New()
  hash.Write(pubkey2)
  pubkey3 := hash.Sum(nil)

  pubkey4 := append([]byte{0x00}, pubkey3...)
  hash = sha256.New()
  hash.Write(pubkey4)
  pubkey5 := hash.Sum(nil)

  hash = sha256.New()
  hash.Write(pubkey5)
  pubkey6 := hash.Sum(nil)

  checksum := pubkey6[:4]
  pubkey7 := append(pubkey4, checksum...)

  return base58Encode(pubkey7)
}

func base58Encode(input []byte) string {
  alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
  bigZero := big.NewInt(0)
  bigFiftyEight := big.NewInt(58)

  x := new(big.Int).SetBytes(input)
  var result []byte

  for x.Cmp(bigZero) > 0 {
    mod := new(big.Int)
    x.DivMod(x, bigFiftyEight, mod)
    result = append(result, alphabet[mod.Int64()])
  }

  for _, v := range input {
    if v == 0x00 {
      result = append(result, alphabet[0])
    } else {
      break
    }
  }

  // Reverse the encoded data
  for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
    result[i], result[j] = result[j], result[i]
  }

  return string(result)
}

