module "stackgen_6e777fca-4788-4cac-ad6f-293ff76e3831" {
  source                       = "./modules/aws_s3"
  block_public_access          = true
  bucket_name                  = "asrgaegaradrgar"
  bucket_policy                = ""
  enable_versioning            = true
  enable_website_configuration = false
  sse_algorithm                = "aws:kms"
  tags                         = {}
  website_error_document       = "404.html"
  website_index_document       = "index.html"
}

