// components/report/DateNavigator.tsx
import React, { useState } from 'react';
import { useDateNavigator } from './useDateNavigator';
import { Calendar } from './Calendar';
import styles from './DateNavigator.module.css';

interface DateNavigatorProps {
  currentDateStr: string;
}

export const DateNavigator: React.FC<DateNavigatorProps> = ({ currentDateStr }) => {
  const { currentDate, formattedDate, goToPreviousDay, goToNextDay, selectDate } = useDateNavigator(currentDateStr);
  const [isCalendarOpen, setIsCalendarOpen] = useState(false);

  return (
    <div className={styles.dateNavContainer}>
      <button onClick={goToPreviousDay} className={styles.arrowButton} aria-label="前の日へ">
        &lt;
      </button>
      <div className={styles.dateDisplay} onClick={() => setIsCalendarOpen(true)}>
        {formattedDate}
      </div>
      <button onClick={goToNextDay} className={styles.arrowButton} aria-label="次の日へ">
        &gt;
      </button>

      {isCalendarOpen && (
        <Calendar 
          selectedDate={currentDate} 
          onDateSelect={selectDate}
          onClose={() => setIsCalendarOpen(false)} 
        />
      )}
    </div>
  );
};