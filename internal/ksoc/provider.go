package ksoc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"ksoc_api_url": {
					Type:        schema.TypeString,
					Description: "Ksoc API to target. Defaults to https://api.ksoc.com",
					Default:     "https://api.ksoc.com",
					Optional:    true,
				},
				"access_key_id": {
					Type:        schema.TypeString,
					Description: "Ksoc Customer Access ID",
					ForceNew:    true,
					Required:    true,
					Sensitive:   true,
				},
				"secret_key": {
					Type:        schema.TypeString,
					Description: "Ksoc Customer Secret Key",
					ForceNew:    true,
					Required:    true,
					Sensitive:   true,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"ksoc_aws_register":   resourceAwsRegister(),
				"ksoc_azure_register": resourceAzureRegister(),
			},
			ConfigureContextFunc: configureProvider,
		}
		return p
	}
}

type Config struct {
	KsocApiUrl    string
	KsocAccountId string
	AccessKeyId   string
	SecretKey     string
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		KsocApiUrl:  d.Get("ksoc_api_url").(string),
		AccessKeyId: d.Get("access_key_id").(string),
		SecretKey:   d.Get("secret_key").(string),
	}

	return &config, nil
}

type RegistrationPayload struct {
	Type        string      `json:"type"`
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	AzureSubscription AzureSubscriptionCredential `json:"azure_subscription"`
	AWSAccount        AWSAccountCredential        `json:"aws_account"`
}

type AWSAccountCredential struct {
	AWSAccountID string `db:"aws_account_id" json:"aws_account_id"`
	AWSRoleArn   string `db:"aws_role_arn" json:"aws_role_arn"`
}

type AzureSubscriptionCredential struct {
	TenantID       string `json:"tenant_id"`
	SubscriptionID string `json:"subscription_id"`
}
