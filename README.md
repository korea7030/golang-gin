### 참고 강좌
https://www.youtube.com/watch?v=Ypwv1mFZ5vU&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=1

### 폴더 구조
```
project
|   go.mod
│   go.sum
│   README.md
│   server.go
│   .gitignore
│───api
│   │   video-api.go
└───controller
│   │   video-controller.go
│   │   login-controller.go
│───docs
│   │   docs.go
│   │   swagger.json
│   │   swagger.yaml
│───dto
│   │   credentials.go (jwt credentials)
│   │   response.go
│───entity
│   │   video.go
│───service
│   │   video-service.go
│   │   jwt-service.go
│   │   login-service.go
│───middlewares
│   │   basic-auth.go
│   │   jwt-auth.go
│   │   logger.go
│───templates
│   │───css
│   │   │   index.css
│   │   footer.html
│   │   header.html
│   │   index.html
│───validators
│   │   validators.go
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

#### Docker with ElasticBeanstalk
1. Dockerfile 생성
2. Docker image 생성(생성한 Dockerfile 이미지로)
3. Dockerrun.aws.json 파일 생성
```json
// Dockerrun.aws.json
{
  "AWSEBdockerrunVersion": "1",
  "Image": {
    "Name": "{ImageName}"
  },
  "Ports": [
    {
      "ContainerPort": "{portNumber}"
    }
  ]
}
```
4. .ebignore 파일 생성(docker로 생성하기 떄문에 eb에 소스들을 올리면 안됨)

#### jwt go
https://github.com/dgrijalva/jwt-go
#### JWT 참고
http://www.opennaru.com/opennaru-blog/jwt-json-web-token/

#### Golang ORM(gorm)
https://gorm.io/  
go ORM 비교 : https://blog.billo.io/devposts/go_orm_recommandation/


#### REST API swagger
1. API comment 등록 
2. swag library 다운로드
```bash
$ go get -u github.com/swaggo/swag/cmd/swag
```
3. go root folder에서 다음 명령어 실행. 주의사항: 무조건 main.go 파일이 있어야 한다.  
다른파일명을 root source로 할경우 실행 안됨
```
$ swag init
```
4. gin-swagger 다운로드
```
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```
5. import 
```go
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files" // swagger embed files
```
github 사이트 : https://github.com/swaggo/gin-swagger