<script lang="ts">
// @ts-ignore
import { GetScreenObjects } from "../wailsjs/go/main/App.js"
import { onMount } from "svelte";
import Recorder from "./pages/Recorder.svelte";
import Settings from "./pages/Settings.svelte";
import { WindowGetPosition, WindowSetPosition, type Position } from '../wailsjs/runtime/runtime'

enum StatusCode {
	ProcessSuccess = 0,
	ProcessRetry,
	ProcessCanceled,
	ProcessError,
}

enum Window {
	MAIN = 'main',
	SETTINGS = 'settings',
}


let window_name = Window.MAIN;
let prev_page_position: Position;
let sourceId = "primary";
let inputSources: any[] = [];
let config: any = {
	savePath: "",
	format: "mp4",
	defaultScreen: "1",
};

async function changeWindow(w: Window) {
	if (prev_page_position)
		WindowSetPosition(prev_page_position.x, prev_page_position.y)
	window_name = w;
	prev_page_position = await WindowGetPosition()
}

onMount(async function() {
	const { response } = await GetScreenObjects();

	inputSources = response.map(function(s, i) {
		return {
			id: s.isPrimary ? 'primary' : 'secondary-' + i,
			height: s.height,
			width: s.width,
		};
	})
})

// Log unhandled errors

</script>

<main class="grid h-dvh w-dvw relative">
	{#if window_name === Window.MAIN}
		<Recorder {sourceId} {inputSources} openSettings={() => changeWindow(Window.SETTINGS)} />
	{:else if window_name === Window.SETTINGS}
		<Settings {sourceId} {config} {inputSources} closeSettings={() => changeWindow(Window.MAIN)} />
	{/if}
</main>
