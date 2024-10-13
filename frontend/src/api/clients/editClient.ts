import { IdBody } from '@/types';
import client from '../client';

type EditClientBody = {
  Id: string;
  Name: string;
  EmployeeId: string;
  Phone: string;
  HasDocuments: boolean;
  Passport: string;
};

export const editClient = (body: EditClientBody) => {
  return client.patch<IdBody>(`/clients/${body.Id}`, body);
};

editClient.queryKey = 'editClient';
