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

