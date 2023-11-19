FROM golang:1.19.2-bullseye
 #la imagen de linux con Golang

 #ahora necesito una carpeta dentro del contenedor donde estara la info

 WORKDIR /app

 COPY ./mvc-go/ .
 # el punto seria el workdir

#para instalar los modulos hacemos:
RUN go mod download

RUN go build -o back .



EXPOSE 8090

#8090 es el puerto que comunica la computadora con el docker

CMD ["./back"]