# Use the official go docker image built on debian.
FROM golang:1.13

# Grab the source code and add it to the workspace.
ADD . src/github.com/jgermita/jgermita.me

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get github.com/revel/modules


RUN revel build github.com/jgermita/jgermita.me app prod
# add db file to file structure
COPY app.db /go/

CMD /go/app/run.sh
# Open up the port where the app is running.
# EXPOSE 8080