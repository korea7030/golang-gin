### 참고 강좌
https://www.youtube.com/watch?v=Ypwv1mFZ5vU&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=1

### 폴더 구조
```
project
|   go.mod
│   go.sum
│   README.md
│   server.go
└───controller
│   │   video-controller.go
│───entity
│   │   video.go
│───service
│   │   video-service.go
```

#### package
- request 와 response의 header/body를 dump 해주는 Gin Middleware/handler  
github : https://github.com/tpkeeper/gin-dump (gin-dump)

#### Gin Model binding and validation
1. Model bind
  > - Bind, BindJSON, BindXML, BindQuery, BindYAML 이 있음
  > - request로 넘어온 데이터 중 bind할 대상이 없으면 404 error return
2. Should bind
  > - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML 로 구성
  > - bind 에러 발생 시 개발자가 처리 가능