import { useState, useMemo } from 'react';

const formatDate = (date: Date): string => {
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
};

export const useDateNavigator = (initialDateStr: string) => {
  const initialDate = useMemo(() => {
    const date = new Date(initialDateStr);
    return isNaN(date.getTime()) ? new Date() : date;
  }, [initialDateStr]);

  const [currentDate, setCurrentDate] = useState<Date>(initialDate);

  const navigateToDate = (date: Date) => {
    const dateString = formatDate(date);
    window.location.href = `/reports/${dateString}`;
  };

  const goToPreviousDay = () => {
    const newDate = new Date(currentDate);
    newDate.setDate(newDate.getDate() - 1);
    navigateToDate(newDate);
  };

  const goToNextDay = () => {
    const newDate = new Date(currentDate);
    newDate.setDate(newDate.getDate() + 1);
    navigateToDate(newDate);
  };
  
  const selectDate = (date: Date) => {
    navigateToDate(date);
  };

  const formattedDate = useMemo(() => {
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1;
    const day = currentDate.getDate();
    const week = ['日', '月', '火', '水', '木', '金', '土'][currentDate.getDay()];
    return `${year}年${month}月${day}日 (${week})`;
  }, [currentDate]);

  return {
    currentDate,
    formattedDate,
    goToPreviousDay,
    goToNextDay,
    selectDate,
  };
};