module "stackgen_e1dc5bff-7c94-4db2-8bca-26dc9b8e9d19" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "asfaefa"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

