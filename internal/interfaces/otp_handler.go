package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diazharizky/rest-otp-generator/internal/application"
	"github.com/diazharizky/rest-otp-generator/internal/domain"
	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
	"github.com/go-chi/chi"
)

const (
	otpMsgValid   = "Your OTP is valid."
	otpMsgInvalid = "Your OTP is invalid."
)

type resCode struct {
	Code string `json:"code"`
}

type otpHandler struct {
	oa application.OTPAppInterface
}

func newOTPHandler(oa application.OTPAppInterface) otpHandler {
	return otpHandler{oa}
}

func (o *otpHandler) getHandler() (r *chi.Mux) {
	r = chi.NewRouter()
	basePath := "/{key}"
	r.Post(basePath, o.generateOTP)
	r.Put(fmt.Sprintf("%s/verify", basePath), o.verifyOTP)
	return
}

func (o *otpHandler) generateOTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var p domain.OTP
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	p.Key = chi.URLParam(r, "key")
	p.Attempts = 0
	p.FixProps()
	code, err := o.oa.GenerateOTP(&p)
	if err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	httpUtils.ResponseSuccess(w, resCode{Code: code})
}

func (o *otpHandler) verifyOTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var p domain.OTP
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	if len(p.Code) != int(p.Digits) {
		httpUtils.ResponseBadRequest(w, []string{otpMsgInvalid})
		return
	}

	p.Key = chi.URLParam(r, "key")
	p.FixProps()
	valid, err := o.oa.VerifyOTP(&p)
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
