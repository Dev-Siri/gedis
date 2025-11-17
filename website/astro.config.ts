import tailwindcss from "@tailwindcss/vite";
import compress from "astro-compress";
import icon from "astro-icon";
import { defineConfig } from "astro/config";

export default defineConfig({
  integrations: [compress(), icon()],
  vite: {
    plugins: [tailwindcss()],
  },
});
