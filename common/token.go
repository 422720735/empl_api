package common

import (
	"empl_api/aes"
	"empl_api/defs"
	"encoding/base64"
	"encoding/json"
	"time"
)

type Token struct {
	AdminMananger defs.Users
	Sign          string // 签名
	Indate        int64  // 有效期
	Time          int64  // 时间
}

func NewToken(user *defs.Users, sign, aesKey string, indate int64) (string, error) {
	token := &Token{
		AdminMananger: *user,
		Sign:          sign,
		Indate:        indate,
		Time:          time.Now().Unix(),
	}
	ciphertext,err:=token.Encrypt(aesKey)
	if err!=nil{
		return "",err
	}
	return ciphertext,nil
}

/*
*加密一个验证码
 */
func (t *Token)Encrypt(aesKey string)(string, error)  {
	key := aesKey
	enc := &aes.AesEncrypt{}
	buffer, err := json.Marshal(*t)
	if err != nil{
		return "",err
	}
	src,_ := enc.Encrypt(key,string(buffer))
	encodeString := base64.StdEncoding.EncodeToString(src)
	return encodeString,nil


}