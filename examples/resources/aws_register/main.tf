resource "ksoc_aws_register" "this" {
  ksoc_api_url          = "<api endpoint>"
  ksoc_assumed_role_arn = "arn:aws:iam::<aws_account_number>:role/ksoc-connector"
  access_key_id         = "ksoc_access_key"
  secret_key            = "ksoc_secret_key"
  aws_account_id        = "aws_account_id"
}
