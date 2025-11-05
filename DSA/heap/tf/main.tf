module "stackgen_d8ba30d6-d397-44fb-936d-f787fab565bb" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "zdfvgdzfbv"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

