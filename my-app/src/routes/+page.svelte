<script lang="ts">
	import type { Movie } from '$lib/movies';
	import { store } from '$lib/movies';
	import Card from '$lib/components/card.svelte';

	import { spring } from 'svelte/motion';

	let suggestedMovie: Movie;

	let movieID: string;

	type Card = {
		title: string;
		imageUrl: string;
		description: string;
		swiped: boolean;
		activeDrag: boolean;
	};
	let coords = spring(
		{ x: 50, y: 50 },
		{
			stiffness: 0.1,
			damping: 0.25
		}
	);

	$: cards = [
		{
			title: 'Card 1',
			imageUrl: 'https://picsum.photos/id/237/300/400',
			description: 'This is a card',
			swiped: false,
			activeDrag: false
		},
		{
			title: 'Card 2',
			imageUrl: 'https://picsum.photos/id/238/300/400',
			description: 'This is a card',
			swiped: false,
			activeDrag: false
		},
		{
			title: 'Card 3',
			imageUrl: 'https://picsum.photos/id/239/300/400',
			description: 'This is a card',
			swiped: false,
			activeDrag: false
		},
		{
			title: 'Card 4',
			imageUrl: 'https://picsum.photos/id/240/300/400',
			description: 'This is a card',
			swiped: false,
			activeDrag: false
		}
	];
	$: displayedCards = cards.filter((card) => !card.swiped);

	async function getMovie() {
		const movie = await store.movies.getMovie(movieID);
		suggestedMovie = movie;
	}

	async function getRecommendation() {
		suggestedMovie = await store.movies.getRecommendation();
	}
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-200 dark:bg-gray-800">
	<header class="mb-10">
		<h1 class="text-4xl font-bold text-gray-900 dark:text-gray-100">Movie Recommendation</h1>
	</header>
	{#each displayedCards as card, index}
		<Card {card} onRemove={console.log} />
	{/each}
	<main>
		<div class="mb-6">
			<p class="text-gray-700 dark:text-gray-300">Get the movie by id</p>
			<input
				bind:value={movieID}
				type="text"
				class="mt-2 w-full px-3 py-2 border border-gray-300 dark:border-gray-700 rounded-md shadow-sm focus:ring-primary focus:border-primary"
			/>

			<button
				on:click={getMovie}
				class="inline-flex items-center justify-center whitespace-nowrap ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-8 py-3 bg-blue-500 text-white rounded-md text-lg font-medium"
			>
				Get Movie
			</button>
		</div>
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
						src={suggestedMovie.imageUrl}
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
