module "stackgen_bd6bdcfe-4fb9-4475-87a3-53a303d10979" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "sdfsdhbsthsthtrh"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

