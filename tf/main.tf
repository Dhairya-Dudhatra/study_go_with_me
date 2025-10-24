module "stackgen_37d5eb4a-9e76-4054-b124-655c55735424" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "assgaesgaege"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

