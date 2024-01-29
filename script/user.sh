
# 注册
curl --location 'localhost:8888/api/user/signup' \
--header 'Content-Type: application/json' \
--data '{
    "username": "tyr",
    "password": "1234",
    "re_password": "1234",
    "gender": 2
}'

# 登录
curl --location 'localhost:8888/api/user/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "tyr",
    "password": "1234"
}'

# 详情
curl --location --request GET 'localhost:8888/api/user/detail?user_id=1706503790' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcxMDg2MDQsImlhdCI6MTcwNjUwMzgwNCwidXNlcl9pZCI6MTcwNjUwMzc5MH0.gMqUOEtjlrhRkZfYMyirX1Hq9yj1EFBH1gM09C_hj4g' 
