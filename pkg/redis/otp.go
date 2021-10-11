package redis

type OTPValue struct {
	Passcode string `json:"passcode" redis:"passcode"`
	Attempts int8   `json:"attempts" redis:"attempts"`
}
