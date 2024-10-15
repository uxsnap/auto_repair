import { IdBody } from '@/types';
import client from '../client';

type AddReceiptBody = {
  Sum: number;
};

export const addReceipt = (body: AddReceiptBody) => {
  return client.post<IdBody[]>('/receipts', body);
};

addReceipt.queryKey = 'addReceipt';
