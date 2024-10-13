import { create } from 'zustand';
import { ClientFilters, FilterValues } from '@/types';

type Store = {
  filters: FilterValues<ClientFilters>;
  setFilters: (payload: FilterValues<ClientFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    name: '',
  },
  setFilters: (payload: FilterValues<ClientFilters>) => set(() => ({ filters: payload })),
}));
