# ElasticSearch Terraform Provider
- ``es-terraform`` has plugin code.
- ``iac`` has tf files for provisioning.

## Steps to build provider
- cd ``es-terraform``
- ```go mod tidy```
- ```go mod fmt```
- ```go build -o terraform-provider-elasticsearch```

## For Mac users
- ```cp terraform-provider-elasticsearch ~/.terraform.d/plugins/terraform-example.com/elasticsearchprovider/elasticsearch/1.0.0/darwin_amd64/```

## For Linux users
- ```cp terraform-provider-elasticsearch ~/.terraform.d/plugins/terraform-example.com/elasticsearchprovider/elasticsearch/1.0.0/linux_amd64/```

## Boostrap Elasticsearch using docker-compose
- ```docker-compose up```

## Provision index in Elasticsearch
- cd ```iac```
- ```terraform init```
- ```terraform plan```
- ```terraform apply -auto-approve=true```