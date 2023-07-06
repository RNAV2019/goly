<script lang="ts">
	import { onMount } from 'svelte';
	import { golies } from '../utils/store';
	import Card from './Card.svelte';
	import type { Goly } from '../utils/helper';
	import { PUBLIC_DB } from '$env/static/public';

	onMount(async () => {
		const res = await fetch(`${PUBLIC_DB}goly`);
		const data = (await res.json()) as Goly[];
		golies.set(data);
	});
</script>

{#each $golies as goly (goly.id)}
	<Card {goly} />
{/each}
