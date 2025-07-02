import type { APIRoute } from 'astro';

export const GET: APIRoute = async ({ redirect, cookies }) => {
  const state = crypto.randomUUID();

  const githubAuthUrl = new URL('https://github.com/login/oauth/authorize');

  githubAuthUrl.searchParams.set('client_id', import.meta.env.GITHUB_CLIENT_ID);
  console.log(import.meta.env.ASTRO_SITE)
  githubAuthUrl.searchParams.set('redirect_uri', new URL('/api/auth/github/callback', import.meta.env.ASTRO_SITE).toString());
  githubAuthUrl.searchParams.set('state', state);
  githubAuthUrl.searchParams.set('scope', 'read:user repo'); 

  cookies.set('github_oauth_state', state, {
    path: '/',
    maxAge: 60 * 10,
    httpOnly: true,
    secure: import.meta.env.PROD,
    sameSite: 'lax',
  });

  return redirect(githubAuthUrl.toString());
};