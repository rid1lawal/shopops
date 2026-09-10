module "eks" {
  source             = "terraform-aws-modules/eks/aws"
  version            = "21.25.0"
  name               = "shopops"
  kubernetes_version = "1.34"

  vpc_id     = var.vpc_id
  subnet_ids = var.private_subnet_ids

  compute_config = {
    enabled    = true
    node_pools = ["general-purpose"]
  }

  enable_cluster_creator_admin_permissions = true

  tags = {
    Environment = var.environment
    Project     = "shopops"
  }
}
