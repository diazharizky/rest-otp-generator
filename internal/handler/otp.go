package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	core "github.com/diazharizky/rest-otp-generator/internal/core"
	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

const (
	otpMsgValid   = "Your OTP is valid."
	otpMsgInvalid = "Your OTP is invalid."
)

type pscode struct {
	Code string `json:"code"`
}

type otpManager interface {
	GenerateOTP(*otp.OTPBase) (string, error)
	VerifyOTP(*otp.OTPV) (bool, error)
}

type otpDriver struct {
	otpManager
}

var odr otpDriver

func init() {
	odr = otpDriver{&core.Core}
}

func otpHandler() (r *chi.Mux) {
	r = chi.NewRouter()
	basePath := "/{key}"
	r.Post(basePath, generateOTPHandler)
	r.Put(fmt.Sprintf("%s/verify", basePath), verifyOTPHandler)
	return
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
	p.FixParams()
	code, err := odr.GenerateOTP(&p)
	if err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	httpUtils.ResponseSuccess(w, pscode{Code: code})
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

	if len(p.Code) != int(p.Digits) {
		httpUtils.ResponseBadRequest(w, []string{otpMsgInvalid})
		return
	}

	p.Key = chi.URLParam(r, "key")
	p.FixParams()
	valid, err := odr.VerifyOTP(&p)
	if err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	if !valid {
		httpUtils.ResponseBadRequest(w, []string{otpMsgInvalid})
		return
	}
	httpUtils.ResponseSuccess(w, otpMsgValid)
}
