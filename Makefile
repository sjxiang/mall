


start:
	@docker-compose up -d

end:
	@docker-compose -f ./docker-compose.yml down

net:
	@docker inspect mysql | grep IPAddress

mysql:
	@echo '打开 MySQL 控制台'
# @docker exec -it mysql bash 
	@docker-compose -f ./docker-compose.yml exec mysql sh -c 'mysql -uroot -p'

redis:
	@echo '打开 Redis 控制台'
	@docker-compose -f ./docker-compose.yml exec redis7 sh -c 'redis-cli'


.PHONY: start end net mysql redis 
