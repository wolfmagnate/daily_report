
export interface IndexData {
    dailyReports: DailyReportData[];
}

export interface DailyReportData {
    date: string;
    status: 'submitted' | 'pending';
    isToday: boolean;
    url: string;
}

export async function GetIndexData(userID: string): Promise<IndexData> {
    // 本来であれば「今週」を取得して、サーバーに7日間まとめてリクエストする
    // それらをまとめてURLを構築する作業が含まれる

    return {
        dailyReports: [
            {
                date: '6月15日 (日)',
                status: 'submitted',
                isToday: false,
                url: '/reports/2025-06-15',
            },
            {
                date: '6月16日 (月)',
                status: 'pending',
                isToday: false,
                url: '/reports/2025-06-16',
            },
            {
                date: '6月17日 (火)',
                status: 'submitted',
                isToday: true,
                url: '/reports/2025-06-17',
            },
            {
                date: '6月18日 (水)',
                status: 'pending',
                isToday: false,
                url: '/reports/2025-06-18',
            },
            {
                date: '6月19日 (木)',
                status: 'pending',
                isToday: false,
                url: '/reports/2025-06-19',
            },
            {
                date: '6月20日 (金)',
                status: 'pending',
                isToday: false,
                url: '/reports/2025-06-20',
            },
            {
                date: '6月21日 (土)',
                status: 'pending',
                isToday: false,
                url: '/reports/2025-06-21',
            },
        ],
    };
}