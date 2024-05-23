NAME=jump_assessment

all: $(NAME)

$(NAME):
	docker-compose -f docker-compose.yml -f docker-compose-prod.yml up -d --build

dev:
	docker-compose up -d --build
	make logs

logs:
	docker-compose logs -f server

ps:
	docker-compose ps

re: fclean dev

fclean: clean
	docker image rm jump_server_image

clean:
	docker-compose down