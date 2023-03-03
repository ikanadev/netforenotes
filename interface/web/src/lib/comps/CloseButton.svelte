<script lang="ts">
	export let size: number = 20;
	export let type: 'button' | 'submit' = 'button';
</script>

<button on:click {type}>
	<svg viewBox="0 0 20 20" fill="currentColor" width={size} height={size}>
		<path
			d="M6.28 5.22a.75.75 0 00-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 101.06 1.06L10 11.06l3.72 3.72a.75.75 0 101.06-1.06L11.06 10l3.72-3.72a.75.75 0 00-1.06-1.06L10 8.94 6.28 5.22z"
		/>
	</svg>
</button>

<style>
	button {
		padding: 2px;
		border-radius: 9999px;
		border: none;
		background: none;
		cursor: pointer;
		transition: background 0.3s;
	}
	button:hover {
		background: rgba(0, 0, 0, 0.0825);
	}
</style>
