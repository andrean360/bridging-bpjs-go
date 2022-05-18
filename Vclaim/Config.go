package Vclaim

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	cons_id    = os.Getenv("CONS_ID")
	Secret_key = os.Getenv("SECRET_KEY")
	User_key   = os.Getenv("USER_KEY")
)

func SetHeader() (string, string, string, string) {

	timenow := time.Now().UTC()
	t, err := time.Parse(time.RFC3339, "1970-01-01T00:00:00Z")
	if err != nil {
		log.Fatal(err)
	}

	tstamp := timenow.Unix() - t.Unix()
	secret := []byte(Secret_key)
	message := []byte(cons_id + "&" + fmt.Sprint(tstamp))
	hash := hmac.New(sha256.New, secret)
	hash.Write(message)
	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))
	// to base64
	X_signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return cons_id, User_key, strconv.FormatInt(tstamp, 16), X_signature

}
