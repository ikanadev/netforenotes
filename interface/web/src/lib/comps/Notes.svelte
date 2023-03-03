<script lang="ts">
	import dayjs from 'dayjs';
	import type { NoteListItem } from '$lib/api';

	export let notes: NoteListItem[];
</script>

<div class="cont">
	<p>Found {notes.length} note{notes.length > 1 ? 's' : ''}</p>
	<div class="items_cont">
		{#each notes as note}
			<div class="item_cont">
				<p class="item_title">{note.title}</p>
				<p class="item_date" data-tooltip={dayjs(note.createdAt).format('lll')}>
					{dayjs(note.createdAt).fromNow()}
				</p>
			</div>
		{/each}
	</div>
</div>

<style>
	[data-tooltip] {
		position: relative;
		cursor: help;
	}
	[data-tooltip]:hover::before {
		opacity: 1;
	}
	[data-tooltip]::before {
		content: attr(data-tooltip);
		position: absolute;
		display: block;
		background: rgba(0, 0, 0, 0.75);
		padding: 4px 8px;
		top: -24px;
		color: white;
		border-radius: 3px;
		right: 0;
		z-index: 1;
		opacity: 0;
		pointer-events: none;
		transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
		white-space: nowrap;
	}
	.cont > p {
		margin-bottom: 8px;
	}
	.items_cont {
		display: grid;
		grid-template-columns: repeat(1, 1fr);
		row-gap: 16px;
		column-gap: 16px;
	}
	@media (min-width: 600px) {
		.items_cont {
			grid-template-columns: repeat(2, 1fr);
		}
	}
	@media (min-width: 850px) {
		.items_cont {
			grid-template-columns: repeat(3, 1fr);
		}
	}
	.item_cont {
		display: flex;
		flex-direction: column;
		border-radius: 4px;
		border: 1px solid rgba(0, 0, 0, 0.125);
		box-shadow: 1px 1px 3px 0 rgba(0, 0, 0, 0.125);
		padding: 8px 10px;
		line-height: 1.2;
		cursor: pointer;
		transition: all 0.3s;
	}
	.item_cont:hover {
		box-shadow: 1px 1px 6px 0 rgba(0, 0, 0, 0.25);
	}
	.item_title {
		font-weight: 600;
		font-size: 1.4rem;
		flex: 1;
	}
	.item_date {
		align-self: flex-end;
		font-size: 0.9rem;
		font-style: italic;
		color: #666;
	}
</style>
