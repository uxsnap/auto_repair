import { create } from 'zustand';
import { ReceiptsFilters, FilterValues } from '@/types';

type Store = {
  filters: FilterValues<ReceiptsFilters>;
  setFilters: (payload: FilterValues<ReceiptsFilters>) => void;
};

export const useFiltersStore = create<Store>()((set) => ({
  filters: {
    name: '',
  },
  setFilters: (payload: FilterValues<ReceiptsFilters>) => set(() => ({ filters: payload })),
}));
