# [{{ .Service.Namespace }}] {{ .Service.Name }}
> Powered by Godin {{ .Godin.Version }} ({{ .Godin.Commit }})

* **Godin init:** {{ .Project.Created }}
* **Last godin update :** {{ .Project.Updated }}

## gRPC service: {{ .Protobuf.Service }}
{{- range .Service.Methods }}
**{{ .Name }}**
{{ range .Comments }}
*{{ replace "// " "" . }}*
{{- end }}
```go
func {{ .Name }}({{ .ParamList }}) ({{ .ReturnList }})
```
{{- end }}

{{- if .Service.Transport.AMQP }}
## AMQP subscriptions

All handlers are located in: `./internal/service/subscriber/`
Each handler has it's own file, named after the subscription topic.

| **Routing key** | **Exchange** | **Queue** | **Handler** |
|-----------------|--------------|-----------|-------------|
{{- range .Service.Subscriber }}
| {{ .Subscription.Topic }} | {{ .Subscription.Exchange }} | {{ .Subscription.Queue.Name }} | {{ .Handler }} |
{{- end }}
{{- end }}

## Transport Options
| **Option**      | **Enabled**                                                                          |
|--------------|----------------------------------------------------------------------------------|
| gRPC Transport layer | {{ ReadmeOptionCheckbox "transport.grpc.enabled" }} |
| gRPC Server | {{ ReadmeOptionCheckbox "transport.grpc.server.enabled" }} |
| gRPC Client | {{ ReadmeOptionCheckbox "transport.grpc.client.enabled" }} |
| AMQP Transport | {{ ReadmeOptionCheckbox "transport.amqp.enabled" }} |
| AMQP Subscriber | {{ ReadmeOptionCheckbox "transport.amqp.subscriber" }} |
| AMQP Publisher | {{ ReadmeOptionCheckbox "transport.amqp.publisher" }} |

## Endpoint middleware

Endpoint middleware is automatically injected by Godin. It is provided by: [go-godin/middleware](github.com/go-godin/middleware)

| **Middleware**      | **Enabled**                                                               |
|--------------|----------------------------------------------------------------------------------|
| InstrumentGRPC |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| Logging |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| RequestID |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)

## Service middleware

Service middleware is use-case specific middleware which the developer has to take care of.
Godin only assits in creating and maintaining service middleware, but will never overwrite middleware.

| **Middleware**      | **Enabled**                                                               |
|--------------|----------------------------------------------------------------------------------|
| Authorization |  {{ ReadmeOptionCheckbox "service.middleware.authorization" }}
| Caching |  {{ ReadmeOptionCheckbox "service.middleware.caching" }}
| Logging |  {{ ReadmeOptionCheckbox "service.middleware.logging" }}
| Recovery |  {{ ReadmeOptionCheckbox "service.middleware.recovery" }}
| Monitoring |  {{ ReadmeOptionCheckbox "service.middleware.monitoring" }}

{{- if .Service.Transport.AMQP }}
## Subscription middleware

Subscription middleware is automatically injected by Godin. It is provided by: [go-godin/middleware/amqp](github.com/go-godin/middleware/amqp)

| **Middleware**      | **Enabled**                                                               |
|--------------|----------------------------------------------------------------------------------|
| RequestID |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| Logging |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| PrometheusInstrumentation |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
{{- end }}