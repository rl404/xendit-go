package bank

import (
	"context"

	"github.com/xendit/xendit-go"
)

// InquiryBank to inquiry bank account.
func InquiryBank(data *InquiryBankParams) (*xendit.BankAccount, *xendit.Error) {
	return InquiryBankWithContext(context.Background(), data)
}

// InquiryBankWithContext to inquiry bank account.
func InquiryBankWithContext(ctx context.Context, data *InquiryBankParams) (*xendit.BankAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.InquiryBankWithContext(ctx, data)
}
