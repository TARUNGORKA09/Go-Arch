package main

import (
	"fmt"
	"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {

		return fmt.Errorf("Token doesnt matches")
	
	}
	if u.SessionID == 0 {
		return fmt.Errorf("Invalid claims")
	}
	return nil
}


var Key []byte ={"Tarun"}

func createToken(c *UserClaims) (string,error) {                  // it is used by a user to send the request....it may be contain in the Authorization header of the request .
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedtoken , err := t.SignedString(Key)
	if err != nil {
		fmt.Errorf("Error while generating token")
	}
	return signedtoken,_
}

var claims *UserClaims = &UserClaims{} 

func ParseToken(signedtoken string) (*UserClaims,error) {

	t,err := jwt.ParseWithclaims(signedtoken,claims,func verify(tok *jwt.Token) (interface{},error){
		tok.Methods.Alg() != jwt.SigningMethodES512.Alg() {
			return nil, fmt.Errorf("Invalid algorithm used...means aletered")
		}
		return Key,err
	})
	if err != nil {
		return nil, fmt.Errorf("error in parse token")
	}
	if  !t.Valid(){
		return nil,t
	}
	return t.claims.(*userClaims),nil
}



func main() {

}
