package mail

import (
	"crypto/rand"
	"math/big"
  )
  
  func GenerateOtp(length int) (string) {
  
	const charset = "0123456789"
  
	otp := make([]byte,length)
  
	for i := range otp {
  
	  random, err := rand.Int(rand.Reader,big.NewInt(int64(len(charset))))
  
	  if err != nil {
  
		return ""
	  }
  
	  otp[i] = charset[random.Int64()]
	}
  
	return string(otp)
  }