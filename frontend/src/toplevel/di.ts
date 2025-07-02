import type { ReportState } from "../pages/reports/[date].astro";

function DI() : ReportState {
  return {
    editorTabEvents: {
      privateReflectionEditorEvents: {
        onDidContentChange: (markdown: string) => {},
        onThoughtContentChange: (markdown: string) => {},
      },
      publicReportEditorEvents: {
          onContentChange: (markdown: string) => {},
          onGenerateDraft: async () => {},
          onSubmit: async () => {},
      },
    }
  }
}

export const reportState = DI();