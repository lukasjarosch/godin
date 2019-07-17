apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: {{ .Service.Name }}
    platform: go
    framework: gokit
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
        framework: gokit
        generator: godin
        version: {{ .Service.Namespace }}-{{ .Service.Name }}-version
    spec:
      containers:
        - env:
            - name: TZ
              value: Europe/Zurich
            - name: GRPC_ADDRESS
              value: "0.0.0.0:50051"
            - name: DEBUG_ADDRESS
              value: "0.0.0.0:3000"
            - name: LOG_LEVEL
              value: "info"
            {{- if .Service.Transport.AMQP }}
            - name: AMQP_ADDRESS
              valueFrom:
                secretKeyRef:
                  key: MESSAGE_BROKER_URL
                  name: rabbitmq
                  optional: false
            {{- end }}
          image: {{ .Docker.Registry }}/{{ .Service.Namespace }}-{{ .Service.Name }}:{{ .Service.Namespace }}-{{ .Service.Name }}-version
          imagePullPolicy: IfNotPresent
          name: {{ .Service.Name }}
          resources:
          	requests:
          	  cpu: 1m
              memory: 15Mi
          securityContext:
            allowPrivilegeEscalation: false
            privileged: false
            capabilities:
              drop:
                - all
            readOnlyRootFilesystem: true
            runAsNonRoot: true
            runAsUser: 65534
          stdin: true
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          tty: true
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
