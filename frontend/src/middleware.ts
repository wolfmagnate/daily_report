import { defineMiddleware } from 'astro:middleware';
import { getSession } from './toplevel/session';

const publicRoutes = [
  '/login',
  '/api/auth/github/redirect',
  '/api/auth/github/callback',
];

export const onRequest = defineMiddleware(async (context, next) => {
  const currentPath = context.url.pathname;

  if (publicRoutes.includes(currentPath)) {
    return next();
  }

  const sessionId = context.cookies.get('session_id')?.value;

  if (!sessionId) {
    return context.redirect('/login');
  }

  const user = getSession(sessionId);

  if (!user) {
    context.cookies.delete('session_id', { path: '/' });
    return context.redirect('/login');
  }

  context.locals.user = user;
  return next();
});