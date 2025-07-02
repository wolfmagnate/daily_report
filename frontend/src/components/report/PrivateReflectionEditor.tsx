
import React from 'react';
import { MilkdownEditor } from './MilkdownEditor'; // 既存の汎用コンポーネント
import commonStyles from './ReportEditorCommon.module.css'; // 共通スタイルをインポート
import styles from './PrivateReflectionEditor.module.css';   // 固有スタイルをインポート

// 親コンポーネントから渡されるpropsの型定義
export interface PrivateReflectionProps {
  didSection: {
    markdownContent: string;
  };
  thoughtSection: {
    prompts: { id: string; text: string }[];
  };
  events: PrivateReflectionEditorEvents;
}

export interface PrivateReflectionEditorEvents {
  onDidContentChange: (markdown: string) => void;
  onThoughtContentChange: (markdown: string) => void;
}

export const PrivateReflectionEditor: React.FC<PrivateReflectionProps> = ({
  didSection,
  thoughtSection,
  events: {
    onDidContentChange,
    onThoughtContentChange,
  },
}) => {
  return (
    <div className={`${commonStyles.container} ${styles.privateContainer}`}>
      <div className={commonStyles.header}>
        <div>
          <h2 className={`${commonStyles.mainTitle}`}>プライベート・リフレクション</h2>
          <p className={commonStyles.mainDescription}>
            この内容はあなただけが見ることができます。安心して思考や感情を書き出してください。
          </p>
        </div>
      </div>

      {/* --- やったこと (事実) セクション --- */}
      <section className={commonStyles.section}>
        <h3 className={commonStyles.sectionTitle}>やったこと</h3>
        <MilkdownEditor
          key={`did-editor-${didSection.markdownContent}`}
          initialContent={didSection.markdownContent}
          onContentChange={onDidContentChange}
          isEditable={true}
        />
      </section>

      {/* --- 考えたこと・感じたこと (内省) セクション --- */}
      <section className={commonStyles.section}>
        <h3 className={commonStyles.sectionTitle}>考えたこと・感じたこと</h3>
        
        <div className={styles.promptsContainer}>
          {thoughtSection.prompts.map((prompt) => (
            <div key={prompt.id} className={styles.prompt}>
              <span className={styles.promptIcon}>💡</span>
              <p>{prompt.text}</p>
            </div>
          ))}
        </div>

        <MilkdownEditor
          key={`thought-editor`}
          initialContent=""
          onContentChange={onThoughtContentChange}
          isEditable={true}
        />
      </section>
    </div>
  );
};
