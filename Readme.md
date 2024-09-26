### Middleware
- The standard pattern for creating a middleware looks like this:
```go
func myMiddleware(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
    // TODO: Execute our middleware logic here...
    next.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}

```
- Simplifying the Middleware : A tweak to this pattern is to use an anonymous function to rewrite the
  myMiddleware middleware like so:
```go
func myMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // TODO: Execute our middleware logic here...
    next.ServeHTTP(w, r)
})
}
```

- Positioning the Middleware

    _If you position your middleware before the servemux in the chain then it
  will act on every request that your application receives._
    ```
  myMiddleware → servemux → application handler
    ```
  
    _you can position the middleware after the servemux in the
  chain — by wrapping a specific application handler. This would cause
  your middleware to only be executed for specific routes. (An example of this would be something like authorization middleware,
  which you may only want to run on specific routes.)_
    ```
  servemux → myMiddleware → application handler
  ```
  
### RESTful Routing
Out of all the third-party routers, The
recommend as a starting point: Pat and Gorilla Mux

### Security Improvements
Make some improvements to
our application so that our data is kept secure during transit and our
server is better able to deal with some common types of Denial-ofService attacks.

- How to quickly and easily create a self-signed TLS certificate, using
only Go.
- The fundamentals of setting up your application so that all requests
and responses are served securely over HTTPS.
- Some sensible tweaks to the default TLS settings to help keep user
information secure and our server performing quickly.
- How to set connection timeouts on our server to mitigate Slowloris
and other slow-client attacks

#### Generating a Self-Signed TLS Certificate

HTTPS is essentially HTTP sent across a TLS (Transport Layer Security)
connection. Because it’s sent over a TLS connection the data is encrypted and
signed, which helps ensure its privacy and integrity during transit.

If you’re not familiar with the term, TLS is essentially the modern version of SSL
(Secure Sockets Layer).

-  create a new 'tls' directory in the root of your project repository to hold the certificate and change into it
- generate a TLS certificate :

  Handily, the crypto/tls package in Go’s standard library includes a
  generate_cert.go tool that we can use to easily create our own self-signed
  certificate.
  ```
  go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
  ```

- It then stores the private key in a key.pem file, and generates a self-signed
  TLS certificate for the host localhost containing the public key — which it
  stores in a cert.pem file
  ```
  tls/
      cert.pem
      key.pem
  ```

## User Authentication/Authorization
- Implement basic signup, login and logout functionality for
users.
- A secure approach to encrypting and storing user passwords securely
in your database using Bcrypt.
- A solid and straightforward approach to verifying that a user is logged
in using middleware and sessions.
- Prevent Cross-Site Request Forgery (CSRF) attacks

