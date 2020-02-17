NAME = hello-istio-golang:1.2.1

default: run

image:
	docker build -t $(NAME) .

push:
	docker tag $(NAME) lreimer/$(NAME)
	docker push lreimer/$(NAME)

run:
	go run main.go
