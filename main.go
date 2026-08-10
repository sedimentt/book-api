package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"method", "path", "status"},
)

var httpErrorsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_errors_total",
		Help: "Total HTTP errors",
	},
)

var dbQueriesTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "db_queries_total",
		Help: "Total database queries",
	},
)

var booksCreatedTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "books_created_total",
		Help: "Total books created",
	},
)

var booksReturnedTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "books_returned_total",
		Help: "Total books returned to clients",
	},
)

var httpRequestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP request duration in seconds",
	},
)

func main() {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpErrorsTotal)
	prometheus.MustRegister(dbQueriesTotal)
	prometheus.MustRegister(booksCreatedTotal)
	prometheus.MustRegister(booksReturnedTotal)
	prometheus.MustRegister(httpRequestDuration)
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/books", booksHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	statusCode := http.StatusOK
	defer func() {
		httpRequestDuration.Observe(time.Since(start).Seconds())
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(statusCode)).Inc()
	}()

	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte("OK")); err != nil {
		statusCode = http.StatusInternalServerError
		httpErrorsTotal.Inc()
		log.Printf("health: write error: %v", err)
	}
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	statusCode := http.StatusOK
	defer func() {
		httpRequestDuration.Observe(time.Since(start).Seconds())
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(statusCode)).Inc()
	}()
	switch r.Method {

	case http.MethodGet:
		rows, err := db.Query("SELECT id, title, author FROM books")
		if err != nil {
			statusCode = http.StatusInternalServerError
			httpErrorsTotal.Inc()
			http.Error(w, err.Error(), statusCode)
			return
		}
		dbQueriesTotal.Inc()
		defer func() {
			if err := rows.Close(); err != nil {
				log.Println(err)
			}
		}()

		var books []Book

		for rows.Next() {
			var b Book
			if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
				statusCode = http.StatusInternalServerError
				httpErrorsTotal.Inc()
				http.Error(w, err.Error(), statusCode)
				return
			}
			books = append(books, b)
		}
		if err := rows.Err(); err != nil {
			statusCode = http.StatusInternalServerError
			httpErrorsTotal.Inc()
			http.Error(w, err.Error(), statusCode)
			return
		}

		if err := json.NewEncoder(w).Encode(books); err != nil {
			statusCode = http.StatusInternalServerError
			httpErrorsTotal.Inc()
			http.Error(w, err.Error(), statusCode)
			return
		}
		booksReturnedTotal.Add(float64(len(books)))

	case http.MethodPost:
		var b Book

		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			statusCode = http.StatusBadRequest
			http.Error(w, err.Error(), statusCode)
			httpErrorsTotal.Inc()
			return
		}

		_, err := db.Exec(
			"INSERT INTO books(title, author) VALUES($1,$2)",
			b.Title,
			b.Author,
		)

		if err != nil {
			statusCode = http.StatusInternalServerError
			http.Error(w, err.Error(), statusCode)
			httpErrorsTotal.Inc()
			return
		}

		booksCreatedTotal.Inc()
		dbQueriesTotal.Inc()
		statusCode = http.StatusCreated
		w.WriteHeader(statusCode)

	default:
		statusCode = http.StatusMethodNotAllowed
		httpErrorsTotal.Inc()
		w.WriteHeader(statusCode)
	}
}
