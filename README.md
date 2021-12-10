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
Stop:  
```sh
docker-compose down
```
  
# What is going on there
What does the client do:  
- sending address to download products (for example, http://164.92.251.245:8080/api/v1/products/)  
- requesting products with pagination functionality (limit, skip)  
- requesting products continuously, simulate endless loading (based on bidirectional-stream)  
  
What does the server do:  
- going to the url, download and save the csv-file, extract the products from it, save to your database  
- giving products with pagination functionality (limit, skip)  
- giving away products in stream  
  
Services are raised in docker-compose, the client starts and executes requests with output to the console, and then exits.  
The server continues to work  

# Separate services
[Client](client)  
[Server](server)  

Example external csv-product store source (**if still works**)  
[swagger-doc](http://164.92.251.245:8080/api/v1/products/)  
[products](http://164.92.251.245:8080/api/v1/products/)  
  
If first resource isn't work - check this example:  
[swagger-doc](https://github.com/p12s/csv-create-api/tree/master/docs)
