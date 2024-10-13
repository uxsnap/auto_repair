import { create } from 'zustand';
import { ClientFilters, FilterValues, VehiclesFilters } from '@/types';

type Store = {
  filters: FilterValues<VehiclesFilters>;
  setFilters: (payload: FilterValues<VehiclesFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    vehicleNumber: '',
  },
  setFilters: (payload: FilterValues<VehiclesFilters>) => set(() => ({ filters: payload })),
}));
