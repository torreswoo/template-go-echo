# template-go-echo

```bash
$ docker build -t torreswoo/template-go-echo:latest .
$ docker push torreswoo/template-go-echo:latest
$ docker run -p 3100:3100 torreswoo/template-go-echo:latest
```

```bash
$ kubectl apply -f k8s/manifest.yaml
$ kubectl port-forward  svc/test-api-svc 3100:3100
```