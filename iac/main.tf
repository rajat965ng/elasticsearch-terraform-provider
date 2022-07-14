resource "elasticsearch_index" "my_collection" {
  doc_id  = 3
  name    = "employee"
  payload = "{\"name\":\"user d\", \"age\":22,\"department\":\"finance\"}"
}
