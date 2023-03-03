<script lang="ts">
	import { fly } from 'svelte/transition';
	import CloseButton from '$lib/comps/CloseButton.svelte';
	import LoadingBar from '$lib/comps/LoadingBar.svelte';
	import type { DefaultApi, NoteListItem, Note } from '$lib/api';

	export let note: NoteListItem | null;
	export let onClose: () => void;
	export let refetchNotes: () => void;
	export let api: DefaultApi;

	let noteComplete: Note | null = null;
	let title = '';
	let body = '';
	let isFetching = true;

	function setTitle(text: string) {
		title = text;
	}
	function setBody(text: string) {
		body = text;
	}

	function handleUpdate() {
		if (!noteComplete) return;
		api
			.updateNote({
				note: {
					...noteComplete,
					title,
					body
				}
			})
			.then(() => {
				fetchNote();
				refetchNotes();
			});
	}

	function handleDelete() {
		if (!noteComplete) return;
		api
			.deleteNote({
				id: noteComplete.id
			})
			.then(() => {
				noteComplete = null;
				title = '';
				body = '';
				refetchNotes();
				onClose();
			});
	}

	function fetchNote() {
		if (!note) return;
		isFetching = true;
		api
			.getNote({ id: note.id })
			.then((data) => {
				noteComplete = data;
			})
			.finally(() => {
				isFetching = false;
			});
	}

	$: {
		if (note) {
			setTitle(note.title);
			fetchNote();
		}
	}

	$: {
		if (noteComplete) {
			setTitle(noteComplete.title);
			setBody(noteComplete.body);
		}
	}

	$: needSave = noteComplete
		? noteComplete.title.trim() !== title.trim() || noteComplete.body.trim() !== body.trim()
		: false;
	$: isFilled = title.trim() !== '' && body.trim() !== '';
</script>

{#if note != null}
	<div class="background" in:fly out:fly>
		<div class="cont">
			<div class="close_cont">
				<CloseButton on:click={onClose} />
			</div>
			<p class="helper_msg">(Click to edit)</p>
			<form>
				<input placeholder="Title..." bind:value={title} />
				<LoadingBar isLoading={isFetching} />
				<textarea placeholder="Content..." bind:value={body} rows={6} />
				<div class="actions_cont">
					{#if needSave}
						<button class="action_btn" type="button" on:click={handleUpdate} disabled={!isFilled}>
							Save changes
						</button>
					{/if}
					<button class="action_btn delete" type="button" on:click={handleDelete}> Delete </button>
				</div>
			</form>
		</div>
	</div>
{/if}

<style>
	.helper_msg {
		font-size: 0.875rem;
	}
	form {
		display: flex;
		flex-direction: column;
	}
	form input,
	form textarea {
		border: none;
	}
	form input:focus,
	form textarea:focus {
		outline: none;
	}
	form input {
		flex: 1;
		font-weight: 600;
		font-size: 1.2rem;
	}
	form textarea {
		font-weight: 500;
		font-size: 1.125rem;
	}
	.action_btn {
		border: none;
		font-weight: 600;
		border-radius: 6px;
		padding-top: 4px;
		padding-bottom: 4px;
		padding-left: 12px;
		padding-right: 12px;
		background: transparent;
		transition: background 0.3s;
		cursor: pointer;
	}
	.action_btn:hover,
	.action_btn:focus {
		background: rgba(0, 0, 0, 0.06);
	}
	.action_btn:disabled {
		opacity: 0.6;
		cursor: not-allowed;
		background: transparent;
	}
	.action_btn.delete {
		color: rgba(200, 0, 0, 0.8);
	}
	.action_btn.delete:hover,
	.action_btn.delete:focus {
		background: rgba(200, 0, 0, 0.06);
	}
	.close_cont {
		align-self: flex-end;
	}
	.actions_cont {
		display: flex;
		justify-content: flex-end;
		align-items: center;
		margin-top: 8px;
		gap: 8px;
	}
	.background {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: rgba(0, 0, 0, 0.75);
		z-index: 5;
		padding: 8px;
	}
	.cont {
		background: white;
		width: 100%;
		max-width: 400px;
		margin-left: auto;
		margin-right: auto;
		margin-top: 60px;
		border-radius: 4px;
		padding: 12px 16px;
		display: flex;
		flex-direction: column;
	}
</style>
