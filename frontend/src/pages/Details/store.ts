import { create } from 'zustand';
import { DetailsFilters, FilterValues } from '@/types';

type Store = {
  filters: FilterValues<DetailsFilters>;
  setFilters: (payload: FilterValues<DetailsFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    name: '',
  },
  setFilters: (payload: FilterValues<DetailsFilters>) => set(() => ({ filters: payload })),
}));
