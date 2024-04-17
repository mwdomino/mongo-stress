# Stress test for MongoDB in a cluster...

### Running it

``` sh
docker-compose up -d
docker-compose run app
```

The `app` container will crash on first start as the leader election hasn't completed between the 3 mongo instances. Once it's running, fail the instances to ensure the nodes continue to process requests:

``` sh
alias dc=docker-compose
dc kill mongo2 ; sleep 5 ; dc restart mongo2 ; sleep 5 ; dc kill mongo1 ; sleep 5 ; dc restart mongo1 ; sleep 10 ; dc kill mongo3 ; sleep 5 ; dc restart mongo3
```

Once it's complete, open up a `mongosh` and run the following to verify all documents were stored:

``` sh
mongosh mongodb://localhost:27017

use testdb
db.numbers.countDocuments()
```
