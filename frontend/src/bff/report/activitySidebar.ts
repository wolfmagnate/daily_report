
export interface Activity {
  source: 'GitHub' | 'Trello' | 'Slack';
  summary: string;
  details: string;
  sourceUrl: string;
  timestamp: string;
}

export interface TrelloCard {
  title: string;
  url: string;
}

export interface ActivityGroupData {
  trelloCard: TrelloCard;
  activities: Activity[];
}

export interface ActivitySidebarData {
  groups: ActivityGroupData[];
}


export async function GetActivitySidebarData(userID: string): Promise<ActivitySidebarData> {
    // 本当だったら、userIDから関連するtrello一覧と収集したアクティビティ一覧を持ってくる
    // activityテーブルの外部キーなので上手く関連付けたり、一部のidを捨てたりする作業がある

    return {
        groups: [
            {
                trelloCard: {
                    title: 'ユーザー認証の実装',
                    url: 'https://trello.com/c/1',
                },
                activities: [
                    {
                        source: 'GitHub',
                        summary: 'mainブランチに3コミットをプッシュ',
                        details: 'ログインページの追加、ミドルウェアのセットアップ、セッションバグの修正',
                        sourceUrl: 'https://github.com/user/repo/commit/abc',
                        timestamp: '2023-10-27T09:00:00Z',
                    },
                    {
                        source: 'Slack',
                        summary: 'チームと認証フローについて議論',
                        details: '当面はGitHub OAuthを利用することで合意',
                        sourceUrl: 'https://slack.com/archives/C123/p12345',
                        timestamp: '2023-10-27T10:30:00Z',
                    },
                ],
            },
            {
                trelloCard: {
                    title: '日報UIの設計',
                    url: 'https://trello.com/c/2',
                },
                activities: [
                    {
                        source: 'Trello',
                        summary: 'カードにモックアップを追加',
                        details: 'レビュー用にFigmaのリンクを添付',
                        sourceUrl: 'https://trello.com/c/2#comment-1',
                        timestamp: '2023-10-27T11:45:00Z',
                    },
                ],
            },
        ],
    }
}