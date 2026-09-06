data "terraform_remote_state" "secrets_manager" {
  backend = "s3"

  config = {
    bucket = "shopops-terraform-states-bucket"
    key    = "secrets-manager-state.tfstate"
    region = "eu-west-1"
  }
}