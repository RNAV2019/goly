import { writable } from 'svelte/store';
import type { Goly } from './helper';
import { PUBLIC_DB } from '$env/static/public';

export const golies = writable([] as Goly[]);

export async function refetchData() {
	const data = await fetch(`${PUBLIC_DB}goly`);
	const res = (await data.json()) as Goly[];
	golies.set(res);
}
