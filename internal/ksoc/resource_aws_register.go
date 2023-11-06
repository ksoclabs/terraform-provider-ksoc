package ksoc

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ksoclabs/terraform-provider-ksoc/internal/request"
)

func resourceAwsRegister() *schema.Resource {
	return &schema.Resource{
		Description: "Register service with Ksoc",

		CreateContext: resourceAwsRegisterCreate,
		ReadContext:   resourceAwsRegisterRead,
		UpdateContext: resourceAwsRegisterUpdate,
		DeleteContext: resourceAwsRegisterDelete,

		Schema: map[string]*schema.Schema{
			"ksoc_assumed_role_arn": {
				Type:        schema.TypeString,
				Description: "Ksoc Role to Trust",
				//ForceNew:    true,
				Required: true,
			},
			"aws_account_id": {
				Type:        schema.TypeString,
				Description: "Ksoc Customer AWS account ID",
				ForceNew:    true,
				Required:    true,
				Sensitive:   true,
			},
			"ksoc_registered": {
				Type:        schema.TypeBool,
				Description: "Target of the API path",
				Computed:    true,
			},

			// Computed values
			"api_path": {
				Type:        schema.TypeString,
				Description: "Target of the API path",
				Computed:    true,
			},
		},
	}
}

func resourceAwsRegisterCreate(ctx context.Context, d *schema.ResourceData, meta any) (diags diag.Diagnostics) {
	config := meta.(*Config)
	httpMethod := http.MethodPost
	setValueOnSuccess := config.KsocApiUrl
	diags = resourceAwsRegisterGeneric(ctx, httpMethod, d, setValueOnSuccess, meta)
	return diags
}

func resourceAwsRegisterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	config := meta.(*Config)
	apiUrlBase := config.KsocApiUrl
	targetURI := apiUrlBase + "/cloud/aws/register"
	err := d.Set("api_path", targetURI)
	if err != nil {
		return diag.Errorf("Error setting api_path: %s", err)
	}
	return nil
}

func resourceAwsRegisterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// Update has not yet been implemented
	return nil
}

func resourceAwsRegisterDelete(ctx context.Context, d *schema.ResourceData, meta any) (diags diag.Diagnostics) {
	httpMethod := http.MethodDelete
	setValueOnSuccess := ""
	diags = resourceAwsRegisterGeneric(ctx, httpMethod, d, setValueOnSuccess, meta)
	return diags
}

func resourceAwsRegisterGeneric(ctx context.Context, httpMethod string, d *schema.ResourceData, setValueOnSuccess string, meta any) (diags diag.Diagnostics) {
	config := meta.(*Config)
	apiUrlBase := config.KsocApiUrl

	targetURI := apiUrlBase + "/cloud/register"
	accessKey := config.AccessKeyId
	secretKey := config.SecretKey
	awsAccountID := d.Get("aws_account_id").(string)

	payload := &RegistrationPayload{
		Type: "aws",
		Credentials: Credentials{
			AWSAccount: AWSAccountCredential{
				AWSAccountID: awsAccountID,
				AWSRoleArn:   "arn:aws:iam::" + awsAccountID + ":role/ksoc-connect",
			},
		},
	}

	statusCode, _, diags := request.AuthenticatedRequest(ctx, apiUrlBase, httpMethod, targetURI, accessKey, secretKey, payload)
	if statusCode != http.StatusOK {
		return append(diags, diag.Errorf("Failed to register with KSOC, received HTTP status: %d", statusCode)...)
	}

	err := d.Set("api_path", targetURI)
	if err != nil {
		return diag.Errorf("Error setting api_path: %s", err)
	}

	if err := d.Set("ksoc_registered", statusCode == http.StatusOK); err != nil {
		return append(diags, diag.Errorf("Error setting ksoc_registered: %s", err)...)
	}

	d.SetId(setValueOnSuccess)

	return nil
}
