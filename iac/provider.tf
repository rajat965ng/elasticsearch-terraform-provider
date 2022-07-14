terraform {
  required_providers {
    elasticsearch = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/elasticsearchprovider/elasticsearch"
    }
  }
}
