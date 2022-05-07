# play with mongoDB

> bring up the stack (compose file)

```bash
docker-compose -f stack.yml up -d
```

> copy sample data to mongoDB container

```bash
docker cp movies.json a82ecd165f4c:/movies.json
```

> import data and check the result (mongo shell)

```bash
docker exec -it a82ecd165f4c bash
# in the container
mongoimport -u root -p root --authenticationDatabase admin --db sample_mflix --collection movies --file movies.json
mongosh -u root -p root --authenticationDatabase admin
```
