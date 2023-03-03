<script lang="ts">
	import dayjs from 'dayjs';
	import relativeTime from 'dayjs/plugin/relativeTime';
	import localizedFormat from 'dayjs/plugin/localizedFormat';
	import type { PageData } from './$types';
	import { GetNotesOrderEnum, GetNotesSortEnum, type NoteListItem } from '$lib/api';
	import { APP_NAME } from '$lib/constants';
	import AddNoteForm from '$lib/comps/AddNoteForm.svelte';
	import FiltersBar from '$lib/comps/FiltersBar.svelte';
	import Header from '$lib/comps/Header.svelte';
	import Notes from '$lib/comps/Notes.svelte';
	import Message from '$lib/comps/Message.svelte';
	import NoteModal from '$lib/comps/NoteModal.svelte';
	import LoadingBar from '$lib/comps/LoadingBar.svelte';

	dayjs.extend(relativeTime);
	dayjs.extend(localizedFormat);
	import './global.css';

	export let data: PageData;

	let sort: GetNotesSortEnum = GetNotesSortEnum.Title;
	let order: GetNotesOrderEnum = GetNotesOrderEnum.Asc;
	let search: string = '';
	let notes: NoteListItem[] = [];
	let isFetching = false;
	let fetchController: AbortController;
	let selectedNote: NoteListItem | null = null;

	function onSelectNote(note: NoteListItem) {
		selectedNote = note;
	}
	function clearSelectedNote() {
		selectedNote = null;
	}

	function toggleOrder() {
		if (order === GetNotesOrderEnum.Asc) {
			order = GetNotesOrderEnum.Desc;
		} else {
			order = GetNotesOrderEnum.Asc;
		}
	}

	function refetchNotes() {
		isFetching = true;
		data.api
			.getNotes({ order, sort, search }, { signal: fetchController.signal })
			.then((res) => {
				notes = res;
			})
			.finally(() => {
				isFetching = false;
			});
	}

	$: {
		if (fetchController) fetchController.abort();
		fetchController = new AbortController();
		isFetching = true;
		data.api
			.getNotes({ order, sort, search }, { signal: fetchController.signal })
			.then((res) => {
				notes = res;
				isFetching = false;
			})
			.catch((err) => {
				if (err?.cause?.message === 'The user aborted a request.') {
					// if it's an aborted request we did not finish fetching
					return;
				}
				isFetching = false;
			});
	}
</script>

<svelte:head>
	<title>{APP_NAME}</title>
</svelte:head>
<Header />
<NoteModal api={data.api} onClose={clearSelectedNote} note={selectedNote} {refetchNotes} />
<AddNoteForm api={data.api} {refetchNotes} />
<FiltersBar bind:search bind:sort {toggleOrder} {order} />
<main class="container">
	<LoadingBar isLoading={isFetching} />
	{#if notes.length === 0 && !isFetching}
		<Message message="No notes found..." />
	{/if}
	{#if notes.length > 0}
		<Notes {notes} {onSelectNote} />
	{/if}
</main>
