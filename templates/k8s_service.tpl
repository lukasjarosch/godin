apiVersion: v1
kind: Service
metadata:
  name: {{ .Service.Name }}
  namespace: {{ .Service.Namespace }}
spec:
  ports:
    - name: grpc
      port: 50051
      protocol: TCP
      targetPort: 50051
    - name: http-debug
      port: 3000
      protocol: TCP
      targetPort: 3000
  selector:
    app: {{ .Service.Name }}
  sessionAffinity: None
  type: ClusterIP
