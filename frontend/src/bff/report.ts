import { GetActivitySidebarData, type ActivitySidebarData } from "./report/activitySidebar";
import { GetEditorData, type EditorData } from "./report/editor";

export interface ReportData {
    activitySidebar: ActivitySidebarData;
    editor: EditorData;
}

export async function GetReportData(userID: string): Promise<ReportData> {
    const activitySiderbar = await GetActivitySidebarData(userID);
    const editor = await GetEditorData(userID);
    return {
        activitySidebar: activitySiderbar,
        editor: editor,
    };
}