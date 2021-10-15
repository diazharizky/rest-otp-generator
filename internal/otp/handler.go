package otp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

func Handler() (r *chi.Mux) {
	r = chi.NewRouter()
	basePath := "/{key}"
	r.Post(basePath, generateOTPHandler)
	r.Put(fmt.Sprintf("%s/verify", basePath), verifyOTP)
	return
}

func generateOTPHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p otp.OTP
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	v := validator.New()
	if err = v.Struct(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.Key = chi.URLParam(r, "key")
	if p.Digits > 6 {
		p.Digits = 6
	}

	passcode, err := c.generateOTP(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{"passcode": passcode}
	resByte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resByte)
}

func verifyOTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var p otp.OTP
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	message := "Incomplete parameter"
	v := validator.New()
	if err = v.Struct(p); err != nil {
		w.Write([]byte(message))
		return
	}

	message = "Your OTP is invalid!"
	if len(p.Passcode) != int(p.Digits) {
		w.Write([]byte(message))
		return
	}

	p.Key = chi.URLParam(r, "key")
	if err := c.verifyOTP(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message = "Your OTP is valid!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}
