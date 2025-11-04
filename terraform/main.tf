module "stackgen_e243a5ef-f7cb-4d88-955d-1d70c0787338" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "afasfas"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

