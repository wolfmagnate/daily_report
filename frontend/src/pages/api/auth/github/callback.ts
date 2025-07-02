import type { APIRoute } from 'astro';
import { createSession, type SessionUser } from '../../../../toplevel/session';

export const GET: APIRoute = async ({ url, cookies, redirect }) => {
  const code = url.searchParams.get('code');
  const state = url.searchParams.get('state');

  const savedState = cookies.get('github_oauth_state')?.value;
  if (!state || !savedState || state !== savedState) {
    return new Response('Invalid state parameter', { status: 400 });
  }
  cookies.delete('github_oauth_state', { path: '/' });

  if (!code) {
    return new Response('Error: No code provided by GitHub.', { status: 400 });
  }

  try {
    const tokenResponse = await fetch('https://github.com/login/oauth/access_token', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify({
        client_id: import.meta.env.GITHUB_CLIENT_ID,
        client_secret: import.meta.env.GITHUB_CLIENT_SECRET,
        code: code,
      }),
    });

    const tokenData = await tokenResponse.json();

    if (tokenData.error) {
      console.error('GitHub token error:', tokenData.error_description);
      return new Response('Failed to retrieve access token.', { status: 500 });
    }

    const accessToken = tokenData.access_token;

    const userResponse = await fetch('https://api.github.com/user', {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        'User-Agent': 'Astro-GitHub-Login-App',
      },
    });

    const userData = await userResponse.json();
    
    const user: SessionUser = {
      login: userData.login,
      id: userData.id,
      avatar_url: userData.avatar_url,
      name: userData.name,
    };

    const sessionId = createSession(user);

    cookies.set('session_id', sessionId, {
      path: '/',
      httpOnly: true,
      secure: import.meta.env.PROD,
      maxAge: 60 * 60 * 24 * 7,
      sameSite: 'lax',
    });

    return redirect('/');

  } catch (error) {
    console.error('Authentication callback error:', error);
    return new Response('An unexpected error occurred during authentication.', { status: 500 });
  }
};