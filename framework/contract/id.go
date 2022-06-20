package contract

const IDKey = "ginx:id"

type IDService interface {
	NewID() string
}
