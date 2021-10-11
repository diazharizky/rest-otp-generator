package otp

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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

	rVal := myredis.OTPValue{Passcode: passcode, Attempts: 3}
	fVal, err := toMSI(rVal)
	if err != nil {
		panic(err)
	}

	if err = o.DB.HSet(p.Key, fVal, p.Period*time.Second); err != nil {
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

	var d myredis.OTPValue
	err = o.DB.HGetAll(p.Key).Scan(&d)
	if err != nil {
		panic(err)
	}

	v := reflect.ValueOf(d)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s", typeOfS.Field(i).Name)
	}

	if err = o.DB.HRemove(p.Key); err != nil {
		panic(err)
	}

	message = "Your OTP is valid!"
	w.Write([]byte(message))
}

func toMSI(val interface{}) (interface{}, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	var i map[string]interface{}
	if err = json.Unmarshal(b, &i); err != nil {
		return nil, err
	}

	return i, nil
}
