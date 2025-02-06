package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/XSAM/otelsql"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ricardominze/gobootstrap/api"
	"github.com/ricardominze/gobootstrap/core/domain/account"
	"github.com/ricardominze/gobootstrap/core/domain/customer"
	"github.com/ricardominze/gobootstrap/infra/adapter"
	"github.com/ricardominze/gobootstrap/infra/middleware"
	"github.com/ricardominze/gobootstrap/infra/telemetry"

	// _ "github.com/go-sql-driver/mysql" //Driver MySQL
	// _ "github.com/godror/godror"       //Driver Oracle
	// _ "github.com/mattn/go-sqlite3" //Driver SQLite
	_ "github.com/lib/pq" //Driver PostgreSQL
)

// Prometheus Global Vars

var (
	HttpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of Requests",
	}, []string{"path"})
)

//Lembrete: A função init() é executada automaticamente pelo runtime do Go antes da execução da função main(), sem necessidade de ser chamada explicitamente.

func init() {

	// Prometheus Register Vars

	prometheus.MustRegister(HttpRequests)

}

func main() {

	// Context

	ctx := context.Background()

	// Environment

	godotenv.Load(".env")

	// Registro de Tracer no Banco de Dados.
	driverName, err := otelsql.Register(os.Getenv("DATABASE"), otelsql.WithAttributes())
	if err != nil {
		fmt.Println("Failed to register otel driver: %v", err)
		return
	}

	// Database
	db, err := sql.Open(driverName, os.Getenv("DBSTRING"))

	if err != nil {
		fmt.Println("Database Error:", err)
		return
	}
	defer db.Close()

	//Roda as Migrations
	if err := goose.Up(db, os.Getenv("MIGRATION_DIR")); err != nil {
		log.Fatalf("Falha ao rodar migrations: %v", err)
	}

	// Jaeger Trace

	cleanup, ctx := telemetry.InitTracer(ctx, "GOBOOTSTRAP", os.Getenv("JAEGER"))
	defer cleanup()

	// Router

	router := adapter.NewRouterClassic() //A implementação utiliza o ServerMux
	// router := adapter.NewRouterGorilla() //A implementação utiliza o GorillaMux

	// Middleware de Metrica

	router.Use(func(next http.Handler) http.Handler {
		return middleware.TelemetryMetricMiddleware(next, HttpRequests)
	})

	// Middleware de Trace

	router.Use(func(next http.Handler) http.Handler {
		return middleware.TelemetryTraceMiddleware(next, ctx)
	})

	// Rotas

	router.Handle("/metrics", promhttp.Handler())
	api.NewCustomerController(customer.NewCustomerDependenciesInjection(db), account.NewAccountDependenciesInjection(db)).MakeHandlers(router)
	api.NewAccountController(account.NewAccountDependenciesInjection(db)).MakeHandlers(router)

	// Run

	fmt.Printf("Server is listening on port %s\n", os.Getenv("SERVICE_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), router)

	if err != nil {
		fmt.Printf("ERR: %s", err.Error())
	}
}
