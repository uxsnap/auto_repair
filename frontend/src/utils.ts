export const fromDateToString = (date: Date) => {
  const ndate = new Date(+date);

  ndate.setTime(ndate.getTime() - ndate.getTimezoneOffset() * 60000);

  return `${ndate.toISOString().slice(0, 19)}Z`;
};
