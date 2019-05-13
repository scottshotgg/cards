package main

import (
	"fmt"
	"log"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	fmt.Println("hey hey")

	stripe.Key = "sk_test_BQokikJOvBiI2HlWgH4olfQ2"

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	params.SetSource("tok_visa")
	params.AddMetadata("key", "value")

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.ID)

	card, err := card.New(&stripe.CardParams{})
	if err != nil {

	}

	fmt.Printf("card %+v", card)
}
