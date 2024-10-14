import { IdBody } from '@/types';
import client from '../client';

type AddActBody = {
  Name: string;
  ApplicationId: string;
  ServiceId: string;
};

export const addAct = (body: AddActBody) => {
  return client.post<IdBody[]>('/acts', body);
};

addAct.queryKey = 'addAct';
