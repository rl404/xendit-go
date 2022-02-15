package bank

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke bank API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// InquiryBank to inquiry bank account.
func (c *Client) InquiryBank(data *InquiryBankParams) (*xendit.BankAccount, *xendit.Error) {
	return c.InquiryBankWithContext(context.Background(), data)
}

// InquiryBankWithContext to inquiry bank account with context.
func (c *Client) InquiryBankWithContext(ctx context.Context, data *InquiryBankParams) (*xendit.BankAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.BankAccount{}
	header := &http.Header{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/bank_account_data_requests", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
