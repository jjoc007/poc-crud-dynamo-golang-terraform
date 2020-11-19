resource "aws_api_gateway_rest_api" "poc_rest_api" {
  name        = "poc_rest_api_crud"
  description = "API Gateway"
}

resource "aws_api_gateway_deployment" "poc_api_rest_dev_development" {
  depends_on = [
    aws_api_gateway_integration.poc_person_post_integration,
    aws_api_gateway_integration.poc_person_get_integration,
    aws_api_gateway_integration.poc_person_put_integration,
    aws_api_gateway_integration.poc_person_delete_integration,

    aws_api_gateway_integration.poc_animal_post_integration,
    aws_api_gateway_integration.poc_animal_get_integration,
    aws_api_gateway_integration.poc_animal_put_integration,
    aws_api_gateway_integration.poc_animal_delete_integration,
  ]

  rest_api_id = aws_api_gateway_rest_api.poc_rest_api.id
  stage_name  = "dev"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_resource" "poc_person_resource" {
  rest_api_id = aws_api_gateway_rest_api.poc_rest_api.id
  parent_id   = aws_api_gateway_rest_api.poc_rest_api.root_resource_id
  path_part   = "person"
}

resource "aws_api_gateway_method" "poc_person_post_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_person_resource.id
  http_method   = "POST"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_person_post_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_person_resource.id
  http_method             = aws_api_gateway_method.poc_person_post_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_create_person_lambda.invoke_arn
}

resource "aws_api_gateway_method" "poc_person_get_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_person_resource.id
  http_method   = "GET"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_person_get_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_person_resource.id
  http_method             = aws_api_gateway_method.poc_person_get_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_get_person_lambda.invoke_arn
}

resource "aws_api_gateway_method" "poc_person_put_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_person_resource.id
  http_method   = "PUT"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_person_put_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_person_resource.id
  http_method             = aws_api_gateway_method.poc_person_put_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_update_person_lambda.invoke_arn
}

resource "aws_api_gateway_method" "poc_person_delete_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_person_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_person_delete_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_person_resource.id
  http_method             = aws_api_gateway_method.poc_person_delete_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_delete_person_lambda.invoke_arn
}

resource "aws_api_gateway_resource" "poc_animal_resource" {
  rest_api_id = aws_api_gateway_rest_api.poc_rest_api.id
  parent_id   = aws_api_gateway_rest_api.poc_rest_api.root_resource_id
  path_part   = "animal"
}

resource "aws_api_gateway_method" "poc_animal_post_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_animal_resource.id
  http_method   = "POST"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_animal_post_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_animal_resource.id
  http_method             = aws_api_gateway_method.poc_animal_post_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_create_animal_lambda.invoke_arn
}

resource "aws_api_gateway_method" "poc_animal_get_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_animal_resource.id
  http_method   = "GET"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_animal_get_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_animal_resource.id
  http_method             = aws_api_gateway_method.poc_animal_get_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_get_animal_lambda.invoke_arn
}

resource "aws_api_gateway_method" "poc_animal_put_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_animal_resource.id
  http_method   = "PUT"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_animal_put_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_animal_resource.id
  http_method             = aws_api_gateway_method.poc_animal_put_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_update_animal_lambda.invoke_arn
}

resource "aws_api_gateway_method" "poc_animal_delete_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_animal_resource.id
  http_method   = "DELETE"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_animal_delete_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_animal_resource.id
  http_method             = aws_api_gateway_method.poc_animal_delete_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_delete_animal_lambda.invoke_arn
}