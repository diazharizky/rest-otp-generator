package domain

const (
	MinPeriod      = uint(60)
	MaxPeriod      = uint(300)
	MinDigits      = uint(3)
	MaxDigits      = uint(6)
	MinMaxAttempts = uint(3)
	MaxMaxAttemtps = uint(5)

	OTPCodeRequiredMsg = "OTP code required."
	MismatchOTPLenMsg  = "OTP code's length doesn't equal with digit value."
)

type OTP struct {
	Key         string
	Period      uint   `json:"period"`
	Digits      uint   `json:"digits"`
	Code        string `json:"code"`
	MaxAttempts uint   `json:"max_attempts"`
	Attempts    uint   `json:"attempts"`
}

func (p *OTP) FixProps() {
	if p.Period < MinPeriod {
		p.Period = MinPeriod
	}
	if p.Period > MaxPeriod {
		p.Period = MaxPeriod
	}
	if p.Digits < MinDigits {
		p.Digits = MinDigits
	}
	if p.Digits > MaxDigits {
		p.Digits = MaxDigits
	}
	if p.MaxAttempts < MinMaxAttempts {
		p.MaxAttempts = MinMaxAttempts
	}
	if p.MaxAttempts > MaxMaxAttemtps {
		p.MaxAttempts = MaxMaxAttemtps
	}
}

func (p OTP) ValidateProps() []string {
	errMessages := []string{}
	codeLen := len(p.Code)
	if codeLen <= 0 {
		errMessages = append(errMessages, OTPCodeRequiredMsg)
	} else {
		if codeLen != int(p.Digits) {
			errMessages = append(errMessages, MismatchOTPLenMsg)
		}
	}
	return errMessages
}
