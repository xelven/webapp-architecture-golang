# webapp-architecture-golang
Purpose is create an Web Application foundation construction on production level, any product/projects could reuse it from 0 to 1 in the future.
First version focus on Restful API service. 


# folder structure
```
xxxxapp
├── cmd
│  ├── api
│  │  └── main.go // Restful API entry point 
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
│     │  └── repository.go  // all db related cud here
│     │  └── querier.go     // all db related r here
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
└── Dockerfile.dev
```


