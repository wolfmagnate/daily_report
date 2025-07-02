/// <reference types="astro/client" />

interface SessionUser {
  login: string;
  id: string;
  avatar_url: string;
  name: string;
}

declare namespace App {
  interface Locals {
    user: SessionUser;
  }
}

interface ImportMetaEnv {
  readonly GITHUB_CLIENT_ID: string;
  readonly GITHUB_CLIENT_SECRET: string;
  readonly ASTRO_SITE: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}