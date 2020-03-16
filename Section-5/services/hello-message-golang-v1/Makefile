NAME = hello-message-golang:1.3.1

default: run

image:
	docker build -t $(NAME) .

push:
	docker tag $(NAME) lreimer/$(NAME)
	docker push lreimer/$(NAME)

run:
	go run main.go
