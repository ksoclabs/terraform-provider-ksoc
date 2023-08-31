# KSOC Terraform Provider
This is the official Terraform Provider for KSOC. Use this provider to interact with the KSOC api. The provider can be found on the [Terraform Provider Registery](https://registry.terraform.io/providers/ksoclabs/ksoc/latest).

To configure the provider, you will need a set of cloud api keys. The keys consist of an access and a secret key that can be generated from the KSOC platform.

To connect your AWS account to your KSOC account, create a `ksoc_aws_register` resource where you run terraform for your AWS resources.

An example of leveraging this resource can be found in our terraform module examples directory [here](https://github.com/ksoclabs/terraform-aws-ksoc-connect/blob/main/examples/main.tf)
