package protocol

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/mohwa/ci-cd-github-action/api/graphql/generated"
	graph "github.com/mohwa/ci-cd-github-action/api/graphql/resolver"
)

func InitRouterGroupGraphQL(graphqlRouterGroup *gin.RouterGroup) {
	graphqlRouterGroup.Use(GinContextToContextMiddleware())

	graphqlRouterGroup.POST("/query", graphqlHandler())
	// graphql playground(query test)
	graphqlRouterGroup.GET("/", playgroundHandler())
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md#accessing-gincontext
// Resolver 레벨에서, gin.Context 객체에 접근하기위한 미들웨어를 추가한다.
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
