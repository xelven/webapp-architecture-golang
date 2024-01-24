## webapp-architecture-golang
Purpose is create an Web Application foundation construction on production level, any product/projects could reuse it from 0 to 1 in the future.
* First version focus on Restful API service.
* The idiomatic structure based on the resource-oriented design.
* The usage of Docker, Docker compose, Alpine images, and linters on development.
* unit test with mock db
* key foundation:
  * code level db replicate implementation. 
  * Server syslog.
  * structure design flexible & easy to scale up.
 
* todo list: 
  * find a way to do the api parameters validation more consistent & smarter 
  * add user table /  user sign-up/in api.
  * user auth check in router middleware with jwt check
  * create PostgreSQL script for init enable UUID & create multiple user with different role. (create real slave for local)
  * Benchmark big sql query with SQL Builder.
  * inlcuding sentry.
  * Server syslog with datadog.
  * didn't decide which where should unit test should put it(right beside the source file or different folder) https://www.reddit.com/r/golang/comments/u23z67/why_are_tests_file_created_right_beside_the/
  * CI / CD integration with circleci https://circleci.com/.
  * documents generation didn't decide.
  * looking for local bucket solution for GCP services, something like localstack provided.
  * 

The idiomatic structure based on the resource-oriented design.

## 📦 container image size 
* DB: 249MB
* API: 
  * Development environment: 666MB
  * Production environment: 19MB ; 💡`docker build -f Dockerfile . -t allen_webapp_core`

## 🗂️ folder structure
```
xxxxapp
├── cmd
│  ├── api
│  │  └── main.go // Restful API entry point
│  ├── apix
│  │  └── main.go // Restful API x entry point
│  ├── websocket
│  │  └── main.go // websocket API entry point, split if is in-need.
│  └── migrate 
│     └── main.go // db migration entry point
│
├── core
│  ├── middleware
│  │  └── content_type_json.go // anything custom need during api call, such as content_type,basic auth,CROS 
│  ├── router
│  │  └── router.go  //api router list
│  │
│  └── resource or domain 
│     ├── health   // provide funciton for health check
│     │  └── handler.go  
│     ├── product   // domian product provide all product related function
│     │  ├── handler.go     // business logic handling, called from router
│     │  ├── model.go       // all this (product) domain related struct
│     │  └── repository.go  // all db related CUD here
│     │  └── querier.go     // all db related R here <- read only
│     └── common
│        └── err 
│           └── err.go  // common error handle response
│
├── migrations
│  └── 20240123020903_create_products_table.sql
│
├── config
│  └── config.go
│
├── util
│  ├── logger
│  │  └── logger.go // system logger
│  └──  test
│     └── test.go // provide common method for unit test
│
├── .env  // local env
│
├── go.mod
├── go.sum
│
├── docker-compose.yml
├── Dockerfile.dev //local dev, hot reload included
└── Dockerfile // production build
```


