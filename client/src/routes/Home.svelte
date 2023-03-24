<script>
	import {readable, writable} from 'svelte/store'
	import { Timeline, TimelineItem, Button, ButtonGroup, Input } from 'flowbite-svelte';
	import {api} from "../service";
	import {ContentType, StoredContent} from "../rpc/api";
	import AudioRecorder from "../components/AudioRecorder.svelte";
	import {clips} from "../store/content";
	import {onMount} from "svelte";

	let media = []
	let mediaRecorder = null;
	let recording = writable(false);
	let content = writable('');
	const storedContent = writable([]);

	const load = async () => {
		try {
			const results = await api.Search({
				query: ''
			});
			let content = [];
			results.storedContent.forEach((item) => {
				let c = {
					type: item.content.type,
					data: item.content.data,
					normal: item.normalized ? item.normalized.data : undefined,
					createdAt: item.createdAt,
				};

				if (item.content.type === ContentType.AUDIO) {
					console.log(item.content.data);
					const blob = new Blob(item.content.data, { type: "audio/mp3" });
					c.url = window.URL.createObjectURL(blob);
					console.log(c.url)
				}

				content.push(c);
			});
			content = content.reverse();
			storedContent.set(content);
		} catch (e) {
			console.error(e);
		}
	}

	const encoder = new TextEncoder()

	onMount(async () => {
		await load();
	});

	async function saveContent(){
		try {
			const contentID = await api.Save({
				data: encoder.encode($content),
				type: ContentType.TEXT,
				metadata: {},
				createdAt: new Date().toTimeString()
			});
			console.log(contentID);
			await load();
		} catch (e) {
			console.error(e);
		}
	}

	async function saveRecording(){
		try {
			const clip = $clips[0];
			const contentID = await api.Save({
				data: clip.content,
				type: ContentType.AUDIO,
				metadata: {},
				createdAt: new Date().toTimeString()
			});
			console.log(contentID);
			await load();
		} catch (e) {
			console.error(e);
		}
	}
</script>

<section>
	<label for="content-input">Content</label>
	<ButtonGroup class="w-full">
		<Input id="content-input" type="text" bind:value={$content} />
		<Button color="blue" on:click={saveContent}>Save</Button>
	</ButtonGroup>

	<AudioRecorder />
	<Button color="blue" on:click={saveRecording}>Save</Button>

	<Timeline>
		{#each $storedContent as item}
			{#if item.type === ContentType.TEXT}
				<TimelineItem title="text" date="March 2022">
					<p class="text-base font-normal text-gray-500 dark:text-gray-400">
						{new TextDecoder().decode(item.data)}
					</p>
				</TimelineItem>
			{:else if item.type === ContentType.AUDIO}
				<TimelineItem title="Audio" date="March 2022">
					<p class="text-base font-normal text-gray-500 dark:text-gray-400">
						{#if item.normal}
							{item.normal}
						{:else}
							No Normalization
						{/if}
						<audio src={item.url} controls/>
					</p>
				</TimelineItem>
			{/if}
		{/each}
	</Timeline>
</section>
