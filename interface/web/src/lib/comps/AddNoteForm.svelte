<script lang="ts">
	import { slide } from 'svelte/transition';
	import type { DefaultApi } from '$lib/api';
	import CloseButton from '$lib/comps/CloseButton.svelte';
	import Message from '$lib/comps/Message.svelte';

	export let api: DefaultApi;
	export let refetchNotes: () => void;

	let showForm = false;
	let title = '';
	let body = '';
	let errMsg = '';
	let isSaving = false;
	$: isFilled = title.trim() !== '' && body.trim() !== '';

	function toggleShowForm() {
		showForm = !showForm;
	}

	function handleSubmit(e: Event) {
		e.preventDefault();
		if (!isFilled) return;
		isSaving = true;
		api
			.addNote({
				noteCreation: {
					title,
					body,
					createdAt: Date.now()
				}
			})
			.then(() => {
				refetchNotes();
				title = '';
				body = '';
				showForm = false;
			})
			.catch((err) => {
				errMsg = err.message;
			})
			.finally(() => {
				isSaving = false;
			});
	}
</script>

<div class="container">
	<div class="form_cont">
		{#if !showForm}
			<p
				in:slide
				out:slide
				role="textbox"
				class="placeholder"
				tabindex={0}
				on:click={toggleShowForm}
				on:keydown={(ev) => {
					if (ev.code === 'Enter') {
						toggleShowForm();
					}
				}}
			>
				New note...
			</p>
		{:else}
			<form in:slide out:slide on:submit={handleSubmit}>
				<div class="title_row">
					<!-- svelte-ignore a11y-autofocus -->
					<input bind:value={title} tabindex={1} placeholder="Título..." autofocus />
					<CloseButton size={24} on:click={toggleShowForm} />
				</div>
				<textarea bind:value={body} tabindex={1} placeholder="Contenido..." rows="3" />
				<Message message={errMsg} variant="error" />
				{#if isFilled}
					<button in:slide out:slide class="submit" type="submit" tabindex={1} disabled={isSaving}>
						{isSaving ? 'Saving...' : 'Save'}
					</button>
				{/if}
			</form>
		{/if}
	</div>
</div>

<style>
	.container {
		margin-bottom: 20px;
	}
	.form_cont {
		transition: all 0.5s;
		width: 100%;
		max-width: 400px;
		margin-left: auto;
		margin-right: auto;
		border-radius: 4px;
		padding-top: 6px;
		padding-bottom: 6px;
		padding-left: 12px;
		padding-right: 12px;
		box-shadow: 0 1px 2px 0 rgba(60, 64, 67, 0.302), 0 2px 6px 2px rgba(60, 64, 67, 0.149);
		border: 1px solid transparent;
	}
	.placeholder {
		font-weight: 600;
		color: #666;
		cursor: text;
		padding: 2px;
	}
	form {
		padding: 3px;
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
	.title_row {
		display: flex;
		align-items: start;
	}
	.title_row input {
		flex: 1;
	}
	form input {
		font-weight: 600;
		font-size: 1.2rem;
		margin-bottom: 12px;
	}
	form textarea {
		font-weight: 500;
		font-size: 1.125rem;
		margin-bottom: 8px;
	}
	.submit {
		align-self: end;
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
	.submit:hover,
	.submit:focus {
		background: rgba(0, 0, 0, 0.0825);
	}
	.submit:disabled {
		opacity: 0.6;
		cursor: not-allowed;
		background: transparent;
	}
</style>
