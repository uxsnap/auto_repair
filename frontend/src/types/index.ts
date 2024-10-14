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

export type Storage = {
  id: string;
  employeeId: string;
  storageNum: string;
  detailId: string;
  detailCount: number;
};

export type Client = {
  id: string;
  employeeId: string;
  name: string;
  phone: string;
  passport: string;
  hasDocuments: boolean;
};

export type ClientWithData = {
  id: string;
  employee: {
    id: string;
    name: string;
  };
  name: string;
  phone: string;
  passport: string;
  hasDocuments: boolean;
};

export type StorageWithData = {
  id: string;
  storageNum: string;
  employee: {
    id: string;
    name: string;
  };
  detail: {
    id: string;
    name: string;
  };
  isDeleted: boolean;
  detailCount: number;
};

export type StorageFilters = {
  storageNum: string;
  employeeName: string;
  detailName: string;
};

export type ClientFilters = {
  name: string;
  employeeName: string;
  passport: string;
  phone: string;
};

export type VehiclesFilters = {
  vehicleNumber: string;
  clientName: string;
  brand: string;
  model: string;
};

export type ActFilters = {
  name: string;
  applicationName: string;
  serviceName: string;
  createdAt: Date;
};

export type Vehicle = {
  id: string;
  vehicleNumber: string;
  clientId: string;
  brand: string;
  model: string;
};

export type VehicleWithData = {
  id: string;
  vehicleNumber: string;
  client: {
    id: string;
    name: string;
  };
  brand: string;
  model: string;
};

export type Act = {
  id: string;
  name: string;
  createdAt: string;
  applicationId: string;
  serviceId: string;
};

export type ActWithData = {
  id: string;
  name: string;
  createdAt: string;
  application: {
    id: string;
    name: string;
  };
  service: {
    id: string;
    name: string;
  };
};

export type Service = {
  id: string;
  name: string;
};

export type FilterValues<T> = Partial<Omit<T, 'id'>>;
