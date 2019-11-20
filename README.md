# template-go-echo

## start up
```
#. use Go version
$ gvm use go1.11
 
#. Basic Go mod
$ go mod init github.com/torreswoo/template-go-echo
$ go get -u github.com/labstack/echo
$ go mod download
($ GO111MODULE=auto)
 
$ go mod tidy [-v]   # go.mod 파일과 소스코드를 비교하여, import 되지 않은 의존성은 제거하고, import 되었지만 의존성 리스트에 추가되지 않은 모듈은 추가
$ go mod vendor [-v]  # vendor/ 디렉터리를 생성합니다.
$ go mod verify    # 로컬에 설치된 모듈의 해시 값과 go.sum을 비교하여 모듈의 유효성을 검증
$ go list -m all
 
#. run
$ go run cmd/server/main.go
 
#. make binary & start
$ go build -o bin/server cmd/server/main.go
$ bin/server
```

## Docker & k8s
```bash
$ docker build -t torreswoo/template-go-echo:latest .
$ docker push torreswoo/template-go-echo:latest
$ docker run -p 3100:3100 torreswoo/template-go-echo:latest
```

```bash
$ kubectl apply -f k8s/manifest.yaml
$ kubectl port-forward  svc/test-api-svc 3100:3100
```
