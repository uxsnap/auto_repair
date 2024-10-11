export type IdBody = {
  id: string;
};

export type Application = {
  id: string;
  employeeId: string;
  clientId: string;
  name: string;
  createdAt: string;
  status: string;
  contractId: string;
};

export type DetailsFilters = {
  name: string;
  minPrice: number;
  maxPrice: number;
  type: string;
};

export type Detail = {
  id: string;
  name: string;
  price: number;
  type: string;
};

export type Employee = {
  id: string;
  name: string;
  position: string;
  employeeNum: string;
};

export type FilterValues<T> = Partial<Omit<T, 'id'>>;
