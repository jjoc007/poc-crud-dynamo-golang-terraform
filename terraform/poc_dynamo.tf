resource "aws_dynamodb_table" "poc_person_table" {
  name           = "poc_person_table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "uid"

  attribute {
    name = "uid"
    type = "S"
  }

}

resource "aws_dynamodb_table" "poc_animal_table" {
  name           = "poc_animal_table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "animal_name"

  attribute {
    name = "animal_name"
    type = "S"
  }

}