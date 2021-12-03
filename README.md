<div align="center">
<article style="display: flex; flex-direction: column; align-items: center; justify-content: center;">
  <h1 style="width: 100%; text-align: center;">Product store</h1>
  <p>Product store. Task description is <a href="task.md">here</a></p>
</article>
</div>

# 🔥 Run
Download:
```sh
git clone https://github.com/p12s/product-store.git && cd product-store
mv server/.env.example server/.env
mv client/.env.example client/.env

docker-compose up --build
```
Run gRPC-client:
```sh
cd client
run cmd/main.go
```
Stop:  
```sh
docker-compose down
```

There is a bug with migrations - they are not performed automatically as they should.  
You have to run the command manually.
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

# Separate services
[Client](client)  
[Server](server)  

Example external csv-product store source (**if still works**)  
[swagger-doc](http://164.92.251.245:8080/api/v1/products/)  
[products](http://164.92.251.245:8080/api/v1/products/)  
  
If first resource isn't work - check this example:  
[swagger-doc](https://github.com/p12s/csv-create-api/tree/master/docs)

