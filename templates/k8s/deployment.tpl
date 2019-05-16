apiVersion: apps/v1beta2
kind: Deployment
metadata:
  labels:
    app: {{ .Service.Name }}
    component: {{ .Service.Namespace }}
  name: {{ .Service.Name }}
  namespace: {{  .Service.Namespace }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Service.Name }}
      component: {{ .Service.Namespace }}
  strategy: {}
  template:
    metadata:
      labels:
        app: {{ .Service.Name }}
        component: {{ .Service.Namespace }}
    spec:
      containers:
      - image: DOCKER_IMAGE
        name: {{ .Service.Name }}
        ports:
        - containerPort: 50051
        resources: {}
status: {}
