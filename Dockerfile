FROM golang:1.20-alpine AS build-stage
WORKDIR /project
ENV GOPATH="${GOPATH}:/project"
COPY . .
RUN apk add build-base
RUN go install .

RUN go test -short -coverprofile=cov.out
RUN go build -tags musl -o runnable .

#Sonar scan
FROM  sonarsource/sonar-scanner-cli AS sonar-stage
WORKDIR /project
COPY --from=build-stage /project /project
ARG SONARQUBE_TOKEN
ARG SONARQUBE_URL
ARG SONARQUBE_ENV
RUN if [ $SONARQUBE_ENV = "prod" ] ; then sonar-scanner -Dsonar.host.url=$SONARQUBE_URL -Dsonar.login=$SONARQUBE_TOKEN ; else echo "$SONARQUBE_ENV no envia a sonar"; fi


# Add final docker image
FROM golang:1.20-alpine
COPY --from=build-stage /project/runnable /project/runnable

# Expose correct port
EXPOSE 8080

# command to run
ENTRYPOINT ["/project/runnable"]