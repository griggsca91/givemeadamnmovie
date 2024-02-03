export type Movie = {
  title: string
  description: string
  imageUrl: string
}


export const store = {
  movies: {
    async getRecommendation(): Promise<Movie> {
      const response = await fetch("http://localhost:8080/movies.v1.MoviesService/GetMovieRecommendation", { method: "POST", body: "{}", headers: { "content-type": "application/json" } }).then(r => r.json());

      return response.movie
    },

    async getMovie(movieID: string): Promise<Movie> {
      const response = await fetch(
        "http://localhost:8080/movies.v1.MoviesService/GetMovie", 
        { method: "POST", body: JSON.stringify({"movie_id": movieID}), headers: { "content-type": "application/json" } }
        ).then(r => r.json());

      return response.movie
    }

  },
}