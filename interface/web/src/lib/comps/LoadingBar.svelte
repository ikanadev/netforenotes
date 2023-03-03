<script lang="ts">
	export let isLoading = false;
</script>

<div class="loader_cont">
	{#if isLoading}
		<div class="loader">
			<span class="dot1" />
		</div>
	{/if}
</div>

<style>
	.loader_cont {
		height: 10px;
	}
	.loader {
		position: relative;
		overflow: hidden;
		border-radius: 4px;
		width: 100%;
		height: 4px;
		background: #00000010;
	}
	.loader span {
		position: absolute;
		height: 4px;
		width: 10px;
		border-radius: 4px;
		background: #00000043;
	}
	.dot1 {
		animation: loader_width 2s linear infinite, loader_left 2.2s linear infinite;
	}
</style>
