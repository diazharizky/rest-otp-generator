package otp

import (
	"encoding/base32"
	"encoding/json"
	"net/http"
	"time"

	myredis "github.com/diazharizky/rest-otp-generator/pkg/redis"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type MyOTP struct {
	DB myredis.RDB
}

type BaseOTPPayload struct {
	Key    string
	Period time.Duration `json:"period"`
	Digits int8          `json:"digits"`
}

type VerifyOTPPayload struct {
	BaseOTPPayload

	Passcode string `json:"passcode"`
}

var myOTP MyOTP

func init() {
	myOTP.DB.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Handler() (r *chi.Mux) {
	r = chi.NewRouter()
	basePath := "/{key}"
	r.Post(basePath, myOTP.GenerateOTP)
	r.Put(basePath+"/verifications", myOTP.VerifyOTP)
	return
}

func (o *MyOTP) GenerateOTP(w http.ResponseWriter, r *http.Request) {
	var p BaseOTPPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		panic(err)
	}

	p.Key = chi.URLParam(r, "key")
	if p.Digits > 6 {
		p.Digits = 6
	}

	secret := base32.StdEncoding.EncodeToString([]byte(p.Key))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period: uint(p.Period),
		Digits: otp.Digits(p.Digits),
	})
	if err != nil {
		panic(err)
	}

	if err = o.DB.Set(p.Key, passcode, p.Period*time.Second); err != nil {
		panic(err)
	}

	res := map[string]interface{}{"passcode": passcode}
	resByte, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resByte)
}

func (o *MyOTP) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var p VerifyOTPPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		panic(err)
	}

	p.Key = chi.URLParam(r, "key")
	secret := base32.StdEncoding.EncodeToString([]byte(p.Key))
	valid, err := totp.ValidateCustom(p.Passcode, secret, time.Now(), totp.ValidateOpts{
		Period: uint(p.Period),
		Digits: otp.Digits(p.Digits),
	})
	if err != nil {
		panic(err)
	}

	message := "Your OTP is invalid!"
	if !valid {
		w.Write([]byte(message))
		return
	}

	exists, err := o.DB.Get(p.Key)
	if err != nil {
		panic(err)
	}

	if len(exists) <= 0 {
		w.Write([]byte(message))
		return
	}

	if err = o.DB.Remove(p.Key); err != nil {
		panic(err)
	}

	message = "Your OTP is valid!"
	w.Write([]byte(message))
}
