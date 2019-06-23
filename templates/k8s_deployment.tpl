apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: {{ .Service.Name }}
    platform: go
    framework: go-kit
    generator: godin
    version: {{ .Service.Namespace }}-{{ .Service.Name }}-version
  name: {{ .Service.Name }}
  namespace: {{ .Service.Namespace }}
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: {{ .Service.Name }}
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      annotations:
        prometheus.io.scrape: "true"
        prometheus.io.port: "3000"
        prometheus.io.path: /metrics
      labels:
        app: {{ .Service.Name }}
        platform: go
        framework: go-kit
        generator: godin
        version: {{ .Service.Namespace }}-{{ .Service.Name }}-version
    spec:
      containers:
        - env:
            - name: GRPC_PORT
              value: "50051"
            - name: LOG_LEVEL
              value: "info"
          image: {{ .Docker.Registry }}/{{ .Service.Namespace }}-{{ .Service.Name }}:{{ .Service.Namespace }}-{{ .Service.Name }}-version
          imagePullPolicy: IfNotPresent
          name: contact
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities: {}
            privileged: false
            readOnlyRootFilesystem: false
            runAsNonRoot: false
          stdin: true
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          tty: true
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
