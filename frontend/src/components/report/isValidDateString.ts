
export function isValidDateString(dateString: string): boolean {
  if (!dateString) return false;
  const regex = /^\d{4}-\d{2}-\d{2}$/;
  if (!regex.test(dateString)) {
    return false;
  }
  const d  = new Date(dateString);
  if (isNaN(d.getTime())) {
    return false;
  }

  return d.toISOString().startsWith(dateString);
}
