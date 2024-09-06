package main

import (
	"bdr/bdsnippetbox/pkg/models/mysql" // New import
	"database/sql"                      // New import
	"flag"
	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/golangcollege/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Add a snippets field to the application struct. This will allow us to
// make the SnippetModel object available to our handlers
// Add a new session field to the application struct.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// THe flag value will be stored in the addr variable at runtime. We’ve used the flag.String() function to define the command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "root:@/snippetbox?parseTime=true", "MySQL database access")
	// Define a new command-line flag for the session secret (a random key whic
	// will be used to encrypt and authenticate session cookies). It should be
	// bytes long.
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret")
	flag.Parse()

	// Create new two loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  // custom log for writing information messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // custom log for writing error messages

	// To keep the main() function tidy I've put the code for creating a connection
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Use the sessions.New() function to initialize a new session manager,
	// passing in the secret key as the parameter. Then we configure it so
	// sessions always expires after 12 hours.
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true // Set the Secure flag on our session cookies
	// Initialize a mysql.SnippetModel instance and add it to the application
	// dependencies.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		session:       session,
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields
	// that the server uses the same network address and routes as before, and
	// the ErrorLog field so that the server now uses the custom errorLog logge
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Call the new app.routes() method
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
