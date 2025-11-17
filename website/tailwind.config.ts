import type { Config } from "tailwindcss";

export default {
  content: ["./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}"],
  theme: {
    extend: {
      boxShadow: {
        dark: "0px 2px 122px 18px rgba(0, 0, 0, 0.23)",
        light: "-22px 10px 300px -97px hsl(0, 0%, 100%)",
      },
      colors: {
        "dark-gray": "rgb(17, 17, 17)",
        "semi-dark-gray": "#1c1c1c",
      },
    },
  },
  future: {
    hoverOnlyWhenSupported: true,
  },
} satisfies Config;
