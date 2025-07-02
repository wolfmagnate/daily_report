import { actions as astroActions } from 'astro:actions';
import React, { useState } from 'react';
import { MilkdownEditor } from './MilkdownEditor';
import commonStyles from './ReportEditorCommon.module.css';
import styles from './PublicReportEditor.module.css';

export interface PublicReportProps {
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
  events: PublicReportEditorEvents;
}

export interface PublicReportEditorEvents {
  onContentChange: (markdown: string) => void;
  onGenerateDraft: () => Promise<void>;
  onSubmit: () => Promise<void>;
}


export const PublicReportEditor: React.FC<PublicReportProps> = ({
  markdownContent,
  status,
  actions,
  events: {
    onContentChange,
    onGenerateDraft,
    onSubmit,
  },
}) => {
  const [isGenerating, setIsGenerating] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleGenerateDraft = async () => {
    if (!actions.generateDraft.isEnabled || isGenerating) return;
    setIsGenerating(true);
    try {
      await astroActions.createReportData({
        didContent: markdownContent,
        thoughtContent: '',
        currentContent: '',
      });
    } catch (error) {
      console.error("Failed to generate draft:", error);
    } finally {
      setIsGenerating(false);
    }
  };
  
  const handleSubmit = async () => {
    if (!actions.submit.isEnabled || isSubmitting) return;
    if (window.confirm("この内容で日報をチームに共有しますか？")) {
      setIsSubmitting(true);
      try {
        await onSubmit();
      } catch (error) {
        console.error("Failed to submit report:", error);
      } finally {
        setIsSubmitting(false);
      }
    }
  };

  return (
    <div className={`${commonStyles.container} ${styles.publicContainer}`}>
      <div className={commonStyles.header}>
        <div>
          <h2 className={commonStyles.mainTitle}>チーム共有日報</h2>
          <p className={commonStyles.mainDescription}>
            プライベートな内省を元に、チームに共有する成果や課題をまとめましょう。
          </p>
        </div>
        <div className={styles.actionsContainer}>
          <button
            onClick={handleGenerateDraft}
            disabled={!actions.generateDraft.isEnabled || isGenerating}
            className={`${commonStyles.button} ${styles.generateButton}`}
          >
            {isGenerating ? '生成中...' : '🤖 AIに下書きを作成させる'}
          </button>
          <button
            onClick={handleSubmit}
            disabled={!actions.submit.isEnabled || isSubmitting}
            className={`${commonStyles.button} ${styles.submitButton}`}
          >
            {isSubmitting ? '提出中...' : '🚀 この内容で提出する'}
          </button>
        </div>
      </div>
      
      <div className={commonStyles.editorWrapper}>
        <MilkdownEditor
          key={`public-editor-${status}-${markdownContent}`}
          initialContent={markdownContent}
          onContentChange={onContentChange}
          isEditable={status !== 'published'}
        />
      </div>
    </div>
  );
};
