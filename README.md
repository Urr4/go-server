This is a small example-project to tryout go/golang.

This was implemented to see if I can use go to do the following:

* Use modules to split my code
* Read a config-file
* Use external modules (like the UUID module)
* Build a REST-Endpoint that can read Authorization-Headers and Query-Params
* Reject requests with a wrong Authorization header
* Write data to a postgres db
* Execute a request agains another REST-Server

## How to start
The service expects a postgres running at port 5432 with default user "postgres" and password
Just run `docker run --name postgres -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres` and create the following table:

```sql
CREATE TABLE users 
(
    id VARCHAR NOT NULL,
    name VARCHAR NOT NULL    
)
```

To run just execute `go run application.go` in the root directory, this will start the server on post 8090.

Executing `curl -v localhost:8090/hello?name=Peter -H "Authorization: Bearer top-secret"` should get you a 200,
instead executing `curl -v localhost:8090/hello?name=Peter -H "Authorization: Bearer dont-know"` should result in a 403.

