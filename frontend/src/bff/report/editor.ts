
export interface EditorData {
  currentDate: string;
  privateReflection: PrivateReflectionData;
  publicReport: PublicReportData;
}

export interface PrivateReflectionData {
    didSection: {
        markdownContent: string;
    };
    thoughtSection: {
        prompts: { id: string; text: string }[];
    };
}

export interface PublicReportData {
    markdownContent: string;
    status: 'draft' | 'published';
    actions: {
        generateDraft: {
            isEnabled: boolean;
        };
        submit: {
            isEnabled: boolean;
        };
    };
}

export async function GetEditorData(userID: string): Promise<EditorData> {
    return {
        currentDate: '2023-10-27',
        privateReflection: {
            didSection: {
                markdownContent: '## 今日の作業\n\n- ユーザー認証の実装\n- 日報UIの設計\n- チームのコードレビュー',
            },
            thoughtSection: {
                prompts: [
                    { id: 'p1', text: '今日直面した課題は何ですか？' },
                    { id: 'p2', text: '今日学んだことは何ですか？' },
                ],
            },
        },
        publicReport: {
            markdownContent: '## 日報 - 2023-10-27\n\n本日は、アプリケーションのユーザー認証機能の実装に注力し、ユーザーの安全なアクセスを確保しました。また、日報のユーザーインターフェース設計にも貢献し、直感的で効率的な体験を目指しました。さらに、コードレビューに参加し、フィードバックを提供してチーム内のコード品質を確保しました。\n\n**主な成果:**\n- ユーザー認証の実装完了\n- 日報UIデザインの完了\n- コードレビューへの貢献',
            status: 'draft',
            actions: {
                generateDraft: {
                    isEnabled: true,
                },
                submit: {
                    isEnabled: true,
                },
            },
        },
    };
}
