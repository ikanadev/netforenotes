import type { PageLoad } from "./$types";
import { DefaultApi, Configuration } from "$lib/api";

export const load = (() => {
	const api = new DefaultApi(
		new Configuration({ basePath: import.meta.env.VITE_API_URL }),
	);
	return { api };
}) satisfies PageLoad;
