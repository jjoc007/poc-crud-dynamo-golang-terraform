resource "aws_lambda_permission" "lambda_permission_person_post_rest" {
  depends_on    = [aws_lambda_function.poc_create_person_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_create_person_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_person_post_method.http_method}${aws_api_gateway_resource.poc_person_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_person_get_rest" {
  depends_on    = [aws_lambda_function.poc_get_person_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_get_person_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_person_get_method.http_method}${aws_api_gateway_resource.poc_person_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_person_put_rest" {
  depends_on    = [aws_lambda_function.poc_update_person_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_update_person_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_person_put_method.http_method}${aws_api_gateway_resource.poc_person_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_person_delete_rest" {
  depends_on    = [aws_lambda_function.poc_delete_person_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_delete_person_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_person_delete_method.http_method}${aws_api_gateway_resource.poc_person_resource.path}"
}


resource "aws_lambda_permission" "lambda_permission_animal_post_rest" {
  depends_on    = [aws_lambda_function.poc_create_animal_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_create_animal_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_animal_post_method.http_method}${aws_api_gateway_resource.poc_animal_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_animal_get_rest" {
  depends_on    = [aws_lambda_function.poc_get_animal_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_get_animal_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_animal_get_method.http_method}${aws_api_gateway_resource.poc_animal_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_animal_put_rest" {
  depends_on    = [aws_lambda_function.poc_update_animal_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_update_animal_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_animal_put_method.http_method}${aws_api_gateway_resource.poc_animal_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_animal_delete_rest" {
  depends_on    = [aws_lambda_function.poc_delete_animal_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_delete_animal_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_animal_delete_method.http_method}${aws_api_gateway_resource.poc_animal_resource.path}"
}