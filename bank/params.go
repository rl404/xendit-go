package bank

// InquiryBankParams contains parameters for InquiryBank.
type InquiryBankParams struct {
	BankAccountNumber string `json:"bank_account_number" validate:"required"`
	BankCode          string `json:"bank_code" validate:"required"`
}
