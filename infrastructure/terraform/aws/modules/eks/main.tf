module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "21.25.0"

  name               = "shopops"
  kubernetes_version = "1.33"

  endpoint_public_access = true

  enable_cluster_creator_admin_permissions = true

  vpc_id = var.vpc_id

  subnet_ids = var.private_subnet_ids

  enable_irsa = true

  addons = {
    coredns = {
      most_recent = true
    }

    kube-proxy = {
      most_recent = true
    }

    vpc-cni = {
      most_recent = true
    }

    eks-pod-identity-agent = {
      most_recent = true
    }
  }

  eks_managed_node_groups = {
    shopops = {
      name = "shopops"

      instance_types = ["t3.small"]

      min_size     = 1
      max_size     = 3
      desired_size = 1

      subnet_ids = var.private_subnet_ids

      capacity_type = "ON_DEMAND"
    }
  }

  tags = {
    Environment = var.environment
    Project     = "shopops"
  }
}