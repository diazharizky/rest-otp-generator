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

	p.Key = chi.URLParam(r, "key")
	if p.Digits > 6 {
		p.Digits = 6
	}

	if err = mCore.generateOTP(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{"passcode": p.Passcode}
	rByte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rByte)
}

func verifyOTPHandler(w http.ResponseWriter, r *http.Request) {
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

	message := "Your OTP is invalid!"
	if len(p.Passcode) != int(p.Digits) {
		w.WriteHeader(400)
		w.Write([]byte(message))
		return
	}

	p.Key = chi.URLParam(r, "key")
	if err = mCore.verifyOTP(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message = "Your OTP is valid!"
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}
