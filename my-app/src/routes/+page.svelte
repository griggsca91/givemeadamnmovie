<script lang="ts">
    	import type { Movie } from '$lib/movies';
        import { store } from '$lib/movies';

	let suggestedMovie: Movie;

	async function getRecommendation() {
        suggestedMovie = await store.movies.getRecommendation()
    }
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-200 dark:bg-gray-800">
	<header class="mb-10">
		<h1 class="text-4xl font-bold text-gray-900 dark:text-gray-100">Movie Recommendation</h1>
	</header>
	<main>
		<div class="mb-6">
			<button
				on:click={getRecommendation}
				class="inline-flex items-center justify-center whitespace-nowrap ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-8 py-3 bg-blue-500 text-white rounded-md text-lg font-medium"
			>
				Get Movie Recommendation
			</button>
		</div>
        {#if suggestedMovie}
		<div
			class="rounded-lg border bg-card text-card-foreground shadow-sm max-w-md mx-auto"
			data-v0-t="card"
		>
			<div class="p-6">
				<img
					src={suggestedMovie.imageURL}
					alt="Movie Poster"
					class="w-full h-64 object-cover rounded-md"
					width="300"
					height="400"
					style="aspect-ratio: 300 / 400; object-fit: cover;"
				/>
				<h2 class="mt-4 text-2xl font-bold text-gray-900 dark:text-gray-100">
                    {suggestedMovie.title}
                </h2>
				<p class="mt-2 text-gray-700 dark:text-gray-300">
                    {suggestedMovie.description}
				</p>
			</div>
		</div>
        {/if}
	</main>
</div>
