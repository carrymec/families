1. 创建用户

```curl
curl --location 'http://127.0.0.1:8080/api/v1/create_person' \
--header 'Content-Type: application/json' \
--data '{
    "name": "秦庄襄王",
    "birthdate": "前281年－前247年"
}'
```

2. 创建用户带关系

```shell
curl --location 'http://127.0.0.1:8080/api/v1/create_person' \
--header 'Content-Type: application/json' \
--data '{
    "name": "秦王政",
    "birthdate": "前259年－前210年",
    "relation": {
        "relationId": 430,
        "relationType": "son"
    }
}'
```
3. 条件查询
```shell
curl --location 'http://127.0.0.1:8080/api/v1/query_persons' \
--header 'Content-Type: application/json' \
--data '{
    "name": "政",
    "page": 1,
    "pageSize": 10
}'
```
4. 单个查询
```shell
curl --location 'http://127.0.0.1:8080/api/v1/persons/431'
```