// @ts-check
import { defineConfig } from 'astro/config';

import react from '@astrojs/react';

import netlify from '@astrojs/netlify';

const isProd = process.env.NODE_ENV === 'production';

export default defineConfig({
  // SSRを有効化
  output: 'server',

  site: isProd ? 'https://your-production-domain.com' : 'http://localhost:4321',
  integrations: [react()],

  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          api: "modern-compiler",
        },
      },
    },
  },

  adapter: netlify()
});