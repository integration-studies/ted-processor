````shell script

curl -v "http://localhost:8080" \
    -X POST \
    -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
    -H "Ce-Specversion: 1.0" \
    -H "Ce-Type: dev.knative.samples.helloworld" \
    -H "Ce-Source: dev.knative.samples/helloworldsource" \
    -H "Content-Type: application/json" \
    -d '{"msg":"Hello World from the curl pod."}'
````