FROM golang:1.19 AS build
WORKDIR /app
COPY ./calc .
RUN CGO_ENABLED=0 go build -a main.go


FROM scratch
COPY --from=build /app /
EXPOSE 8080
CMD [ "/main" ]    
