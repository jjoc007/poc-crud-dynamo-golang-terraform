#!/bin/bash

cd ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas
CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_create/main functions/person/create/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_create/main.zip build/person_create/main

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_delete/main functions/person/delete/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_delete/main.zip build/person_delete/main

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_get/main functions/person/get/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_get/main.zip build/person_get/main

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_update/main functions/person/update/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/person_update/main.zip build/person_update/main



CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_create/main functions/animal/create/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_create/main.zip build/animal_create/main

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_delete/main functions/animal/delete/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_delete/main.zip build/animal_delete/main

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_get/main functions/animal/get/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_get/main.zip build/animal_get/main

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_update/main functions/animal/update/main.go
build-lambda-zip.exe -o ~/dev/guild-golang/poc-crud-dynamo-golang-terraform/lambdas/build/animal_update/main.zip build/animal_update/main