import { Editor, rootCtx, defaultValueCtx, editorViewOptionsCtx } from '@milkdown/core';
import { Milkdown, MilkdownProvider, useEditor } from '@milkdown/react';
import { commonmark } from '@milkdown/preset-commonmark';
import { nord } from '@milkdown/theme-nord';
import '@milkdown/theme-nord/style.css';
import { listener, listenerCtx } from '@milkdown/plugin-listener';

import styles from './MilkdownEditor.module.css';

export interface MilkdownEditorProps {
  initialContent: string;
  isEditable?: boolean;
  onContentChange?: (markdown: string) => void;
}

const MilkdownInternal: React.FC<MilkdownEditorProps> = ({
  initialContent,
  isEditable,
  onContentChange,
}) => {
  useEditor((root) => {
    return Editor.make()
      .config((ctx) => {
        ctx.set(rootCtx, root);
        ctx.set(defaultValueCtx, initialContent);
        ctx.set(editorViewOptionsCtx, { editable: () => isEditable ?? true });
        if (onContentChange) {
          ctx.get(listenerCtx).markdownUpdated((_, markdown) => {
            onContentChange(markdown);
          });
        }
      })
      .config(nord)
      .use(commonmark)
      .use(listener);
  });
  return <Milkdown />;
};

export const MilkdownEditor: React.FC<MilkdownEditorProps> = (props) => {
  
  return (
    <div className={`${styles.editorContainer}`}>
      <MilkdownProvider>
        <MilkdownInternal {...props} />
      </MilkdownProvider>
    </div>
  );
};