locals {
  environment = "dev"
}
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "6.7.2"

  name = "shopops-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b"]
  public_subnets  = ["10.0.1.0/24", "10.0.2.0/24"]
  private_subnets = ["10.0.11.0/24", "10.0.12.0/24"]

  enable_nat_gateway = true
  single_nat_gateway = true

  tags = {
    Environment = local.environment
  }
}

module "eks" {
  source = "../../modules/eks"

  vpc_id = module.vpc.vpc_id

  private_subnet_ids = module.vpc.private_subnets

  environment = local.environment
}