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

	$: values = $golies.sort((a, b) => {
		return a.id! - b.id!;
	});
</script>

<div class="grid grid-cols-1 gap-4 lg:grid-cols-3 lg:gap-8">
	{#each values as goly (goly.id)}
		<Card {goly} />
	{/each}
</div>
