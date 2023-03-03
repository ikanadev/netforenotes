<script lang="ts">
	import { GetNotesOrderEnum, GetNotesSortEnum } from '$lib/api';

	export let search: string;
	export let sort: GetNotesSortEnum;
	export let toggleOrder: () => void;
	export let order: GetNotesOrderEnum;
</script>

<div class="container">
	<div class="filters_cont">
		<div class="input_cont">
			<svg viewBox="0 0 20 20" fill="currentColor" width={18} height={18}>
				<path
					fill-rule="evenodd"
					d="M9 3.5a5.5 5.5 0 100 11 5.5 5.5 0 000-11zM2 9a7 7 0 1112.452 4.391l3.328 3.329a.75.75 0 11-1.06 1.06l-3.329-3.328A7 7 0 012 9z"
					clip-rule="evenodd"
				/>
			</svg>
			<input bind:value={search} placeholder="Search by title or content..." />
		</div>
		<div class="flex_1" />
		<div class="filters_right_cont">
			<p>Order by:</p>
			<select bind:value={sort}>
				<option value={GetNotesSortEnum.Title}>Title</option>
				<option value={GetNotesSortEnum.Date}>Date</option>
			</select>
			<button on:click={toggleOrder}>
				{#if order === GetNotesOrderEnum.Asc}
					<svg viewBox="0 0 20 20" fill="currentColor" width={20} height={20}>
						<path
							fill-rule="evenodd"
							d="M2 3.75A.75.75 0 012.75 3h11.5a.75.75 0 010 1.5H2.75A.75.75 0 012 3.75zM2 7.5a.75.75 0 01.75-.75h7.508a.75.75 0 010 1.5H2.75A.75.75 0 012 7.5zM14 7a.75.75 0 01.75.75v6.59l1.95-2.1a.75.75 0 111.1 1.02l-3.25 3.5a.75.75 0 01-1.1 0l-3.25-3.5a.75.75 0 111.1-1.02l1.95 2.1V7.75A.75.75 0 0114 7zM2 11.25a.75.75 0 01.75-.75h4.562a.75.75 0 010 1.5H2.75a.75.75 0 01-.75-.75z"
							clip-rule="evenodd"
						/>
					</svg>
				{/if}
				{#if order === GetNotesOrderEnum.Desc}
					<svg viewBox="0 0 20 20" fill="currentColor" width={20} height={20}>
						<path
							fill-rule="evenodd"
							d="M2 3.75A.75.75 0 012.75 3h11.5a.75.75 0 010 1.5H2.75A.75.75 0 012 3.75zM2 7.5a.75.75 0 01.75-.75h6.365a.75.75 0 010 1.5H2.75A.75.75 0 012 7.5zM14 7a.75.75 0 01.55.24l3.25 3.5a.75.75 0 11-1.1 1.02l-1.95-2.1v6.59a.75.75 0 01-1.5 0V9.66l-1.95 2.1a.75.75 0 11-1.1-1.02l3.25-3.5A.75.75 0 0114 7zM2 11.25a.75.75 0 01.75-.75H7A.75.75 0 017 12H2.75a.75.75 0 01-.75-.75z"
							clip-rule="evenodd"
						/>
					</svg>
				{/if}
			</button>
		</div>
	</div>
</div>

<style>
	.filters_cont {
		display: flex;
		flex-wrap: wrap;
		align-items: center;
		justify-content: end;
		margin-bottom: 12px;
	}
	.flex_1 {
		flex: 1;
	}
	.input_cont,
	.filters_right_cont {
		display: flex;
		align-items: center;
		gap: 10px;
		margin-bottom: 10px;
	}
	input {
		width: 260px;
		border: none;
		font-weight: 500;
		border-bottom: 2px solid rgba(0, 0, 0, 0.1);
		padding-left: 6px;
	}
	input:focus {
		outline: none;
		border-bottom: 2px solid rgba(0, 0, 0, 0.5);
	}
	.filters_cont p {
		font-weight: 500;
	}
	select {
		padding-left: 12px;
		padding-right: 20px;
		border-radius: 6px;
		padding-top: 4px;
		padding-bottom: 4px;
		border: none;
		font-weight: 600;
		background-color: transparent;
		cursor: pointer;
		transition: background 0.3s;
	}
	select:hover,
	select:focus {
		background: rgba(0, 0, 0, 0.06);
	}
	button {
		align-self: end;
		border: none;
		font-weight: 600;
		border-radius: 6px;
		padding-top: 4px;
		padding-bottom: 4px;
		padding-left: 8px;
		padding-right: 8px;
		background: transparent;
		transition: background 0.3s;
		cursor: pointer;
	}
	button:hover,
	button:focus {
		background: rgba(0, 0, 0, 0.06);
	}
</style>
