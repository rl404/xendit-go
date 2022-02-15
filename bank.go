package xendit

import "time"

// BankAccount contains data from Xendit's API response of inquiry bank account.
type BankAccount struct {
	ID                    string    `json:"id"`
	Reference             string    `json:"reference"`
	BankCode              string    `json:"bank_code"`
	BankAccountNumber     string    `json:"bank_account_number"`
	BankAccountHolderName string    `json:"bank_account_holder_name"`
	Status                string    `json:"status"`
	FailureReason         string    `json:"failure_reason"`
	Updated               time.Time `json:"updated"`
}
