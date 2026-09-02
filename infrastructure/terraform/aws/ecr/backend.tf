terraform {
  backend "s3" {
    bucket         = "shopops-terraform-states-bucket"
    key            = "ecr-state.tfstate"
    region         = "eu-west-1"
    use_lockfile   = true
  }
}