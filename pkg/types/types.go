package types

type Money int64

type PaymentCategory string

type PaymentStatus string

const (
	StatusOk         PaymentStatus = "OK"
	StatusFail       PaymentStatus = "FAIL"
	StatusInProgress PaymentStatus = "INPROGRESS"
)

type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type PAN string

type Card struct {
	ID       int64
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
