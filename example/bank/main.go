package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/bank"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	bankData := bank.InquiryBankParams{
		BankAccountNumber: "123123",
		BankCode:          "BCA",
	}

	resp, err := bank.InquiryBank(&bankData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bank account: %+v\n", resp)
}
