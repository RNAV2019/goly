<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { Goly, updateFieldsType } from '../utils/helper';
	import TemplateModal from './TemplateModal.svelte';
	import Button from './Button.svelte';
	const dispatch = createEventDispatcher();
	export let show: boolean = true;
	const close = () => {
		show = false;
		dispatch('close');
	};
	export let updateGoly: (data: updateFieldsType) => void;
	export let goly: Goly;
	export let title: string;

	let urlInputError: string = '';
	let redirectInputError: string = '';
	const update = () => {
		if (goly.redirect != undefined && goly.url != undefined && goly.id != undefined) {
			if (goly.redirect.length <= 1) {
				redirectInputError = 'Redirect was too short';
				return;
			} else if (goly.url.length < 8) {
				urlInputError = 'URL was too short (has to be 8 characters long)';
				return;
			}
			let updateFields: updateFieldsType = {
				redirect: goly.redirect,
				url: goly.url,
				id: goly.id
			};
			updateGoly(updateFields);
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
				bind:value={goly.redirect}
				class="flex-grow p-1 px-2 ml-1 border-2 border-primary bg-transparent outline-none focus:outline-none active:outline-none rounded-lg"
			/>
		</div>
		<h4 class="text-red-400 font-bold text-sm">{redirectInputError}</h4>
		<div class="flex flex-row gap-1 items-center mb-2">
			<label for="url">URL: </label>
			<input
				type="text"
				name="url"
				id="url"
				bind:value={goly.url}
				class="flex-grow p-1 px-2 ml-1 border-2 border-primary bg-transparent outline-none focus:outline-none active:outline-none rounded-lg"
				maxlength={8}
			/>
		</div>
		<h4 class="text-red-400 font-bold text-sm">{urlInputError}</h4>
		<div class="flex justify-end gap-3">
			<Button onClickHandler={update}>Update</Button>
			<Button onClickHandler={close}>Close</Button>
		</div>
	</TemplateModal>
{/if}
