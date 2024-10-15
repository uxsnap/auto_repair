import { IdBody } from '@/types';
import client from '../client';

type AddContractBody = {
  Name: string;
  Sum: number;
  SignedAt: string;
  Status: string;
};

export const addContract = (body: AddContractBody) => {
  return client.post<IdBody[]>('/contracts', body);
};

addContract.queryKey = 'addContract';
