module "stackgen_47f90a8b-f70a-40dd-a41f-0ff98808dcec" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "dfdfgdfgdgdd"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

