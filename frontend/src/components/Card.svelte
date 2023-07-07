<script lang="ts">
	import { onMount } from 'svelte';
	import type { Goly, updateFieldsType } from '../utils/helper';
	import { refetchData } from '../utils/store';
	import UpdateModal from './UpdateModal.svelte';
	import { PUBLIC_DB, PUBLIC_WS } from '$env/static/public';
	import { ArrowUpRightSquare, Globe, Hash } from 'lucide-svelte';
	import Button from './Button.svelte';
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

<div
	class="h-52 rounded-lg bg-gray-100 text-black dark:bg-secondary dark:text-white p-5 flex flex-col justify-between w-full"
>
	<div class="flex flex-col gap-4">
		<div class="flex flex-row gap-2 items-center">
			<Globe size={16} class="flex-shrink-0" />
			<a
				href={`${PUBLIC_DB}r/${goly.url}`}
				target="_blank"
				class="text-base text-ellipsis whitespace-nowrap overflow-hidden hover:underline hover:cursor-pointer flex-grow"
			>
				{PUBLIC_DB}r/{goly.url}
			</a>
		</div>
		<div class="flex flex-row gap-2 items-center">
			<ArrowUpRightSquare size={16} class="flex-shrink-0" />
			<a
				href={goly.redirect}
				target="_blank"
				class="text-base text-ellipsis whitespace-nowrap overflow-hidden hover:underline hover:cursor-pointer flex-grow"
				>{goly.redirect}</a
			>
		</div>
		<div class="flex flex-row gap-2 items-center">
			<Hash size={16} class="flex-shrink-0" />
			<h3 class="text-base">{goly.clicked} {goly.clicked === 1 ? 'time' : 'times'}</h3>
		</div>
	</div>
	<div class="flex flex-row gap-3">
		<Button onClickHandler={() => (showModal = true)}>Update</Button>
		<Button onClickHandler={() => deleteGoly(goly.id)}>Delete</Button>
	</div>
</div>
<UpdateModal show={showModal} title={'Update Goly'} on:close={handleClose} {updateGoly} {goly} />
