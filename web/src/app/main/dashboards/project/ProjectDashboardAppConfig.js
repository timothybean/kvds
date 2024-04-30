import { lazy } from 'react';

const CockpitDashboardApp = lazy(() => import('./ProjectDashboardApp'));

const ProjectDashboardAppConfig = {
  settings: {
    layout: {
      config: {},
    },
  },
  routes: [
    {
      path: 'dashboards/cockpit',
      element: <CockpitDashboardApp />,
    },
  ],
};

export default ProjectDashboardAppConfig;
