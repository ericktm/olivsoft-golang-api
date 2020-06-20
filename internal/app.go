package internal

import (
	"log"
	"syscall"
	"time"

	"github.com/erickoliv/finances-api/account"
	"github.com/erickoliv/finances-api/auth"
	"github.com/erickoliv/finances-api/auth/authhttp"
	"github.com/erickoliv/finances-api/categories/categoryhttp"
	"github.com/erickoliv/finances-api/entries/entryhttp"
	"github.com/erickoliv/finances-api/entries/entrysql"
	"github.com/erickoliv/finances-api/index"
	"github.com/erickoliv/finances-api/internal/db"
	"github.com/erickoliv/finances-api/repository/session"
	"github.com/erickoliv/finances-api/repository/sql"
	"github.com/erickoliv/finances-api/tag"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func buildRouter(conn *gorm.DB) *gin.Engine {

	r := gin.Default()
	r.GET("/", index.Handler)

	r.Use(db.Middleware(conn))

	security := r.Group("/auth")
	authenticator := sql.MakeAuthenticator(conn)
	signer := makeJWTSigner()
	authHandler := authhttp.NewHTTPHandler(authenticator, signer)
	authHandler.Router(security)

	api := r.Group("/api")
	api.Use(authhttp.Middleware(signer))

	accountRepo := sql.MakeAccounts(conn)
	accounts := account.NewHTTPHandler(accountRepo)
	accounts.Router(api)

	tagRepo := sql.BuildTagRepository(conn)
	tags := tag.NewHTTPHandler(tagRepo)
	tags.Router(api)

	categoryRepo := sql.BuildCategoryRepository(conn)
	categories := categoryhttp.NewHandler(categoryRepo)
	categories.Router(api)

	entryRepo := entrysql.BuildRepository(conn)
	entries := entryhttp.NewHandler(entryRepo)
	entries.Router(api)

	return r
}

// TODO: use a configuration service
func makeJWTSigner() auth.SessionSigner {
	appToken, found := syscall.Getenv("APP_TOKEN")
	if !found {
		log.Fatal("env APP_TOKEN not found")
	}

	key := []byte(appToken)
	ttl := time.Hour
	return session.NewJWTSigner(key, ttl)
}

func Run() error {
	conn := db.Prepare()
	defer conn.Close()

	return buildRouter(conn).Run()
}
