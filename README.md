# web-service
Backend microservices for dins.app

* Identity service (Login/Signup/3rd party Oauth)
* Internal Recipe Service (CRUD recipes in our database)
* External Recipe Service (Queries external services for data and returns data to client; for example, add a recipe from a url — would scrape info from the url and return)
* Recommendation Service (Takes a user id and gets user info from identity service, then uses user preferences to query internal recipe service for weighted results)


## API RPC Routes
URL:
```
localhost:8080/rpc
```
APPLICATION/JSON POST:
```json
{
	"service":"dins.app.api.v1.identity",
	"method": "Identity.Create",
	"request":{
		"email": "doesitwqork@gmail.com",
		"password": "passasdasword",
		"first_name": "test"
	}
}
```

### OR

URL:
```
localhost:8080/identity/create
```
APPLICATION/JSON POST:
```json
{
	"email": "doesitwqork@gmail.com",
	"password": "passasdasword",
	"first_name": "test"
}
```


## Stats
View live stats at
```
localhost:8080/stats
```
