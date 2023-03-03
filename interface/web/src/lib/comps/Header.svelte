<script lang="ts">
	import { APP_NAME } from '$lib/constants';
</script>

<header>
	<div class="container">
		<h1>{APP_NAME}</h1>

		<a
			class="source_cont"
			href="https://github.com/vmkevv/netforenotes"
			target="_blank"
			rel="noreferrer"
		>
			<span>Source</span>
			<svg
				height="30px"
				viewBox="0 0 512 512"
				style="enable-background:new 0 0 512 512"
				width="30px"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
			>
				<path
					d="M256 32A227 227 0 0 0 32 262c0 101 64 187 153 218 11 2 15-5 15-11v-40c-62 14-75-30-75-30-11-27-25-34-25-34-21-14 1-14 1-14 23 2 35 24 35 24 20 35 52 25 65 19 2-15 8-25 14-31-50-6-102-25-102-113 0-26 9-46 23-62-2-6-10-29 2-61 0 0 19-6 62 24a210 210 0 0 1 112 0c43-30 61-24 61-24 13 32 5 55 3 61 14 16 23 37 23 62 0 88-53 107-103 113 8 7 16 21 16 42l-1 63c0 6 4 14 16 11 89-30 153-116 153-217 0-127-100-230-224-230z"
					style="fill-rule:evenodd;clip-rule:evenodd"
				/>
			</svg>
		</a>
	</div>
</header>

<style>
	header {
		padding-top: 8px;
		padding-bottom: 8px;
		box-shadow: 0 0 4px 1px rgba(0, 0, 0, 0.2);
		margin-bottom: 20px;
	}
	.container {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.source_cont {
		display: flex;
		align-items: center;
		font-weight: 600;
		color: inherit;
	}
	.source_cont span {
		margin-right: 4px;
	}
</style>
