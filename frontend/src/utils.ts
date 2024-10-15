export const fromDateToString = (date: Date) => {
  console.log(date);

  date = new Date(+date);
  date.setTime(date.getTime() - date.getTimezoneOffset() * 60000);
  let dateAsString = date.toISOString().slice(0, 19);
  return dateAsString + 'Z';
};
