package otp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

func Handler() (r *chi.Mux) {
	r = chi.NewRouter()

	basePath := "/{key}"
	r.Post(basePath, generateOTPHandler)
	r.Put(fmt.Sprintf("%s/verify", basePath), verifyOTPHandler)

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

	p.Attempts = 3
	p.Key = chi.URLParam(r, "key")
	if p.Digits > 6 {
		p.Digits = 6
	}

	if p.Period <= 60 {
		p.Period = 60 * time.Second
	}

	if err = c.generateOTP(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{"passcode": p.Passcode}
	rjs, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(rjs)
}

func verifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p otp.OTPV
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

	message := "Your OTP is invalid!"
	if len(p.Passcode) != int(p.Digits) {
		w.WriteHeader(400)
		w.Write([]byte(message))
		return
	}

	p.Key = chi.URLParam(r, "key")
	if err = c.verifyOTP(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message = "Your OTP is valid!"
	w.WriteHeader(200)
	w.Write([]byte(message))
}
