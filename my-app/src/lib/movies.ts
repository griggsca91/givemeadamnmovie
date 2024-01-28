import { readable } from 'svelte/store';

const _movies = [
  {
    "title": "The Shawshank Redemption",
    "category": "Drama",
    "release_date": "1994-09-23",
    "rating": 9.3,
    "genres": ["Crime", "Drama"],
    "imageURL": "https://example.com/shawshank_redemption.jpg",
    "description": "lipsum orem"
  },
  {
    "title": "The Godfather",
    "category": "Crime",
    "release_date": "1972-03-24",
    "rating": 9.2,
    "genres": ["Crime", "Drama"],
    "imageURL": "https://example.com/godfather.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "The Dark Knight",
    "category": "Action",
    "release_date": "2008-07-18",
    "rating": 9.0,
    "genres": ["Action", "Crime", "Drama"],
    "imageURL": "https://example.com/dark_knight.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "Pulp Fiction",
    "category": "Crime",
    "release_date": "1994-10-14",
    "rating": 8.9,
    "genres": ["Crime", "Drama"],
    "imageURL": "https://example.com/pulp_fiction.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "Forrest Gump",
    "category": "Drama",
    "release_date": "1994-07-06",
    "rating": 8.8,
    "genres": ["Drama", "Romance"],
    "imageURL": "https://example.com/forrest_gump.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "Inception",
    "category": "Action",
    "release_date": "2010-07-16",
    "rating": 8.8,
    "genres": ["Action", "Adventure", "Sci-Fi"],
    "imageURL": "https://example.com/inception.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "The Matrix",
    "category": "Action",
    "release_date": "1999-03-31",
    "rating": 8.7,
    "genres": ["Action", "Sci-Fi"],
    "imageURL": "https://example.com/matrix.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "Schindler's List",
    "category": "Biography",
    "release_date": "1993-12-15",
    "rating": 8.9,
    "genres": ["Biography", "Drama", "History"],
    "imageURL": "https://example.com/schindlers_list.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "Fight Club",
    "category": "Drama",
    "release_date": "1999-10-15",
    "rating": 8.8,
    "genres": ["Drama"],
    "imageURL": "https://example.com/fight_club.jpg",
    "description": "lipsum orem"

  },
  {
    "title": "The Lord of the Rings: The Fellowship of the Ring",
    "category": "Action",
    "release_date": "2001-12-19",
    "rating": 8.8,
    "genres": ["Action", "Adventure", "Drama"],
    "imageURL": "https://example.com/lotr_fellowship.jpg",
    "description": "lipsum orem"

  }
]

export type Movie = {
  title: string
  description: string
  imageURL: string
}



export const store = {
  movies: {
    async getRecommendation(): Promise<Movie> {
      const response = await fetch("http://localhost:8080/movies.v1.MoviesService/GetMovieRecommendation", { method: "POST", body: "{}", headers: { "content-type": "application/json" } }).then(r => r.json());
      console.log(response)
      const index = Math.round(Math.random() * _movies.length)

      const movie = _movies[index]


      return Promise.resolve(response.movie as Movie)
    }
  },
}