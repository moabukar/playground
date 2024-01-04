terraform {
  backend "s3" {
    bucket         = "mo-tf-statefile"
    key            = "terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "mo-tf-lock"
  }
}
