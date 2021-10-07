package otp

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Handler() *chi.Mux {
	r := chi.NewRouter()
	basePath := "/{key}"
	r.Post(basePath, generateOTP)
	r.Put(basePath+"/verifications", verifyOTP)
	return r
}

func generateOTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Your OTP code is ..."))
}

func verifyOTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Your OTP is invalid!"))
}
