"use strict";(self.webpackChunkfuse_react_app=self.webpackChunkfuse_react_app||[]).push([[2083],{68728:function(r,e,n){var t=n(64836);e.Z=void 0;var a=t(n(45045)),o=n(46417),i=(0,a.default)((0,o.jsx)("path",{d:"M9 16.17 4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"}),"Check");e.Z=i},83121:function(r,e,n){var t=n(64836);e.Z=void 0;var a=t(n(45045)),o=n(46417),i=(0,a.default)((0,o.jsx)("path",{d:"M17 3H5c-1.11 0-2 .9-2 2v14c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V7l-4-4zm-5 16c-1.66 0-3-1.34-3-3s1.34-3 3-3 3 1.34 3 3-1.34 3-3 3zm3-10H5V5h10v4z"}),"Save");e.Z=i},98668:function(r,e,n){var t,a,o,i,c,s,l,u,d=n(30168),f=n(63366),v=n(87462),m=n(47313),p=n(83061),b=n(21921),h=n(30686),Z=n(91615),g=n(77342),k=n(17592),y=n(94808),S=n(46417),x=["className","color","disableShrink","size","style","thickness","value","variant"],w=44,C=(0,h.F4)(c||(c=t||(t=(0,d.Z)(["\n  0% {\n    transform: rotate(0deg);\n  }\n\n  100% {\n    transform: rotate(360deg);\n  }\n"])))),P=(0,h.F4)(s||(s=a||(a=(0,d.Z)(["\n  0% {\n    stroke-dasharray: 1px, 200px;\n    stroke-dashoffset: 0;\n  }\n\n  50% {\n    stroke-dasharray: 100px, 200px;\n    stroke-dashoffset: -15px;\n  }\n\n  100% {\n    stroke-dasharray: 100px, 200px;\n    stroke-dashoffset: -125px;\n  }\n"])))),M=(0,k.ZP)("span",{name:"MuiCircularProgress",slot:"Root",overridesResolver:function(r,e){var n=r.ownerState;return[e.root,e[n.variant],e["color".concat((0,Z.Z)(n.color))]]}})((function(r){var e=r.ownerState,n=r.theme;return(0,v.Z)({display:"inline-block"},"determinate"===e.variant&&{transition:n.transitions.create("transform")},"inherit"!==e.color&&{color:(n.vars||n).palette[e.color].main})}),(function(r){return"indeterminate"===r.ownerState.variant&&(0,h.iv)(l||(l=o||(o=(0,d.Z)(["\n      animation: "," 1.4s linear infinite;\n    "]))),C)})),j=(0,k.ZP)("svg",{name:"MuiCircularProgress",slot:"Svg",overridesResolver:function(r,e){return e.svg}})({display:"block"}),N=(0,k.ZP)("circle",{name:"MuiCircularProgress",slot:"Circle",overridesResolver:function(r,e){var n=r.ownerState;return[e.circle,e["circle".concat((0,Z.Z)(n.variant))],n.disableShrink&&e.circleDisableShrink]}})((function(r){var e=r.ownerState,n=r.theme;return(0,v.Z)({stroke:"currentColor"},"determinate"===e.variant&&{transition:n.transitions.create("stroke-dashoffset")},"indeterminate"===e.variant&&{strokeDasharray:"80px, 200px",strokeDashoffset:0})}),(function(r){var e=r.ownerState;return"indeterminate"===e.variant&&!e.disableShrink&&(0,h.iv)(u||(u=i||(i=(0,d.Z)(["\n      animation: "," 1.4s ease-in-out infinite;\n    "]))),P)})),R=m.forwardRef((function(r,e){var n=(0,g.Z)({props:r,name:"MuiCircularProgress"}),t=n.className,a=n.color,o=void 0===a?"primary":a,i=n.disableShrink,c=void 0!==i&&i,s=n.size,l=void 0===s?40:s,u=n.style,d=n.thickness,m=void 0===d?3.6:d,h=n.value,k=void 0===h?0:h,C=n.variant,P=void 0===C?"indeterminate":C,R=(0,f.Z)(n,x),B=(0,v.Z)({},n,{color:o,disableShrink:c,size:l,thickness:m,value:k,variant:P}),D=function(r){var e=r.classes,n=r.variant,t=r.color,a=r.disableShrink,o={root:["root",n,"color".concat((0,Z.Z)(t))],svg:["svg"],circle:["circle","circle".concat((0,Z.Z)(n)),a&&"circleDisableShrink"]};return(0,b.Z)(o,y.C,e)}(B),z={},F={},I={};if("determinate"===P){var L=2*Math.PI*((w-m)/2);z.strokeDasharray=L.toFixed(3),I["aria-valuenow"]=Math.round(k),z.strokeDashoffset="".concat(((100-k)/100*L).toFixed(3),"px"),F.transform="rotate(-90deg)"}return(0,S.jsx)(M,(0,v.Z)({className:(0,p.default)(D.root,t),style:(0,v.Z)({width:l,height:l},F,u),ownerState:B,ref:e,role:"progressbar"},I,R,{children:(0,S.jsx)(j,{className:D.svg,ownerState:B,viewBox:"".concat(22," ").concat(22," ").concat(w," ").concat(w),children:(0,S.jsx)(N,{className:D.circle,style:z,ownerState:B,cx:w,cy:w,r:(w-m)/2,fill:"none",strokeWidth:m})})}))}));e.Z=R},94808:function(r,e,n){n.d(e,{C:function(){return o}});var t=n(77430),a=n(32298);function o(r){return(0,a.Z)("MuiCircularProgress",r)}var i=(0,t.Z)("MuiCircularProgress",["root","determinate","indeterminate","colorPrimary","colorSecondary","svg","circle","circleDeterminate","circleIndeterminate","circleDisableShrink"]);e.Z=i},94149:function(r,e,n){var t,a,o,i,c,s,l,u,d,f,v,m,p=n(30168),b=n(63366),h=n(87462),Z=n(47313),g=n(83061),k=n(21921),y=n(30686),S=n(17551),x=n(91615),w=n(19860),C=n(17592),P=n(77342),M=n(66598),j=n(46417),N=["className","color","value","valueBuffer","variant"],R=(0,y.F4)(l||(l=t||(t=(0,p.Z)(["\n  0% {\n    left: -35%;\n    right: 100%;\n  }\n\n  60% {\n    left: 100%;\n    right: -90%;\n  }\n\n  100% {\n    left: 100%;\n    right: -90%;\n  }\n"])))),B=(0,y.F4)(u||(u=a||(a=(0,p.Z)(["\n  0% {\n    left: -200%;\n    right: 100%;\n  }\n\n  60% {\n    left: 107%;\n    right: -8%;\n  }\n\n  100% {\n    left: 107%;\n    right: -8%;\n  }\n"])))),D=(0,y.F4)(d||(d=o||(o=(0,p.Z)(["\n  0% {\n    opacity: 1;\n    background-position: 0 -23px;\n  }\n\n  60% {\n    opacity: 0;\n    background-position: 0 -23px;\n  }\n\n  100% {\n    opacity: 1;\n    background-position: -200px -23px;\n  }\n"])))),z=function(r,e){return"inherit"===e?"currentColor":r.vars?r.vars.palette.LinearProgress["".concat(e,"Bg")]:"light"===r.palette.mode?(0,S.$n)(r.palette[e].main,.62):(0,S._j)(r.palette[e].main,.5)},F=(0,C.ZP)("span",{name:"MuiLinearProgress",slot:"Root",overridesResolver:function(r,e){var n=r.ownerState;return[e.root,e["color".concat((0,x.Z)(n.color))],e[n.variant]]}})((function(r){var e=r.ownerState,n=r.theme;return(0,h.Z)({position:"relative",overflow:"hidden",display:"block",height:4,zIndex:0,"@media print":{colorAdjust:"exact"},backgroundColor:z(n,e.color)},"inherit"===e.color&&"buffer"!==e.variant&&{backgroundColor:"none","&::before":{content:'""',position:"absolute",left:0,top:0,right:0,bottom:0,backgroundColor:"currentColor",opacity:.3}},"buffer"===e.variant&&{backgroundColor:"transparent"},"query"===e.variant&&{transform:"rotate(180deg)"})})),I=(0,C.ZP)("span",{name:"MuiLinearProgress",slot:"Dashed",overridesResolver:function(r,e){var n=r.ownerState;return[e.dashed,e["dashedColor".concat((0,x.Z)(n.color))]]}})((function(r){var e=r.ownerState,n=r.theme,t=z(n,e.color);return(0,h.Z)({position:"absolute",marginTop:0,height:"100%",width:"100%"},"inherit"===e.color&&{opacity:.3},{backgroundImage:"radial-gradient(".concat(t," 0%, ").concat(t," 16%, transparent 42%)"),backgroundSize:"10px 10px",backgroundPosition:"0 -23px"})}),(0,y.iv)(f||(f=i||(i=(0,p.Z)(["\n    animation: "," 3s infinite linear;\n  "]))),D)),L=(0,C.ZP)("span",{name:"MuiLinearProgress",slot:"Bar1",overridesResolver:function(r,e){var n=r.ownerState;return[e.bar,e["barColor".concat((0,x.Z)(n.color))],("indeterminate"===n.variant||"query"===n.variant)&&e.bar1Indeterminate,"determinate"===n.variant&&e.bar1Determinate,"buffer"===n.variant&&e.bar1Buffer]}})((function(r){var e=r.ownerState,n=r.theme;return(0,h.Z)({width:"100%",position:"absolute",left:0,bottom:0,top:0,transition:"transform 0.2s linear",transformOrigin:"left",backgroundColor:"inherit"===e.color?"currentColor":(n.vars||n).palette[e.color].main},"determinate"===e.variant&&{transition:"transform .".concat(4,"s linear")},"buffer"===e.variant&&{zIndex:1,transition:"transform .".concat(4,"s linear")})}),(function(r){var e=r.ownerState;return("indeterminate"===e.variant||"query"===e.variant)&&(0,y.iv)(v||(v=c||(c=(0,p.Z)(["\n      width: auto;\n      animation: "," 2.1s cubic-bezier(0.65, 0.815, 0.735, 0.395) infinite;\n    "]))),R)})),q=(0,C.ZP)("span",{name:"MuiLinearProgress",slot:"Bar2",overridesResolver:function(r,e){var n=r.ownerState;return[e.bar,e["barColor".concat((0,x.Z)(n.color))],("indeterminate"===n.variant||"query"===n.variant)&&e.bar2Indeterminate,"buffer"===n.variant&&e.bar2Buffer]}})((function(r){var e=r.ownerState,n=r.theme;return(0,h.Z)({width:"100%",position:"absolute",left:0,bottom:0,top:0,transition:"transform 0.2s linear",transformOrigin:"left"},"buffer"!==e.variant&&{backgroundColor:"inherit"===e.color?"currentColor":(n.vars||n).palette[e.color].main},"inherit"===e.color&&{opacity:.3},"buffer"===e.variant&&{backgroundColor:z(n,e.color),transition:"transform .".concat(4,"s linear")})}),(function(r){var e=r.ownerState;return("indeterminate"===e.variant||"query"===e.variant)&&(0,y.iv)(m||(m=s||(s=(0,p.Z)(["\n      width: auto;\n      animation: "," 2.1s cubic-bezier(0.165, 0.84, 0.44, 1) 1.15s infinite;\n    "]))),B)})),T=Z.forwardRef((function(r,e){var n=(0,P.Z)({props:r,name:"MuiLinearProgress"}),t=n.className,a=n.color,o=void 0===a?"primary":a,i=n.value,c=n.valueBuffer,s=n.variant,l=void 0===s?"indeterminate":s,u=(0,b.Z)(n,N),d=(0,h.Z)({},n,{color:o,variant:l}),f=function(r){var e=r.classes,n=r.variant,t=r.color,a={root:["root","color".concat((0,x.Z)(t)),n],dashed:["dashed","dashedColor".concat((0,x.Z)(t))],bar1:["bar","barColor".concat((0,x.Z)(t)),("indeterminate"===n||"query"===n)&&"bar1Indeterminate","determinate"===n&&"bar1Determinate","buffer"===n&&"bar1Buffer"],bar2:["bar","buffer"!==n&&"barColor".concat((0,x.Z)(t)),"buffer"===n&&"color".concat((0,x.Z)(t)),("indeterminate"===n||"query"===n)&&"bar2Indeterminate","buffer"===n&&"bar2Buffer"]};return(0,k.Z)(a,M.E,e)}(d),v=(0,w.Z)(),m={},p={bar1:{},bar2:{}};if("determinate"===l||"buffer"===l)if(void 0!==i){m["aria-valuenow"]=Math.round(i),m["aria-valuemin"]=0,m["aria-valuemax"]=100;var Z=i-100;"rtl"===v.direction&&(Z=-Z),p.bar1.transform="translateX(".concat(Z,"%)")}else 0;if("buffer"===l)if(void 0!==c){var y=(c||0)-100;"rtl"===v.direction&&(y=-y),p.bar2.transform="translateX(".concat(y,"%)")}else 0;return(0,j.jsxs)(F,(0,h.Z)({className:(0,g.default)(f.root,t),ownerState:d,role:"progressbar"},m,{ref:e},u,{children:["buffer"===l?(0,j.jsx)(I,{className:f.dashed,ownerState:d}):null,(0,j.jsx)(L,{className:f.bar1,ownerState:d,style:p.bar1}),"determinate"===l?null:(0,j.jsx)(q,{className:f.bar2,ownerState:d,style:p.bar2})]}))}));e.Z=T},66598:function(r,e,n){n.d(e,{E:function(){return o}});var t=n(77430),a=n(32298);function o(r){return(0,a.Z)("MuiLinearProgress",r)}var i=(0,t.Z)("MuiLinearProgress",["root","colorPrimary","colorSecondary","determinate","indeterminate","buffer","query","dashed","dashedColorPrimary","dashedColorSecondary","bar","barColorPrimary","barColorSecondary","bar1Indeterminate","bar1Determinate","bar1Buffer","bar2Indeterminate","bar2Buffer"]);e.Z=i},42832:function(r,e,n){n.d(e,{Z:function(){return M}});var t=n(4942),a=n(63366),o=n(87462),i=n(47313),c=n(83061),s=n(13019),l=n(21921),u=n(32298),d=n(96694),f=n(14614),v=n(39028),m=n(9456),p=n(54929),b=n(86886),h=n(46417),Z=["component","direction","spacing","divider","children","className","useFlexGap"],g=(0,m.Z)(),k=(0,d.Z)("div",{name:"MuiStack",slot:"Root",overridesResolver:function(r,e){return e.root}});function y(r){return(0,f.Z)({props:r,name:"MuiStack",defaultTheme:g})}function S(r,e){var n=i.Children.toArray(r).filter(Boolean);return n.reduce((function(r,t,a){return r.push(t),a<n.length-1&&r.push(i.cloneElement(e,{key:"separator-".concat(a)})),r}),[])}var x=function(r){var e=r.ownerState,n=r.theme,a=(0,o.Z)({display:"flex",flexDirection:"column"},(0,p.k9)({theme:n},(0,p.P$)({values:e.direction,breakpoints:n.breakpoints.values}),(function(r){return{flexDirection:r}})));if(e.spacing){var i=(0,b.hB)(n),c=Object.keys(n.breakpoints.values).reduce((function(r,n){return("object"===typeof e.spacing&&null!=e.spacing[n]||"object"===typeof e.direction&&null!=e.direction[n])&&(r[n]=!0),r}),{}),l=(0,p.P$)({values:e.direction,base:c}),u=(0,p.P$)({values:e.spacing,base:c});"object"===typeof l&&Object.keys(l).forEach((function(r,e,n){if(!l[r]){var t=e>0?l[n[e-1]]:"column";l[r]=t}}));a=(0,s.Z)(a,(0,p.k9)({theme:n},u,(function(r,n){return e.useFlexGap?{gap:(0,b.NA)(i,r)}:{"& > :not(style) + :not(style)":(0,t.Z)({margin:0},"margin".concat((a=n?l[n]:e.direction,{row:"Left","row-reverse":"Right",column:"Top","column-reverse":"Bottom"}[a])),(0,b.NA)(i,r))};var a})))}return a=(0,p.dt)(n.breakpoints,a)};var w=n(17592),C=n(77342),P=function(){var r=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},e=r.createStyledComponent,n=void 0===e?k:e,t=r.useThemeProps,s=void 0===t?y:t,d=r.componentName,f=void 0===d?"MuiStack":d,m=n(x),p=i.forwardRef((function(r,e){var n=s(r),t=(0,v.Z)(n),i=t.component,d=void 0===i?"div":i,p=t.direction,b=void 0===p?"column":p,g=t.spacing,k=void 0===g?0:g,y=t.divider,x=t.children,w=t.className,C=t.useFlexGap,P=void 0!==C&&C,M=(0,a.Z)(t,Z),j={direction:b,spacing:k,useFlexGap:P},N=(0,l.Z)({root:["root"]},(function(r){return(0,u.Z)(f,r)}),{});return(0,h.jsx)(m,(0,o.Z)({as:d,ownerState:j,ref:e,className:(0,c.default)(N.root,w)},M,{children:y?S(x,y):x}))}));return p}({createStyledComponent:(0,w.ZP)("div",{name:"MuiStack",slot:"Root",overridesResolver:function(r,e){return e.root}}),useThemeProps:function(r){return(0,C.Z)({props:r,name:"MuiStack"})}}),M=P},96694:function(r,e,n){var t=(0,n(36541).ZP)();e.Z=t}}]);