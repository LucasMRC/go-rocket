<script lang="ts">
// @ts-ignore
import { StartRecording, StopRecording } from '../../wailsjs/go/main/App.js'
import { WindowSetSize } from '../wailsjs/runtime/runtime'
import { EventsOn } from '../wailsjs/runtime/runtime';
import { createInterval, deleteInterval, formatTime } from '../utils';
import RecordButton from '../components/RecordButton.svelte';
import { onMount } from 'svelte';

let recording = false;
let timeRecording = 0;
let panelOpened = false;
let processReady = true;

$: formattedTime = formatTime(timeRecording);

onMount(() => {
	WindowSetSize(130, 48)
})

async function startRecording(): Promise<void> {
	StartRecording();
	EventsOn('recording-started', function() {
		console.info('Renderer process: Recording started');
		recording = true;
		panelOpened = true;
		createInterval(() => {
			++timeRecording;
		}, 1000);
	})
}

async function stopRecording() {
	processReady = false;
	panelOpened = false;
	timeRecording = 0;
	recording = false;
	deleteInterval();
	await StopRecording()
		.then(() => processReady = true);
}
</script>

<section class="my-auto ml-8 flex items-center select-none">
	<RecordButton
		start={startRecording}
		stop={stopRecording}
		disabled={!processReady}
		recording={recording}
	/>
	<div class={`flex border border-slate-500 gap-1 transition-all items-center ${panelOpened ? 'w-[25%]' : 'w-[5%]'} h-8 rounded-md pr-1 py-[0.2rem] pl-2 bg-gradient-to-r from-slate-900 to-slate-950 relative left-8 z-10 box-border ${!processReady ? 'justify-center' : ''}`}>
		{#if processReady && recording}
				<p class={`text-xs text-indigo-300 transition-opacity absolute right-2 ${panelOpened ? 'opacity-100' : 'opacity-0'}`}>{formattedTime}</p>
		{/if}
	</div>
</section>
