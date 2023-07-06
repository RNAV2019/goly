<script lang="ts">
	import { onMount } from 'svelte';
	import type { Goly, updateFieldsType } from '../utils/helper';
	import { refetchData } from '../utils/store';
	import UpdateModal from './UpdateModal.svelte';
	import { PUBLIC_DB, PUBLIC_WS } from '$env/static/public';
	export let goly: Goly;
	let showModal: boolean = false;

	function handleClose() {
		showModal = false;
	}

	async function updateGoly(data: updateFieldsType) {
		await fetch(`${PUBLIC_DB}goly`, {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(data)
		}).then((res) => console.log(res));
		refetchData();
	}

	async function deleteGoly(id: number | undefined) {
		if (id === undefined) {
			return;
		}
		await fetch(`${PUBLIC_DB}goly/${id}`, {
			method: 'DELETE'
		}).then((res) => console.log(res));
		refetchData();
	}

	onMount(() => {
		const socket = new WebSocket(PUBLIC_WS);

		socket.onmessage = async (event) => {
			const message: string = event.data;
			console.log(`Message Data: ${message}`);
			if (message == 'redirect') {
				refetchData();
			}
		};
	});
</script>

<div class="rounded-xl bg-teal-300 text-teal-900 py-12 px-10 font-bold text-lg">
	<p>URL: {PUBLIC_DB}r/{goly.url}</p>
	<p>Redirects to: {goly.redirect}</p>
	<p>Clicked: {goly.clicked}</p>
	<div class="flex flex-row gap-4 mt-3 text-base font-medium">
		<button
			class="px-4 py-2 bg-blue-300 text-blue-800 rounded-xl hover:bg-blue-400 hover:text-blue-900 transition-colors"
			on:click={() => (showModal = true)}>Update</button
		>
		<button
			class="px-4 py-2 bg-red-300 text-red-800 rounded-xl hover:bg-red-400 hover:text-red-900 transition-colors"
			on:click={() => deleteGoly(goly.id)}>Delete</button
		>
	</div>
</div>
<UpdateModal show={showModal} title={'Update Goly'} on:close={handleClose} {updateGoly} {goly} />
