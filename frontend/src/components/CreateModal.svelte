<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { createFieldsType } from '../utils/helper';
	import TemplateModal from './TemplateModal.svelte';
	import Button from './Button.svelte';

	const dispatch = createEventDispatcher();
	export let show: boolean = true;
	const close = () => {
		show = false;
		redirect = '';
		redirectInputError = '';
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
		} else {
			redirectInputError = 'Enter a URL to shorten';
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
				class="flex-grow p-1 px-2 ml-1 border-2 border-primary bg-transparent outline-none focus:outline-none active:outline-none rounded-lg"
			/>
		</div>
		<h4 class="text-red-400 font-bold text-sm">{redirectInputError}</h4>
		<div class="flex justify-end gap-3">
			<Button onClickHandler={create}>Create</Button>
			<Button onClickHandler={close}>Close</Button>
		</div>
	</TemplateModal>
{/if}
