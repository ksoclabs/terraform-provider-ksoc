package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/ksoclabs/terraform-provider-ksoc/internal/auth"
)

func AuthenticatedRequest(
	ctx context.Context,
	apiUrlBase string,
	httpMethod string,
	targetUrl string,
	accessKeyID string,
	secretKey string,
	payload any,
) (statusCode int, body []byte, diags diag.Diagnostics) {
	authenticator := auth.New(apiUrlBase)
	token, err := authenticator.Authenticate(ctx, accessKeyID, secretKey)
	if err != nil {
		return 500, nil, append(diags, diag.Errorf("Error authenticating using Access Key %s: %s", accessKeyID, err)...)
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return 500, nil, append(diags, diag.Errorf("Error occurred while encoding data: %s", err)...)
	}

	jsonPayload := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(httpMethod, targetUrl, jsonPayload)
	if err != nil {
		return 500, nil, append(diags, diag.Errorf("Error in http request: %s", err)...)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, append(diags, diag.Errorf("Error reading response. Response code: %d", resp.StatusCode)...)
	}

	if resp.StatusCode >= 300 {
		return resp.StatusCode, nil, append(diags, diag.Errorf("HTTP request error. Response code: %d,  Response body: %s", resp.StatusCode, string(bodyBytes))...)
	}

	return resp.StatusCode, bodyBytes, diags
}
