import { format } from 'date-fns';
import { zonedTimeToUtc } from 'date-fns-tz';
import React, { FC } from 'react';

type week = '日' | '月' | '火' | '水' | '木' | '金' | '土';

const weeks: Array<week> = ['日', '月', '火', '水', '木', '金', '土'];

type Params = {
  raw: string;
};

const Date: FC<Params> = ({ raw }): JSX.Element => {
  const date = zonedTimeToUtc(raw, 'Asia/Tokyo');
  return (
    <div>{format(date, 'yyyy年MM月dd日(') + weeks[date.getDay()] + ')'}</div>
  );
};

export default Date;
