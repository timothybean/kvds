import AnalyticsDashboardAppConfig from './analytics/AnalyticsDashboardAppConfig';
import CockpitDashboardApp from './project/ProjectDashboardAppConfig';
import FinanceDashboardAppConfig from './finance/FinanceDashboardAppConfig';
import CryptoDashboardAppConfig from './crypto/CryptoDashboardAppConfig';

const dashboardsConfigs = [
  AnalyticsDashboardAppConfig,
  CockpitDashboardApp,
  FinanceDashboardAppConfig,
  CryptoDashboardAppConfig,
];

export default dashboardsConfigs;
