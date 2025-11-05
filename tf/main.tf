module "stackgen_2e035ccf-df65-456c-ab41-8773c675d656" {
  source                           = "./modules/aws_athena"
  athena_engine_version            = "AUTO"
  bucket_name                      = "sdfvdsfv"
  bytes_scanned_cutoff_per_query   = 10485760
  database_force_destroy           = false
  database_name                    = "sdvadsd"
  encrypt_query_results            = true
  encryption_option                = "SSE_KMS"
  expected_bucket_owner            = null
  kms_key                          = null
  query                            = "zsdvasvas"
  query_description                = null
  query_name                       = "zxvds"
  query_results_encryption_option  = "SSE_KMS"
  require_encryption_configuration = false
  result_output_location           = null
  set_acl_configuration            = false
  tags                             = {}
  workgroup_description            = null
  workgroup_force_destroy          = false
  workgroup_name                   = "zsvasf"
  workgroup_state                  = "ENABLED"
}

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

