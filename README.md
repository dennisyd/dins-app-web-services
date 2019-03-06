# lasagna-msa
Backend microservices for lasagna

* Identity service (Login/Signup/3rd party Oauth)
* Internal Recipe Service (CRUD recipes in our database)
* External Recipe Service (Queries external services for data and returns data to client; for example, add a recipe from a url — would scrape info from the url and return)
* Recommendation Service (Takes a user id and gets user info from identity service, then uses user preferences to query internal recipe service for weighted results)
