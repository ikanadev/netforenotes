<script lang="ts">
	import type { PageData } from './$types';
	import { GetNotesOrderEnum, GetNotesSortEnum, type NoteListItem } from '$lib/api';
	import { APP_NAME } from '$lib/constants';
	import AddNoteForm from '$lib/comps/AddNoteForm.svelte';
	import FiltersBar from '$lib/comps/FiltersBar.svelte';
	import Header from '$lib/comps/Header.svelte';
	import './global.css';

	export let data: PageData;

	let sort: GetNotesSortEnum = GetNotesSortEnum.Title;
	let order: GetNotesOrderEnum = GetNotesOrderEnum.Asc;
	let search: string = '';
	let fetchController: AbortController;
	let notesPromise: Promise<NoteListItem[]>;

	function toggleOrder() {
		if (order === GetNotesOrderEnum.Asc) {
			order = GetNotesOrderEnum.Desc;
		} else {
			order = GetNotesOrderEnum.Asc;
		}
	}

	function refetchNotes() {
		notesPromise = data.api.getNotes({ order, sort, search }, { signal: fetchController.signal });
	}

	$: {
		if (fetchController) fetchController.abort();
		fetchController = new AbortController();
		notesPromise = data.api.getNotes({ order, sort, search }, { signal: fetchController.signal });
	}
</script>

<svelte:head>
	<title>{APP_NAME}</title>
</svelte:head>
<Header />
<AddNoteForm api={data.api} {refetchNotes} />
<FiltersBar bind:search bind:sort {toggleOrder} {order} />
<main class="container">
	{#await notesPromise}
		<p>Loading...</p>
	{:then notes}
		<p>Found {notes.length} note(s)</p>
	{/await}
</main>
