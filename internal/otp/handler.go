package otp

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

const (
	otpMessageValid   = "Your OTP is valid."
	otpMessageInvalid = "Your OTP is invalid."
)

func Handler() (r *chi.Mux) {
	r = chi.NewRouter()
	basePath := "/{key}"
	r.Post(basePath, generateOTPHandler)
	r.Put(fmt.Sprintf("%s/verify", basePath), verifyOTPHandler)
	return
}

type genOTPRes struct {
	Passcode string `json:"passcode"`
}

func generateOTPHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p otp.OTPBase
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	v := validator.New()
	if err = v.Struct(p); err != nil {
		httpUtils.ResponseBadRequest(w, []string{err.Error()})
		return
	}

	p.Key = chi.URLParam(r, "key")
	p.Attempts = 0
	p.SetDefaultValues()
	passcode, err := c.generateOTP(&p)
	if err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	httpUtils.ResponseSuccess(w, genOTPRes{Passcode: passcode})
}

func verifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var p otp.OTPV
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	v := validator.New()
	if err = v.Struct(p); err != nil {
		httpUtils.ResponseBadRequest(w, []string{err.Error()})
		return
	}

	if len(p.Passcode) != int(p.Digits) {
		httpUtils.ResponseBadRequest(w, []string{otpMessageInvalid})
		return
	}

	p.Key = chi.URLParam(r, "key")
	p.SetDefaultValues()
	if err = c.verifyOTP(&p); err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	httpUtils.ResponseSuccess(w, otpMessageValid)
}
