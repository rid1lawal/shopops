terraform {
  backend "s3" {
    bucket         = "shopops-terraform-states-bucket"
    key            = "dev-state.tfstate"
    region         = "eu-west-1"
    use_lockfile   = true
  }
}