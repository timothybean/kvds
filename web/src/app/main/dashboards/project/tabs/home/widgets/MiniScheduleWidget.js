import Paper from '@mui/material/Paper';
import { lighten, useTheme } from '@mui/material/styles';
import Tab from '@mui/material/Tab';
import Tabs from '@mui/material/Tabs';
import Typography from '@mui/material/Typography';
import { memo, useEffect, useState, useMemo, useCallback } from 'react';
import ReactApexChart from 'react-apexcharts';
import Box from '@mui/material/Box';
import { useSelector } from 'react-redux';
import { selectWidgets } from '../../../store/widgetsSlice';
import { Epg, useEpg, Layout } from "planby";
import { Timeline, ChannelItem, ProgramItem } from "./components";
import { fetchChannels, fetchEpg } from "./helpers";

function MiniScheduleWidget() {
  const theme = useTheme();
  const [awaitRender, setAwaitRender] = useState(true);
  const [tabValue, setTabValue] = useState(0);
  const widgets = useSelector(selectWidgets);
  const { overview, series, ranges, labels } = widgets?.githubIssues;
  const currentRange = Object.keys(ranges)[tabValue];

  const [channels, setChannels] = useState([]);
  const [epg, setEpg] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const channelsData = useMemo(() => channels, [channels]);
  const epgData = useMemo(() => epg, [epg]);

  const handleFetchResources = useCallback(async () => {
    setIsLoading(true);
    const epg = await fetchEpg();
    const channels = await fetchChannels();
    setEpg(epg);
    setChannels(channels);
    setIsLoading(false);
  }, []);

  useEffect(() => {
    handleFetchResources();
  }, [handleFetchResources]);

  const wtheme = {
    primary: {
      600: "#1a202c",
      900: "#171923"
    },
    grey: { 300: "#d1d1d1" },
    white: "#fff",
    green: {
      300: "#2c7a7b"
    },
    scrollbar: {
      border: "#ffffff",
      thumb: {
        bg: "#e1e1e1"
      }
    },
    loader: {
      teal: "#5DDADB",
      purple: "#3437A2",
      pink: "#F78EB6",
      bg: "#171923db"
    },
    gradient: {
      blue: {
        300: "#002eb3",
        600: "#002360",
        900: "#051937"
      }
    },
  
    text: {
      grey: {
        300: "#a0aec0",
        500: "#718096"
      }
    },
  
    timeline: {
      divider: {
        bg: "#718096"
      }
    }
  };

  const { getEpgProps, getLayoutProps } = useEpg({
    channels: channelsData,
    epg: epgData,
    dayWidth: 7200,
    sidebarWidth: 100,
    itemHeight: 80,
    isSidebar: true,
    isTimeline: true,
    isLine: true,
    startDate: "2022-10-18T00:00:00",
    endDate: "2022-10-18T24:00:00",
    isBaseTimeFormat: true,
    wtheme
  });

  useEffect(() => {
    setAwaitRender(false);
  }, []);

  if (awaitRender) {
    return null;
  }

  return (
    <Paper className="flex flex-col flex-auto p-24 shadow rounded-2xl overflow-hidden">
      <div className="flex flex-col sm:flex-row items-start justify-between">
        <Typography className="text-lg font-medium tracking-tight leading-6 truncate">
          Lesson Schedule
        </Typography>
        <div className="mt-12 sm:mt-0 sm:ml-8">
          <Tabs
            value={tabValue}
            onChange={(ev, value) => setTabValue(value)}
            indicatorColor="secondary"
            textColor="inherit"
            variant="scrollable"
            scrollButtons={false}
            className="-mx-4 min-h-40"
            classes={{ indicator: 'flex justify-center bg-transparent w-full h-full' }}
            TabIndicatorProps={{
              children: (
                <Box
                  sx={{ bgcolor: 'text.disabled' }}
                  className="w-full h-full rounded-full opacity-20"
                />
              ),
            }}
          >
            {Object.entries(ranges).map(([key, label]) => (
              <Tab
                className="text-14 font-semibold min-h-40 min-w-64 mx-4 px-12"
                disableRipple
                key={key}
                label={label}
              />
            ))}
          </Tabs>
        </div>
      </div>
      <div className="grid grid-cols-1 lg:grid-cols-1 grid-flow-row gap-24 w-full mt-32 sm:mt-16">
        <div className="flex flex-col flex-auto">
        <Epg isLoading={isLoading} {...getEpgProps()}>
          <Layout
            {...getLayoutProps()}
            renderTimeline={(props) => <Timeline {...props} />}
            renderProgram={({ program, ...rest }) => (
              <ProgramItem key={program.data.id} program={program} {...rest} />
            )}
            renderChannel={({ channel }) => (
              <ChannelItem key={channel.uuid} channel={channel} />
            )}
          />
        </Epg>
        </div>
      </div>
    </Paper>
  );
}

export default memo(MiniScheduleWidget);
