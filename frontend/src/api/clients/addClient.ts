import { IdBody } from '@/types';
import client from '../client';

type AddClientBody = {
  Name: string;
  EmployeeId: string;
  Phone: string;
  HasDocuments: boolean;
  Passport: string;
};

export const addClient = (body: AddClientBody) => {
  return client.post<IdBody[]>('/clients', body);
};

addClient.queryKey = 'addClient';
