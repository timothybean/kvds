"use strict";(self.webpackChunkfuse_react_app=self.webpackChunkfuse_react_app||[]).push([[6931],{44269:function(n,e,t){t.d(e,{Z:function(){return I}});var o=t(29439),r=t(98655),a=t(73428),i=t(65280),l=t(5297),u=t(83061),c=t(47313),s=t(17551),m=t(9506),d=t(1413),p=t(45987),x=t(1168),f=t(10223),h=t(83097),b=t(9496),j=t(54097),B=t(19860),Z=t(17592),v=t(52437),g=t(46417),w=["children","name"];function y(n){var e=n.children,t=n.document,o=(0,B.Z)();c.useEffect((function(){t.body.dir=o.direction}),[t,o.direction]);var r=c.useMemo((function(){return(0,h.Z)({key:"iframe-demo-".concat(o.direction),prepend:!0,container:t.head,stylisPlugins:"rtl"===o.direction?[f.Z]:[]})}),[t,o.direction]),a=c.useCallback((function(){return t.defaultView}),[t]);return(0,g.jsx)(j.StyleSheetManager,{target:t.head,stylisPlugins:"rtl"===o.direction?[f.Z]:[],children:(0,g.jsxs)(b.C,{value:r,children:[(0,g.jsx)(v.Z,{styles:function(){return{html:{fontSize:"62.5%"}}}}),c.cloneElement(e,{window:a})]})})}var G=(0,Z.ZP)("iframe")((function(n){var e=n.theme;return{backgroundColor:e.palette.background.default,flexGrow:1,height:400,border:0,boxShadow:e.shadows[1]}}));function k(n){var e,t=n.children,r=n.name,a=(0,p.Z)(n,w),i="".concat(r," demo"),l=c.useRef(null),u=c.useReducer((function(){return!0}),!1),s=(0,o.Z)(u,2),m=s[0],f=s[1];c.useEffect((function(){var n=l.current.contentDocument;null==n||"complete"!==n.readyState||m||f()}),[m]);var h=null===(e=l.current)||void 0===e?void 0:e.contentDocument;return(0,g.jsxs)(g.Fragment,{children:[(0,g.jsx)(G,(0,d.Z)({onLoad:f,ref:l,title:i},a)),!1!==m?x.createPortal((0,g.jsx)(y,{document:h,children:t}),h.body):null]})}var N=c.memo(k),T=t(56993);function C(n){var e=(0,c.useState)(n.currentTabIndex),t=(0,o.Z)(e,2),d=t[0],p=t[1],x=n.component,f=n.raw,h=n.iframe,b=n.className,j=n.name;return(0,g.jsxs)(a.Z,{className:(0,u.default)(b,"shadow"),sx:{backgroundColor:function(n){return(0,s._j)(n.palette.background.paper,"light"===n.palette.mode?.01:.1)}},children:[(0,g.jsx)(m.Z,{sx:{backgroundColor:function(n){return(0,s._j)(n.palette.background.paper,"light"===n.palette.mode?.02:.2)}},children:(0,g.jsxs)(l.Z,{classes:{root:"border-b-1",flexContainer:"justify-end"},value:d,onChange:function(n,e){p(e)},textColor:"secondary",indicatorColor:"secondary",children:[x&&(0,g.jsx)(i.Z,{classes:{root:"min-w-64"},icon:(0,g.jsx)(T.Z,{children:"heroicons-outline:eye"})}),f&&(0,g.jsx)(i.Z,{classes:{root:"min-w-64"},icon:(0,g.jsx)(T.Z,{children:"heroicons-outline:code"})})]})}),(0,g.jsxs)("div",{className:"flex justify-center max-w-full relative",children:[(0,g.jsx)("div",{className:0===d?"flex flex-1 max-w-full":"hidden",children:x&&(h?(0,g.jsx)(N,{name:j,children:(0,g.jsx)(x,{})}):(0,g.jsx)("div",{className:"p-24 flex flex-1 justify-center max-w-full",children:(0,g.jsx)(x,{})}))}),(0,g.jsx)("div",{className:1===d?"flex flex-1":"hidden",children:(0,c.useMemo)((function(){return f&&1===d?(0,g.jsx)("div",{className:"flex flex-1",children:(0,g.jsx)(r.Z,{component:"pre",className:"language-javascript w-full",sx:{borderRadius:"0!important"},children:f.default})}):null}),[f,d])})]})]})}C.defaultProps={name:"",currentTabIndex:0};var I=C},75582:function(n,e,t){t.d(e,{Z:function(){return i}});t(47313);var o=t(24193),r=t(84060),a=t(46417);function i(){return(0,a.jsxs)(r.Z,{variant:"contained","aria-label":"outlined primary button group",children:[(0,a.jsx)(o.Z,{children:"One"}),(0,a.jsx)(o.Z,{children:"Two"}),(0,a.jsx)(o.Z,{children:"Three"})]})}},90885:function(n,e,t){t.d(e,{Z:function(){return i}});t(47313);var o=t(84060),r=t(24193),a=t(46417);function i(){return(0,a.jsxs)(o.Z,{disableElevation:!0,variant:"contained","aria-label":"Disabled elevation buttons",children:[(0,a.jsx)(r.Z,{children:"One"}),(0,a.jsx)(r.Z,{children:"Two"})]})}},98540:function(n,e,t){t.d(e,{Z:function(){return u}});t(47313);var o=t(24193),r=t(84060),a=t(9506),i=t(46417),l=[(0,i.jsx)(o.Z,{children:"One"},"one"),(0,i.jsx)(o.Z,{children:"Two"},"two"),(0,i.jsx)(o.Z,{children:"Three"},"three")];function u(){return(0,i.jsxs)(a.Z,{sx:{display:"flex","& > *":{m:1}},children:[(0,i.jsx)(r.Z,{orientation:"vertical","aria-label":"vertical outlined button group",children:l}),(0,i.jsx)(r.Z,{orientation:"vertical","aria-label":"vertical contained button group",variant:"contained",children:l}),(0,i.jsx)(r.Z,{orientation:"vertical","aria-label":"vertical contained button group",variant:"text",children:l})]})}},6317:function(n,e,t){t.d(e,{Z:function(){return u}});t(47313);var o=t(24193),r=t(9506),a=t(84060),i=t(46417),l=[(0,i.jsx)(o.Z,{children:"One"},"one"),(0,i.jsx)(o.Z,{children:"Two"},"two"),(0,i.jsx)(o.Z,{children:"Three"},"three")];function u(){return(0,i.jsxs)(r.Z,{sx:{display:"flex",flexDirection:"column",alignItems:"center","& > *":{m:1}},children:[(0,i.jsx)(a.Z,{size:"small","aria-label":"small button group",children:l}),(0,i.jsx)(a.Z,{color:"secondary","aria-label":"medium secondary button group",children:l}),(0,i.jsx)(a.Z,{size:"large","aria-label":"large button group",children:l})]})}},30243:function(n,e,t){t.d(e,{Z:function(){return b}});var o=t(1413),r=t(29439),a=t(47313),i=t(24193),l=t(84060),u=t(54406),c=t(31685),s=t(73365),m=t(70501),d=t(88975),p=t(51405),x=t(14560),f=t(46417),h=["Create a merge commit","Squash and merge","Rebase and merge"];function b(){var n=a.useState(!1),e=(0,r.Z)(n,2),t=e[0],b=e[1],j=a.useRef(null),B=a.useState(1),Z=(0,r.Z)(B,2),v=Z[0],g=Z[1],w=function(n){j.current&&j.current.contains(n.target)||b(!1)};return(0,f.jsxs)(a.Fragment,{children:[(0,f.jsxs)(l.Z,{variant:"contained",ref:j,"aria-label":"split button",children:[(0,f.jsx)(i.Z,{onClick:function(){console.info("You clicked ".concat(h[v]))},children:h[v]}),(0,f.jsx)(i.Z,{size:"small","aria-controls":t?"split-button-menu":void 0,"aria-expanded":t?"true":void 0,"aria-label":"select merge strategy","aria-haspopup":"menu",onClick:function(){b((function(n){return!n}))},children:(0,f.jsx)(u.Z,{})})]}),(0,f.jsx)(d.Z,{sx:{zIndex:1},open:t,anchorEl:j.current,role:void 0,transition:!0,disablePortal:!0,children:function(n){var e=n.TransitionProps,t=n.placement;return(0,f.jsx)(s.Z,(0,o.Z)((0,o.Z)({},e),{},{style:{transformOrigin:"bottom"===t?"center top":"center bottom"},children:(0,f.jsx)(m.Z,{children:(0,f.jsx)(c.Z,{onClickAway:w,children:(0,f.jsx)(x.Z,{id:"split-button-menu",autoFocusItem:!0,children:h.map((function(n,e){return(0,f.jsx)(p.Z,{disabled:2===e,selected:e===v,onClick:function(n){return function(n,e){g(e),b(!1)}(0,e)},children:n},n)}))})})})}))}})]})}},99698:function(n,e,t){t.d(e,{Z:function(){return l}});t(47313);var o=t(24193),r=t(84060),a=t(9506),i=t(46417);function l(){return(0,i.jsxs)(a.Z,{sx:{display:"flex",flexDirection:"column",alignItems:"center","& > *":{m:1}},children:[(0,i.jsxs)(r.Z,{variant:"outlined","aria-label":"outlined button group",children:[(0,i.jsx)(o.Z,{children:"One"}),(0,i.jsx)(o.Z,{children:"Two"}),(0,i.jsx)(o.Z,{children:"Three"})]}),(0,i.jsxs)(r.Z,{variant:"text","aria-label":"text button group",children:[(0,i.jsx)(o.Z,{children:"One"}),(0,i.jsx)(o.Z,{children:"Two"}),(0,i.jsx)(o.Z,{children:"Three"})]})]})}},76931:function(n,e,t){t.r(e);var o=t(44269),r=t(56993),a=t(24193),i=t(61113),l=t(46417);e.default=function(n){return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)("div",{className:"flex flex-1 grow-0 items-center justify-end",children:(0,l.jsx)(a.Z,{className:"normal-case",variant:"contained",color:"secondary",component:"a",href:"https://mui.com/components/button-group",target:"_blank",role:"button",startIcon:(0,l.jsx)(r.Z,{children:"heroicons-outline:external-link"}),children:"Reference"})}),(0,l.jsx)(i.Z,{className:"text-40 my-16 font-700",component:"h1",children:"Button Group"}),(0,l.jsx)(i.Z,{className:"description",children:"The ButtonGroup component can be used to group related buttons."}),(0,l.jsx)(i.Z,{className:"text-32 mt-40 mb-10 font-700",component:"h2",children:"Basic button group"}),(0,l.jsxs)(i.Z,{className:"mb-40",component:"div",children:["The buttons can be grouped by wrapping them with the ",(0,l.jsx)("code",{children:"ButtonGroup"})," component. They need to be immediate children."]}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:(0,l.jsx)(o.Z,{name:"BasicButtonGroup.js",className:"my-24",iframe:!1,component:t(75582).Z,raw:t(96527)})}),(0,l.jsx)(i.Z,{className:"text-32 mt-40 mb-10 font-700",component:"h2",children:"Button variants"}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:"All the standard button variants are supported."}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:(0,l.jsx)(o.Z,{name:"VariantButtonGroup.js",className:"my-24",iframe:!1,component:t(99698).Z,raw:t(1285)})}),(0,l.jsx)(i.Z,{className:"text-32 mt-40 mb-10 font-700",component:"h2",children:"Sizes and colors"}),(0,l.jsxs)(i.Z,{className:"mb-40",component:"div",children:["The ",(0,l.jsx)("code",{children:"size"})," and ",(0,l.jsx)("code",{children:"color"})," props can be used to control the appearance of the button group."]}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:(0,l.jsx)(o.Z,{name:"GroupSizesColors.js",className:"my-24",iframe:!1,component:t(6317).Z,raw:t(98242)})}),(0,l.jsx)(i.Z,{className:"text-32 mt-40 mb-10 font-700",component:"h2",children:"Vertical group"}),(0,l.jsxs)(i.Z,{className:"mb-40",component:"div",children:["The button group can be displayed vertically using the ",(0,l.jsx)("code",{children:"orientation"})," prop."]}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:(0,l.jsx)(o.Z,{name:"GroupOrientation.js",className:"my-24",iframe:!1,component:t(98540).Z,raw:t(81081)})}),(0,l.jsx)(i.Z,{className:"text-32 mt-40 mb-10 font-700",component:"h2",children:"Split button"}),(0,l.jsxs)(i.Z,{className:"mb-40",component:"div",children:[(0,l.jsx)("code",{children:"ButtonGroup"})," can also be used to create a split button. The dropdown can change the button action (as in this example) or be used to immediately trigger a related action."]}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:(0,l.jsx)(o.Z,{name:"SplitButton.js",className:"my-24",iframe:!1,component:t(30243).Z,raw:t(11800)})}),(0,l.jsx)(i.Z,{className:"text-32 mt-40 mb-10 font-700",component:"h2",children:"Disabled elevation"}),(0,l.jsxs)(i.Z,{className:"mb-40",component:"div",children:["You can remove the elevation with the ",(0,l.jsx)("code",{children:"disableElevation"})," prop."]}),(0,l.jsx)(i.Z,{className:"mb-40",component:"div",children:(0,l.jsx)(o.Z,{name:"DisableElevation.js",className:"my-24",iframe:!1,component:t(90885).Z,raw:t(45185)})})]})}},54406:function(n,e,t){var o=t(64836);e.Z=void 0;var r=o(t(45045)),a=t(46417),i=(0,r.default)((0,a.jsx)("path",{d:"m7 10 5 5 5-5z"}),"ArrowDropDown");e.Z=i},96527:function(n,e,t){t.r(e),e.default="import * as React from 'react';\nimport Button from '@mui/material/Button';\nimport ButtonGroup from '@mui/material/ButtonGroup';\n\nexport default function BasicButtonGroup() {\n  return (\n    <ButtonGroup variant=\"contained\" aria-label=\"outlined primary button group\">\n      <Button>One</Button>\n      <Button>Two</Button>\n      <Button>Three</Button>\n    </ButtonGroup>\n  );\n}\n"},45185:function(n,e,t){t.r(e),e.default="import * as React from 'react';\nimport ButtonGroup from '@mui/material/ButtonGroup';\nimport Button from '@mui/material/Button';\n\nexport default function DisableElevation() {\n  return (\n    <ButtonGroup\n      disableElevation\n      variant=\"contained\"\n      aria-label=\"Disabled elevation buttons\"\n    >\n      <Button>One</Button>\n      <Button>Two</Button>\n    </ButtonGroup>\n  );\n}\n"},81081:function(n,e,t){t.r(e),e.default='import * as React from \'react\';\nimport Button from \'@mui/material/Button\';\nimport ButtonGroup from \'@mui/material/ButtonGroup\';\nimport Box from \'@mui/material/Box\';\n\nconst buttons = [\n  <Button key="one">One</Button>,\n  <Button key="two">Two</Button>,\n  <Button key="three">Three</Button>,\n];\n\nexport default function GroupOrientation() {\n  return (\n    <Box\n      sx={{\n        display: \'flex\',\n        \'& > *\': {\n          m: 1,\n        },\n      }}\n    >\n      <ButtonGroup\n        orientation="vertical"\n        aria-label="vertical outlined button group"\n      >\n        {buttons}\n      </ButtonGroup>\n      <ButtonGroup\n        orientation="vertical"\n        aria-label="vertical contained button group"\n        variant="contained"\n      >\n        {buttons}\n      </ButtonGroup>\n      <ButtonGroup\n        orientation="vertical"\n        aria-label="vertical contained button group"\n        variant="text"\n      >\n        {buttons}\n      </ButtonGroup>\n    </Box>\n  );\n}\n'},98242:function(n,e,t){t.r(e),e.default='import * as React from \'react\';\nimport Button from \'@mui/material/Button\';\nimport Box from \'@mui/material/Box\';\nimport ButtonGroup from \'@mui/material/ButtonGroup\';\n\nconst buttons = [\n  <Button key="one">One</Button>,\n  <Button key="two">Two</Button>,\n  <Button key="three">Three</Button>,\n];\n\nexport default function GroupSizesColors() {\n  return (\n    <Box\n      sx={{\n        display: \'flex\',\n        flexDirection: \'column\',\n        alignItems: \'center\',\n        \'& > *\': {\n          m: 1,\n        },\n      }}\n    >\n      <ButtonGroup size="small" aria-label="small button group">\n        {buttons}\n      </ButtonGroup>\n      <ButtonGroup color="secondary" aria-label="medium secondary button group">\n        {buttons}\n      </ButtonGroup>\n      <ButtonGroup size="large" aria-label="large button group">\n        {buttons}\n      </ButtonGroup>\n    </Box>\n  );\n}\n'},11800:function(n,e,t){t.r(e),e.default="import * as React from 'react';\nimport Button from '@mui/material/Button';\nimport ButtonGroup from '@mui/material/ButtonGroup';\nimport ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';\nimport ClickAwayListener from '@mui/material/ClickAwayListener';\nimport Grow from '@mui/material/Grow';\nimport Paper from '@mui/material/Paper';\nimport Popper from '@mui/material/Popper';\nimport MenuItem from '@mui/material/MenuItem';\nimport MenuList from '@mui/material/MenuList';\n\nconst options = ['Create a merge commit', 'Squash and merge', 'Rebase and merge'];\n\nexport default function SplitButton() {\n  const [open, setOpen] = React.useState(false);\n  const anchorRef = React.useRef(null);\n  const [selectedIndex, setSelectedIndex] = React.useState(1);\n\n  const handleClick = () => {\n    console.info(`You clicked ${options[selectedIndex]}`);\n  };\n\n  const handleMenuItemClick = (event, index) => {\n    setSelectedIndex(index);\n    setOpen(false);\n  };\n\n  const handleToggle = () => {\n    setOpen((prevOpen) => !prevOpen);\n  };\n\n  const handleClose = (event) => {\n    if (anchorRef.current && anchorRef.current.contains(event.target)) {\n      return;\n    }\n\n    setOpen(false);\n  };\n\n  return (\n    <React.Fragment>\n      <ButtonGroup variant=\"contained\" ref={anchorRef} aria-label=\"split button\">\n        <Button onClick={handleClick}>{options[selectedIndex]}</Button>\n        <Button\n          size=\"small\"\n          aria-controls={open ? 'split-button-menu' : undefined}\n          aria-expanded={open ? 'true' : undefined}\n          aria-label=\"select merge strategy\"\n          aria-haspopup=\"menu\"\n          onClick={handleToggle}\n        >\n          <ArrowDropDownIcon />\n        </Button>\n      </ButtonGroup>\n      <Popper\n        sx={{\n          zIndex: 1,\n        }}\n        open={open}\n        anchorEl={anchorRef.current}\n        role={undefined}\n        transition\n        disablePortal\n      >\n        {({ TransitionProps, placement }) => (\n          <Grow\n            {...TransitionProps}\n            style={{\n              transformOrigin:\n                placement === 'bottom' ? 'center top' : 'center bottom',\n            }}\n          >\n            <Paper>\n              <ClickAwayListener onClickAway={handleClose}>\n                <MenuList id=\"split-button-menu\" autoFocusItem>\n                  {options.map((option, index) => (\n                    <MenuItem\n                      key={option}\n                      disabled={index === 2}\n                      selected={index === selectedIndex}\n                      onClick={(event) => handleMenuItemClick(event, index)}\n                    >\n                      {option}\n                    </MenuItem>\n                  ))}\n                </MenuList>\n              </ClickAwayListener>\n            </Paper>\n          </Grow>\n        )}\n      </Popper>\n    </React.Fragment>\n  );\n}\n"},1285:function(n,e,t){t.r(e),e.default="import * as React from 'react';\nimport Button from '@mui/material/Button';\nimport ButtonGroup from '@mui/material/ButtonGroup';\nimport Box from '@mui/material/Box';\n\nexport default function VariantButtonGroup() {\n  return (\n    <Box\n      sx={{\n        display: 'flex',\n        flexDirection: 'column',\n        alignItems: 'center',\n        '& > *': {\n          m: 1,\n        },\n      }}\n    >\n      <ButtonGroup variant=\"outlined\" aria-label=\"outlined button group\">\n        <Button>One</Button>\n        <Button>Two</Button>\n        <Button>Three</Button>\n      </ButtonGroup>\n      <ButtonGroup variant=\"text\" aria-label=\"text button group\">\n        <Button>One</Button>\n        <Button>Two</Button>\n        <Button>Three</Button>\n      </ButtonGroup>\n    </Box>\n  );\n}\n"}}]);