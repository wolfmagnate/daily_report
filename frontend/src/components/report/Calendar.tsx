import React, { useState } from 'react';
import styles from './Calendar.module.css';

interface CalendarProps {
  selectedDate: Date;
  onDateSelect: (date: Date) => void;
  onClose: () => void;
}

export const Calendar: React.FC<CalendarProps> = ({ selectedDate, onDateSelect, onClose }) => {
  const [displayDate, setDisplayDate] = useState(new Date(selectedDate));

  const changeMonth = (amount: number) => {
    const newDate = new Date(displayDate);
    newDate.setMonth(newDate.getMonth() + amount);
    setDisplayDate(newDate);
  };
  
  const handleDateClick = (day: number) => {
    const newDate = new Date(displayDate.getFullYear(), displayDate.getMonth(), day);
    onDateSelect(newDate);
    onClose();
  };

  const renderCalendarDays = () => {
    const year = displayDate.getFullYear();
    const month = displayDate.getMonth();
    const firstDayOfMonth = new Date(year, month, 1).getDay();
    const daysInMonth = new Date(year, month + 1, 0).getDate();
    
    const days = [];
    for (let i = 0; i < firstDayOfMonth; i++) {
      days.push(<div key={`empty-${i}`} className={styles.dayCell}></div>);
    }

    for (let day = 1; day <= daysInMonth; day++) {
      const isSelected = selectedDate.getFullYear() === year && selectedDate.getMonth() === month && selectedDate.getDate() === day;
      days.push(
        <div 
          key={day} 
          className={`${styles.dayCell} ${styles.day} ${isSelected ? styles.selected : ''}`}
          onClick={() => handleDateClick(day)}
        >
          {day}
        </div>
      );
    }
    return days;
  };

  return (
    <div className={styles.calendarOverlay} onClick={onClose}>
        <div className={styles.calendarContainer} onClick={(e) => e.stopPropagation()}>
            <div className={styles.header}>
                <button onClick={() => changeMonth(-1)} className={styles.navButton}>&lt;</button>
                <div className={styles.monthYear}>
                    {displayDate.getFullYear()}年 {displayDate.getMonth() + 1}月
                </div>
                <button onClick={() => changeMonth(1)} className={styles.navButton}>&gt;</button>
            </div>
            <div className={styles.weekdays}>
                {['日', '月', '火', '水', '木', '金', '土'].map(day => <div key={day} className={styles.weekday}>{day}</div>)}
            </div>
            <div className={styles.daysGrid}>
                {renderCalendarDays()}
            </div>
        </div>
    </div>
  );
};