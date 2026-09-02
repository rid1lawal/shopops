resource "aws_ecr_repository" "shopops_ecr" {
  name                 = "shopops"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}