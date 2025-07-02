export interface SessionUser {
  login: string;
  id: string;
  avatar_url: string;
  name: string;
}

const sessionStore = new Map<string, SessionUser>();

export function createSession(user: SessionUser): string {
  const sessionId = crypto.randomUUID();
  sessionStore.set(sessionId, user);
  return sessionId;
}

export function getSession(sessionId: string): SessionUser | undefined {
  return sessionStore.get(sessionId);
}

export function deleteSession(sessionId: string): void {
  sessionStore.delete(sessionId);
}