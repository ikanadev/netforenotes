<script lang="ts">
	import { slide } from 'svelte/transition';

	export let message = '';
	export let variant: 'success' | 'warning' | 'error' | 'info' | 'neutral' = 'neutral';
</script>

{#if message.length > 0}
	<p
		in:slide
		out:slide
		class:error={variant === 'error'}
		class:warning={variant === 'warning'}
		class:success={variant === 'success'}
		class:info={variant === 'info'}
	>
		{message}
	</p>
{/if}

<style>
	.error {
		color: #dc3d43;
	}
	.warning {
		color: #ffa01c;
	}
	.success {
		color: #299764;
	}
	.info {
		color: #0081f1;
	}
	p {
		color: #7e868c;
		font-weight: 400;
		text-align: center;
		margin-top: 2px;
		margin-bottom: 2px;
	}
</style>
