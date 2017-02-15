package tokens

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	// Acquire token:
	sema <- struct{}{}

	balance += amount

	// Release token:
	<-sema
}

func Balance() int {
	sema <- struct{}{}

	b := balance

	<-sema

	return b
}
