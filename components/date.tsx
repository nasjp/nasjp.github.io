import { format } from 'date-fns';
import { zonedTimeToUtc } from 'date-fns-tz';
import React from 'react';

type week = '日' | '月' | '火' | '水' | '木' | '金' | '土';

const weeks: Array<week> = ['日', '月', '火', '水', '木', '金', '土'];

const Date = ({ raw }: { raw: string }): JSX.Element => {
  const date = zonedTimeToUtc(raw, 'Asia/Tokyo');
  return (
    <div>{format(date, 'yyyy年MM月dd日(') + weeks[date.getDay()] + ')'}</div>
  );
};

export default Date;
