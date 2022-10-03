package main

import (
	"log"
	"math/rand"
	"time"
)

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o Otp) GenAndSendOTP(length int) error {
	otp := o.iOtp.GenRandomOTP(length)
	o.iOtp.SaveOTPCache(otp)
	message := o.iOtp.GetMessage(otp)
	err := o.iOtp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

type SmsOtp struct{}

func (s *SmsOtp) GenRandomOTP(len int) string {
	randomOTP := RandStringRunes(len)
	log.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *SmsOtp) SaveOTPCache(otp string) {
	log.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *SmsOtp) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *SmsOtp) SendNotification(message string) error {
	log.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type EmailOtp struct{}

func (s *EmailOtp) GenRandomOTP(len int) string {
	randomOTP := RandStringRunes(len)
	log.Printf("Email: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *EmailOtp) SaveOTPCache(otp string) {
	log.Printf("Email: saving otp: %s to cache\n", otp)
}

func (s *EmailOtp) GetMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (s *EmailOtp) SendNotification(message string) error {
	log.Printf("Email: sending sms: %s\n", message)
	return nil
}

var (
	r           = rand.New(rand.NewSource(time.Now().UnixMilli()))
	letterRunes = []rune("0123456789")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	otp := Otp{
		iOtp: &SmsOtp{},
	}
	otp.GenAndSendOTP(4)

	otp = Otp{
		iOtp: &EmailOtp{},
	}
	otp.GenAndSendOTP(8)
}
