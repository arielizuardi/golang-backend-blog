# golang-backend-blog

To start the application:
```
bash ./start.sh
make migrate.up
```

To test the connection:
```
curl -X GET http://localhost:8080/ping
```

You should see
```
{
    "pong":"ok"
}
```

Perform integration & e2e test:
```
make docker.start
make test.integration
```

Perform unit test:
```
make test.unit
```
