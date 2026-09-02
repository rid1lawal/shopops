resource "aws_s3_bucket" "shopops_bucket" {
  bucket = "shopops-terraform-states-bucket"

  tags = {
    Name        = "Shopops bucket"
  }

  object_lock_enabled = true
}

resource "aws_s3_bucket_versioning" "shopops_bucket_versioning" {
  bucket = aws_s3_bucket.shopops_bucket.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "shopops_bucket_encryption" {
  bucket = aws_s3_bucket.shopops_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_public_access_block" "shopops_bucket_public_access_block" {
  bucket = aws_s3_bucket.shopops_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}