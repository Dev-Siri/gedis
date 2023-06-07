import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";
import compress from "astro-compress";

import mdx from "@astrojs/mdx";

export default defineConfig({
  integrations: [tailwind(), compress(), mdx()],
});
