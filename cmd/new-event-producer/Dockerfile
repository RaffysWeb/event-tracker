FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app


COPY ./bin/* ./bin

# Expose the port that your event producer listens on (if applicable)
EXPOSE 8080

# Set the command to run the event producer binary
CMD ["./bin"]
