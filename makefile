



.PHONY: start end net mysql redis 


start:
	@docker-compose up -d


end:
	@docker-compose -f ./docker-compose.yml down

.PHONY: net
net:
	@docker inspect mysql | grep IPAddress

mysql:
	@docker exec -it mysql bash 

redis:
	@docker exec -it redis7 bash

# redis-cli
# 替代
# docker-compose -f ./docker-compose.yml exec redis sh -c 'redis-cli'