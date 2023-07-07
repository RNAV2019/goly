<script lang="ts">
	import type { MouseEventHandler } from 'svelte/elements';

	export let onClickHandler: MouseEventHandler<HTMLButtonElement>;
</script>

<button
	class="px-3.5 py-1.5 bg-gray-800 text-white rounded-lg hover:bg-gray-700 hover:text-white dark:bg-white dark:text-gray-800 dark:hover:bg-primary dark:hover:text-white transition-colors text-sm"
	on:click={onClickHandler}
>
	<slot />
</button>
