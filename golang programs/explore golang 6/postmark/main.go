package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mrz1836/postmark"
)

func main() {
	fmt.Println("learning postmark go client")

	// postmark is paid, so ig it can not be tested ha ha.
	// TODO: learn templated email in golang and use postmark to send it.

	serverToken := "bf89873f-ae26-44ae-b46a-0bc5338a15fa"
	accountToken := "ee7690eb-bb31-48ce-8d0c-d07a1bee09d6"

	client := postmark.NewClient(serverToken, accountToken)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	templateVariables := `{
	"product_url": "product_url_Value",
	"product_name": "product_name_Value",
	"name": "name_Value",
	"action_url": "action_url_Value",
	"login_url": "login_url_Value",
	"username": "username_Value",
	"trial_length": "trial_length_Value",
	"trial_start_date": "trial_start_date_Value",
	"trial_end_date": "trial_end_date_Value",
	"support_email": "support_email_Value",
	"live_chat_url": "live_chat_url_Value",
	"sender_name": "sender_name_Value",
	"help_url": "help_url_Value",
	"company_name": "company_name_Value",
	"company_address": "company_address_Value"
}`

	var variables map[string]interface{}
	err := json.Unmarshal([]byte(templateVariables), &variables)
	if err != nil {
		fmt.Println("error unmarshalling template variables", err)
		return
	}

	res, err := client.SendTemplatedEmail(ctx, postmark.TemplatedEmail{
		TemplateID:    36975312,
		TemplateModel: variables,
		From:          "ketan.rathod@bacancy.com",
		To:            "tridip.chavda@bacancy.com",
	})

	if err != nil {
		fmt.Println("error sending email", err)
		return
	}

	fmt.Println("postmark response:", res)
}
