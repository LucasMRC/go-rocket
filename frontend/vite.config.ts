import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte';
import tailwindcss from '@tailwindcss/vite';
import { resolve } from 'path';

export default defineConfig({
		plugins: [svelte(), tailwindcss()],
		build: {
			rollupOptions: {
				input: {
					index: resolve(__dirname, 'index.html'),
				}
			}
		}
});
