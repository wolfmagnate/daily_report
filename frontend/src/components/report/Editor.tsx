import React, { useState } from 'react';
import { PrivateReflectionEditor, type PrivateReflectionEditorEvents } from './PrivateReflectionEditor';
import { PublicReportEditor, type PublicReportEditorEvents } from './PublicReportEditor';
import { DateNavigator } from './DateNavigator';
import styles from './Editor.module.css';
import type { PrivateReflectionData, PublicReportData } from '../../bff/report/editor';

export interface EditorTabEvents {
    privateReflectionEditorEvents: PrivateReflectionEditorEvents;
    publicReportEditorEvents: PublicReportEditorEvents;
}


export interface EditorProps {
  currentDateStr: string;
  initialTabId?: 'private' | 'public';

  privateReflectionData: PrivateReflectionData;
  publicReportData: PublicReportData;

  events: EditorTabEvents;
}

const TABS = [
  { id: 'private', label: 'プライベート・リフレクション' },
  { id: 'public', label: 'チーム共有日報' },
];

export const Editor: React.FC<EditorProps> = ({
  currentDateStr,
  privateReflectionData,
  publicReportData,
  events: {
    privateReflectionEditorEvents,
    publicReportEditorEvents,
  },
  initialTabId = 'private',
}) => {
  const [activeTab, setActiveTab] = useState(initialTabId);

  return (
    <div className={styles.editorTabsContainer}>
      <div className={styles.header}>
        <div className={styles.dateNavWrapper}>
            <DateNavigator currentDateStr={currentDateStr} />
        </div>
        <div className={styles.tabsWrapper}>
          {TABS.map((tab) => (
            <button
              key={tab.id}
              className={`${styles.tabButton} ${activeTab === tab.id ? styles.active : ''}`}
              onClick={() => setActiveTab(tab.id as 'private' | 'public')}
            >
              {tab.label}
            </button>
          ))}
        </div>
      </div>

      <div className={styles.panelsContainer}>
        <div className={`${styles.panel} ${activeTab === 'private' ? styles.active : ''}`}>
          <PrivateReflectionEditor
            didSection={privateReflectionData.didSection}
            thoughtSection={privateReflectionData.thoughtSection}
            events={privateReflectionEditorEvents}
          />
        </div>
        <div className={`${styles.panel} ${activeTab === 'public' ? styles.active : ''}`}>
          <PublicReportEditor
            markdownContent={publicReportData.markdownContent}
            status={publicReportData.status}
            actions={publicReportData.actions}
            events={publicReportEditorEvents}
          />
        </div>
      </div>
    </div>
  );
};