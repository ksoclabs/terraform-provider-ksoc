resource "ksoc_aws_register" "this" {
  ksoc_assumed_role_arn = "arn:aws:iam::<aws_account_number>:role/ksoc-connector"
  aws_account_id        = "aws_account_id"
}
