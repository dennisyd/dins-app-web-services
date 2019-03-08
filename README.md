# lasagna-msa
Backend microservices for lasagna

* Identity service (Login/Signup/3rd party Oauth)
* Internal Recipe Service (CRUD recipes in our database)
* External Recipe Service (Queries external services for data and returns data to client; for example, add a recipe from a url — would scrape info from the url and return)
* Recommendation Service (Takes a user id and gets user info from identity service, then uses user preferences to query internal recipe service for weighted results)


## API RPC Routes
URL:
```
localhost:8080
```
APPLICATION/JSON POST:
```json
{
	"service":"lasagna.identity.service",
	"method": "IdentityService.Create",
	"request":{
		"email": "doesitwqork@gmail.com",
		"password": "passasdasword"
	}
}
```

## Stats
View live stats at
```
localhost:8080/stats
```
