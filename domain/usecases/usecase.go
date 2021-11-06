package usecases

type Usecase interface {
	Login(string, string) (Account, error)
	ListAccounts() ([]AccountID, error)
}
