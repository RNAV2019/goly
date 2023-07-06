<script lang="ts">
	export let title: string;
</script>

<div class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-90 z-50" on:close={close}>
	<div class="mx-auto my-52 max-w-xl bg-white rounded-lg shadow-lg p-8 flex flex-col gap-2">
		<h3 class="text-lg font-medium mb-4">{title}</h3>
		<slot />
	</div>
</div>
