module "stackgen_6df2622d-2236-4775-80ab-39f5b6c27207" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "zfvdsgdsgsdd"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

