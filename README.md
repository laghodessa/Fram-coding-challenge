# Personia

## API

### Update hierarchy

```http
PUT http://localhost:3000/api/hierarchy
Authorization: Bearer secret
Content-Type: application/json

{
  "Pete": "Nick",
  "Barbara": "Nick",
  "Nick": "Sophie",
  "Sophie": "Jonas"
}
```

### Get hierarchy

```http
GET http://localhost:3000/api/hierarchy
Authorization: Bearer secret

```

### Get supervisor

```http
GET http://localhost:3000/api/hierarchy/Nick
Authorization: Bearer secret
```


## Running with docker

```sh
# build
docker build -t personia .
# run
docker run -p 3000:3000 personia
```

Sample requests:
```sh
# update hierarchy
curl -X PUT http://localhost:3000/api/hierarchy -i \
  -H 'Authorization: Bearer secret' \
  -H 'Content-Type: application/json' \
  -d '
    {
      "Pete": "Nick",
      "Barbara": "Nick",
      "Nick": "Sophie",
      "Sophie": "Jonas"
    }
  '

# get hierarchy
curl http://localhost:3000/api/hierarchy \
  -H 'Authorization: Bearer secret'


# get Nick's superviors
curl http://localhost:3000/api/hierarchy/Nick \
  -H 'Authorization: Bearer secret'
```
