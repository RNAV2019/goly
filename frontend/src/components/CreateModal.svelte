<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { createFieldsType } from '../utils/helper';
	import TemplateModal from './TemplateModal.svelte';

	const dispatch = createEventDispatcher();
	export let show: boolean = true;
	const close = () => {
		show = false;
		dispatch('close');
	};
	export let createGoly: (data: createFieldsType) => void;
	export let title: string;
	let redirect: string;

	let redirectInputError: string = '';
	const create = () => {
		if (redirect != '') {
			if (redirect.length <= 1) {
				redirectInputError = 'Redirect was too short';
				return;
			}
			let createFields: createFieldsType = {
				redirect: redirect,
				url: ''
			};
			createGoly(createFields);
			close();
		}
	};
</script>

{#if show}
	<TemplateModal {title}>
		<div class="flex flex-row gap-1 items-center">
			<label for="redirect">Redirect To: </label>
			<input
				type="text"
				name="redirect"
				id="redirect"
				bind:value={redirect}
				class="flex-grow p-1 px-2 ml-1 border-2 border-teal-300 outline-none focus:outline-none active:outline-none rounded-lg"
			/>
		</div>
		<h4 class="text-red-400 font-bold text-sm">{redirectInputError}</h4>
		<div class="flex justify-end gap-3">
			<button
				class="px-4 py-2 bg-green-300 text-green-800 rounded-xl hover:bg-green-400 hover:text-green-900 transition-colors"
				on:click={create}>Create</button
			>

			<button
				class="px-4 py-2 bg-red-300 text-red-800 rounded-xl hover:bg-red-400 hover:text-red-900 transition-colors"
				on:click={close}>Close</button
			>
		</div>
	</TemplateModal>
{/if}
