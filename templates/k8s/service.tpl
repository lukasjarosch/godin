apiVersion: v1
kind: Service
metadata:
  name: {{ .Service.Name }}
  namespace: {{  .Service.Namespace }}
spec:
  type: ClusterIP
  ports:
    - port: 50051
      targetPort: 50051
      protocol: TCP
      name: grpc-server
  selector:
    app: {{ .Service.Name }}
    component: {{ .Service.Namespace }}
