import FuseSvgIcon from '@fuse/core/FuseSvgIcon';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

function NoSsrDoc(props) {
  return (
    <>
      <div className="flex flex-1 grow-0 items-center justify-end">
        <Button
          className="normal-case"
          variant="contained"
          color="secondary"
          component="a"
          href="https://mui.com/components/no-ssr"
          target="_blank"
          role="button"
          startIcon={<FuseSvgIcon>heroicons-outline:external-link</FuseSvgIcon>}
        >
          Reference
        </Button>
      </div>
      <Typography className="text-40 my-16 font-700" component="h1">
        No SSR
      </Typography>
      <Typography className="description">
        The No-SSR component defers the rendering of children components from the server to the
        client.
      </Typography>

      <Typography className="text-32 mt-40 mb-10 font-700" component="h2">
        This document has moved
      </Typography>
      <Typography className="mb-40" component="div">
        :::warning Please refer to the <a href="/base/react-no-ssr/">No-SSR</a> component page in
        the Base UI docs for demos and details on usage.
      </Typography>
      <Typography className="mb-40" component="div">
        No-SSR is a part of the standalone <a href="/base/getting-started/overview/">Base UI</a>{' '}
        component library. It is currently re-exported from <code>@mui/material</code> for your
        convenience, but it will be removed from this package in a future major version, after{' '}
        <code>@mui/base</code> gets a stable release. :::
      </Typography>
    </>
  );
}

export default NoSsrDoc;
