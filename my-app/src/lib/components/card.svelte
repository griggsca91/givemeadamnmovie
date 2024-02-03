<script lang="ts">
	type Card = {
		title: string;
		imageUrl: string;
		description: string;
		swiped: boolean;
		activeDrag: boolean;
	};
	export let card: Card;
    export let onRemove: (direction: "left" | "right") => void;

	import { onMount } from 'svelte';

	let div: HTMLElement;
	let isDragging = false;
	let x: number, y: number;
	let initialX: number, initialY: number;

	onMount(() => {
		const handleMouseDown = (event: MouseEvent) => {
			isDragging = true;
			x = event.clientX - div.getBoundingClientRect().left;
			y = event.clientY - div.getBoundingClientRect().top;
			initialX = div.getBoundingClientRect().left;
			initialY = div.getBoundingClientRect().top;
		};

		const handleMouseUp = () => {
			isDragging = false;
			// Reset the position of the div if it hasn't been removed
			if (document.body.contains(div)) {
				div.style.left = `${initialX}px`;
				div.style.top = `${initialY}px`;
			}
		};

		const handleMouseMove = (event: MouseEvent) => {
			if (!isDragging) return;
			div.style.left = `${event.clientX - x}px`;
			div.style.top = `${event.clientY - y}px`;
			let bounds = document.body.clientWidth - document.body.clientWidth / 5;
			if (event.pageX > bounds || event.pageX < document.body.clientWidth - bounds) {
                cleanup();
                onRemove(event.pageX > bounds ? "right" : "left");
				div.remove();
			}
		};

		div.addEventListener('mousedown', handleMouseDown);
		window.addEventListener('mouseup', handleMouseUp);
		window.addEventListener('mousemove', handleMouseMove);

		const cleanup =  () => {
			div.removeEventListener('mousedown', handleMouseDown);
			window.removeEventListener('mouseup', handleMouseUp);
			window.removeEventListener('mousemove', handleMouseMove);
		};

        return cleanup
	});
</script>

<div bind:this={div} style="position: absolute; width: 100px; height: 100px; background: red;">
	<div>
		<img draggable="false" src={card.imageUrl} />
	</div>
	<div>
		<h3>{card.title}</h3>
		<p>{card.description}</p>
	</div>
</div>
