import { FilterValues, ReceiptsFilters, ReceiptWithData } from '@/types';
import client from '../client';

export const getReceipts = (filters?: FilterValues<ReceiptsFilters>) => {
  return client.get<ReceiptWithData[]>('/receipts', {
    params: filters,
  });
};

getReceipts.queryKey = 'getReceipts';
