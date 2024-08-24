module github.com/helloworld/go-app

go 1.22.5

replace example.com/app => ./app

require (
	github.com/aws/aws-sdk-go v1.55.5
	github.com/gofor-little/env v1.0.18
)

require (
	example.com/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
