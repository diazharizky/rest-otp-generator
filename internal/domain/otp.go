package domain

type OTP struct {
	Key         string
	Period      uint   `json:"period"`
	Digits      uint   `json:"digits"`
	Code        string `json:"code"`
	MaxAttempts uint   `json:"max_attempts"`
	Attempts    uint   `json:"attempts"`
}

func (p *OTP) FixProps() {
	if p.Period < 60 {
		p.Period = 60
	}
	if p.Period > 300 {
		p.Period = 300
	}
	if p.Digits < 3 {
		p.Digits = 3
	}
	if p.Digits > 6 {
		p.Digits = 6
	}
	if p.MaxAttempts < 3 {
		p.MaxAttempts = 3
	}
	if p.MaxAttempts > 5 {
		p.MaxAttempts = 5
	}
}

func (p OTP) ValidateProps() []string {
	errMessages := []string{}
	if len(p.Code) <= 0 {
		errMessages = append(errMessages, "OTP code required.")
	}
	return errMessages
}
