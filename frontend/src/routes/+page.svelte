<script lang="ts">
	import { PUBLIC_DB } from '$env/static/public';
	import CreateModal from '../components/CreateModal.svelte';
	import List from '../components/List.svelte';
	import type { createFieldsType } from '../utils/helper';
	import { refetchData } from '../utils/store';
	let showModal: boolean = false;

	function handleClose() {
		showModal = false;
	}

	async function createGoly(data: createFieldsType) {
		await fetch(`${PUBLIC_DB}goly`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(data)
		}).then((res) => console.log(res));
		refetchData();
	}
</script>

<main class="flex flex-col gap-6">
	<div class="flex flex-row justify-between">
		<h1 class="text-2xl font-bold">Goly -- The Go URL Shortener!</h1>
		<button
			class="px-4 py-2 bg-green-300 text-green-800 rounded-lg hover:bg-green-400 hover:text-green-900 transition-colors"
			on:click={() => (showModal = true)}>Create</button
		>
	</div>
	<List />
	<CreateModal show={showModal} title={'Create Goly'} on:close={handleClose} {createGoly} />
</main>
