package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	moviesv1 "github.com/griggsca91/givemeadamnmovie/gen/go/movies/v1"
	"github.com/griggsca91/givemeadamnmovie/gen/go/movies/v1/moviesv1connect"
	"github.com/griggsca91/givemeadamnmovie/themoviedb"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

type MovieServer struct {
}

// GetMovie implements moviesv1connect.MoviesServiceHandler.
func (*MovieServer) GetMovie(ctx context.Context, req *connect.Request[moviesv1.GetMovieRequest]) (*connect.Response[moviesv1.GetMovieResponse], error) {
	id := req.Msg.MovieId

	service := themoviedb.NewTheMovieDBService("c8a7f2bacaa6395f0d8868fa64597a26")
	details, err := service.GetMovieDetails(id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &connect.Response[moviesv1.GetMovieResponse]{
		Msg: &moviesv1.GetMovieResponse{
			Movie: &moviesv1.Movie{
				Description: details.Description,
				Title:       details.Title,
				ImageUrl:    details.ImageURL,
			},
		},
	}, nil
}

// GetMovieRecommendation implements moviesv1connect.MoviesServiceHandler.
func (*MovieServer) GetMovieRecommendation(context.Context, *connect.Request[moviesv1.GetMovieRecommendationRequest]) (*connect.Response[moviesv1.GetMovieRecommendationResponse], error) {
	fmt.Println("Hit")
	return &connect.Response[moviesv1.GetMovieRecommendationResponse]{
		Msg: &moviesv1.GetMovieRecommendationResponse{
			Movie: &moviesv1.Movie{
				Description: "farts",
				Title:       "farts",
			},
		},
	}, nil
}

func main() {
	movieServer := MovieServer{}
	mux := http.NewServeMux()

	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	loggingInterceptor := connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
				logger.Println("calling:", request.Spec().Procedure)
				logger.Println("request:", request.Any())
				response, err := next(ctx, request)
				if err != nil {
					logger.Println("error:", err)
				} else {
					logger.Println("response:", response.Any())
				}
				return response, err
			})
		},
	)

	path, handler := moviesv1connect.NewMoviesServiceHandler(
		&movieServer,
		connect.WithInterceptors(loggingInterceptor),
	)

	mux.Handle(path, withCORS(handler))
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
