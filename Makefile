# set env value NAME
NAME = "raid2earn"
# set env value VERSION
VERSION = "1.0.0"

build:
	go build -o bin/$(NAME) -ldflags "-X main.Version=$(VERSION)" $(NAME).go

run:
	make build
	./bin/$(NAME)

start:
	./bin/$(NAME)