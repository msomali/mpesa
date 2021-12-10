package mpesa

type api interface {
	session
	disburser
	push
	reversal
	b2b
	directDebit
}
