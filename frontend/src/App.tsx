import '@mantine/core/styles.css';

import { IconEngine, IconPlus, IconUserFilled } from '@tabler/icons-react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import {
  createBrowserRouter,
  Route,
  RouterProvider,
  Routes,
  useLocation,
  useNavigate,
} from 'react-router-dom';
import { AppShell, Group, MantineProvider, Tabs, Title } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import { DetailsPage } from './pages/Details/Details.page';
import { theme } from './theme';

import '@mantine/notifications/styles.css';
import '@mantine/dates/styles.css';
import 'dayjs/locale/ru';

import { DatesProvider } from '@mantine/dates';
import { useDisclosure } from '@mantine/hooks';
import { AddEmployee } from './components/AddEmployee';
import { EmployeeModal } from './components/EmployeeModal';
import { ActsPage } from './pages/Acts/Act.page';
import { AppsPage } from './pages/Apps/Apps.page';
import { ClientsPage } from './pages/Clients/Client.page';
import { ContractsPage } from './pages/Contracts/Contract.page';
import { ReceiptsPage } from './pages/Receipts/Receipt.page';
import { StoragePage } from './pages/Storage/Storage.page';
import { VehiclesPage } from './pages/Vehicles/Vehicle.page';

const queryClient = new QueryClient();

const router = createBrowserRouter([{ path: '*', element: <Root /> }]);

function Root() {
  const navigate = useNavigate();
  const location = useLocation();

  const [openedEmployeeModal, { open: openEmployeeModal, close: closeEmployeeModal }] =
    useDisclosure(false);
  const [openedAddEmployee, { open: openAddEmployee, close: closeAddEmployee }] =
    useDisclosure(false);

  return (
    <MantineProvider theme={theme}>
      <Notifications />

      <DatesProvider settings={{ locale: 'ru' }}>
        <QueryClientProvider client={queryClient}>
          <EmployeeModal opened={openedEmployeeModal} close={closeEmployeeModal} />
          <AddEmployee opened={openedAddEmployee} close={closeAddEmployee} />

          <AppShell header={{ height: 60 }} padding="md">
            <AppShell.Header
              display="flex"
              style={{ alignItems: 'center', justifyContent: 'space-between' }}
              px={20}
              c="blue"
            >
              <Group gap={12} align="center">
                <IconEngine size={40} />
                <Title order={3}>Auto Repair</Title>
              </Group>

              <Group gap={12} align="center">
                <IconUserFilled
                  size={40}
                  style={{ cursor: 'pointer' }}
                  onClick={openEmployeeModal}
                />
                <IconPlus size={40} style={{ cursor: 'pointer' }} onClick={openAddEmployee} />
              </Group>
            </AppShell.Header>

            <AppShell.Main>
              <Tabs
                defaultValue={location.pathname}
                onChange={(value) => navigate(`${value === '/' ? '' : value}`)}
                value={location.pathname}
                variant="outline"
              >
                <Tabs.List>
                  <Tabs.Tab value="/">Заявки</Tabs.Tab>
                  <Tabs.Tab value="/contracts">Договора</Tabs.Tab>
                  <Tabs.Tab value="/receipts">Чеки об оплате</Tabs.Tab>
                  <Tabs.Tab value="/clients">Клиенты</Tabs.Tab>
                  <Tabs.Tab value="/vehicles">ТС</Tabs.Tab>
                  <Tabs.Tab value="/details">Детали</Tabs.Tab>
                  <Tabs.Tab value="/storage">Склад</Tabs.Tab>
                  <Tabs.Tab value="/act">Акты</Tabs.Tab>
                </Tabs.List>

                <Routes>
                  <Route path="/clients" element={<ClientsPage />} />
                  <Route path="/contracts" element={<ContractsPage />} />
                  <Route path="/receipts" element={<ReceiptsPage />} />
                  <Route path="/vehicles" element={<VehiclesPage />} />
                  <Route path="/details" element={<DetailsPage />} />
                  <Route path="/storage" element={<StoragePage />} />
                  <Route path="/act" element={<ActsPage />} />
                  <Route path="/" element={<AppsPage />} />
                </Routes>
              </Tabs>
            </AppShell.Main>
          </AppShell>
        </QueryClientProvider>
      </DatesProvider>
    </MantineProvider>
  );
}

export default function App() {
  return <RouterProvider router={router} />;
}
