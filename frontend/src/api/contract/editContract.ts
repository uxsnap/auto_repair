import { IdBody } from '@/types';
import client from '../client';

type EditContractBody = {
  Id: string;
  Name: string;
  Sum: number;
  SignedAt: string;
  Status: string;
};

export const editContract = (body: EditContractBody) => {
  return client.patch<IdBody>(`/contracts/${body.Id}`, body);
};

editContract.queryKey = 'editContract';
