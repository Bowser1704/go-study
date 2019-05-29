package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //cost == 10  Typeof(hashBytes) == []byte
	return string(hashedBytes), err
	//生成token
}

func Compare(token, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(token), []byte(password))	//token == hashPassword
}