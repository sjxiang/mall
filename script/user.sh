

curl --location 'localhost:8888/api/user/signup' \
--header 'Content-Type: application/json' \
--data '{
    "username": "",
    "password": "1234",
    "re_password": "1234",
    "gender": 2
}'