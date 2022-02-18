package main

func main() {
	smsNotify := NewNotify(&Sms{})
	sms := NewSms(smsNotify)
	sms.genRandomCode(3)
	emailNotify := NewNotify(&Email{})
	email := NewEmail(emailNotify)
	email.genRandomCode(6)
}
