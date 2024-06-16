## Create the necessary file structure for Go Web Applicaion
<img src="../pictures/go-web-structure.jpg" width="20%" height="40%">

## The content of the application files are below

### main.go
```go
package main

// Import Packages
import (
    "log"
    "net/http"
)

func main() {
    // Server the Desired HTML File
    http.Handle("/", http.FileServer(http.Dir("./content")))
    log.Fatal(http.ListenAndServe(":9091", nil))
}
```

### index.html
```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Simple Go Web Server</title>
    <link rel="stylesheet" href="css/main.css">
  </head>
  <body>
    <h1>Hello from Go...!!</h2>
    <div>
      <img src="images/go.jpg" alt="Go Lang" class="center">
    </div>
  </body>
</html>
```

### about.html
```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Simple Go Web Server</title>
  </head>
  <body>
    <h1>Kubernetes Deployment Test</h2>
  </body>
</html>
```

### main.css
```css
h1 {
    text-align: center;
    width: 100%;
    font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}
```

#### Now if we run the application locally for testing by using the command ``` go run . ``` then we should be able to see the application by usung the web address ```localhost:9091```
<img src="../pictures/go-output.jpg" width="50%" height="50%">

## Containerizing go web application
#### Now we need to create a ``` Dockerfile ```
### Dockerfile
```yml
# Base Image
FROM golang:1.22-alpine

# Make app directory
RUN mkdir /app

# Copy all content to the app directory
ADD . /app

# Make app directory the working directory
WORKDIR /app

# Download any required modules
RUN go mod download

# Build the program to create an executable binary
RUN go build -o webserver .

# Set the startup command
CMD ["/app/webserver"]
```
#### Now build the docker image with the command ``` docker build -t go-web-server . ```

#### We can test the docker image locallly by using the command ``` docker run -d --publish 9091:9091 --name go-web go-web-server ```

#### This will also give us the same output in the browser as previously.
<img src="../pictures/go-output.jpg" width="50%" height="50%">