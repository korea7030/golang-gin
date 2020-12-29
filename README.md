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