# TapToEat

#### how to run

1. copy env from env.example

```
cp .env .env.example
```

2. set `port`,`dbname`, `user` and `root` from your local database
3. get dependencies from this project

```
go get
```

4. run project using this command

```
make run
```

5. on default, port from this project is `8800`, so copy this curl on postman or add new terminal and paste

```
curl --request GET http://localhost:8800/v1/hello
```
