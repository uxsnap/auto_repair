import { IdBody } from '@/types';
import client from '../client';

type AddDetailBody = {
  Name: string;
  Price: number;
  Type: string;
};

export const addDetail = (body: AddDetailBody) => {
  return client.post<IdBody[]>('/details', body);
};

addDetail.queryKey = 'addDetail';
