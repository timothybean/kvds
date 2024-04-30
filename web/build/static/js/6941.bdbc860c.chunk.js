"use strict";(self.webpackChunkfuse_react_app=self.webpackChunkfuse_react_app||[]).push([[6941],{66941:function(e,s,l){l.r(s);var n=l(1413),c=l(29439),t=l(53996),r=l(24193),a=l(70501),i=l(17592),o=l(61113),d=l(56573),x=l(47313),m=l(1550),u=l(75627),h=l(56993),f=l(24631),j=l(15103),p=l(40454),v=l(51405),Z=l(98655),g=l(46417),N=(0,i.ZP)(t.Z)((function(e){var s=e.theme;return{"& .FusePageSimple-header":{backgroundColor:s.palette.background.paper,borderBottomWidth:1,borderStyle:"solid",borderColor:s.palette.divider}}}));s.default=function(e){var s=null===e||void 0===e?void 0:e.data,l=s.pageTitle,t=s.referenceUrl,i=s.apiUrl,b=s.iconName,y=(0,x.useState)(null),k=(0,c.Z)(y,2),w=k[0],S=k[1],z=(0,x.useState)(""),C=(0,c.Z)(z,2),I=C[0],F=C[1],T=(0,x.useState)(null),_=(0,c.Z)(T,2),U=_[0],P=_[1],W=(0,u.cI)({mode:"onChange",defaultValues:{searchText:"",size:24}}),E=W.watch,Q=W.control,B=E(),L=E("searchText");return(0,x.useEffect)((function(){d.Z.get(i).then((function(e){S(e.data),F(e.data[0])}))}),[i]),(0,x.useEffect)((function(){P(L.length>0?w.filter((function(e){return!!e.includes(L)})):w)}),[w,L]),(0,g.jsx)(N,{header:(0,g.jsxs)("div",{className:"flex flex-col sm:flex-row flex-0 sm:items-center sm:justify-between p-24 sm:py-32 sm:px-40",children:[(0,g.jsxs)("div",{className:"flex-1 min-w-0",children:[(0,g.jsx)("div",{className:"flex flex-wrap items-center font-medium",children:(0,g.jsx)("div",{children:(0,g.jsx)(o.Z,{className:"whitespace-nowrap",color:"secondary",children:"User Interface"})})}),(0,g.jsx)("div",{className:"mt-8",children:(0,g.jsx)(o.Z,{className:"text-3xl md:text-4xl font-extrabold tracking-tight leading-7 sm:leading-10 truncate",children:l})})]}),(0,g.jsx)("div",{children:t&&(0,g.jsx)(r.Z,{className:"mt-12 sm:mt-0",variant:"contained",color:"secondary",component:"a",href:t,target:"_blank",role:"button",startIcon:(0,g.jsx)(h.Z,{children:"heroicons-solid:external-link"}),children:"Official docs"})})]}),content:(0,g.jsxs)("div",{className:"flex-auto p-24 sm:p-40",children:[(0,g.jsx)(o.Z,{className:"text-20 font-700 mb-16",children:"Usage"}),(0,g.jsx)(Z.Z,{component:"pre",className:"language-jsx my-24",children:'\n              <FuseSvgIcon className="text-48" size={'.concat(B.size,'} color="action">').concat(b,":").concat(I,"</FuseSvgIcon>\n            ")}),(0,g.jsx)(o.Z,{className:"text-20 font-700 mt-32 mb-16",children:"Icons"}),(0,g.jsxs)("div",{className:"flex flex-col md:flex-row justify-center md:items-end my-24 xs:flex-col md:space-x-24",children:[(0,g.jsx)("div",{className:"flex flex-1",children:(0,g.jsx)(u.Qr,{name:"searchText",control:Q,render:function(e){var s=e.field;return(0,g.jsx)(f.Z,(0,n.Z)((0,n.Z)({},s),{},{id:"searchText",label:"Search an icon",placeholder:"Search..",className:"flex-auto",InputLabelProps:{shrink:!0},variant:"outlined",fullWidth:!0}))}})}),(0,g.jsx)(u.Qr,{name:"size",control:Q,render:function(e){var s=e.field;return(0,g.jsxs)(m.Z,{sx:{mt:2,minWidth:120},children:[(0,g.jsx)(j.Z,{htmlFor:"max-width",children:"Size"}),(0,g.jsxs)(p.Z,(0,n.Z)((0,n.Z)({autoFocus:!0},s),{},{label:"Size",children:[(0,g.jsx)(v.Z,{value:16,children:"16"}),(0,g.jsx)(v.Z,{value:20,children:"20"}),(0,g.jsx)(v.Z,{value:24,children:"24"}),(0,g.jsx)(v.Z,{value:32,children:"32"}),(0,g.jsx)(v.Z,{value:40,children:"40"}),(0,g.jsx)(v.Z,{value:48,children:"48"}),(0,g.jsx)(v.Z,{value:64,children:"64"})]}))]})}})]}),(0,g.jsx)("div",{className:"flex grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-12 sm:gap-32 py-24",children:(0,x.useMemo)((function(){return U&&(U.length>0?U.map((function(e){return(0,g.jsxs)(a.Z,{className:"flex flex-col items-center justify-center p-16 border-2 min-h-120 rounded cursor-pointer",elevation:0,onClick:function(){return F(e)},sx:{borderColor:e===I&&"secondary.main"},children:[(0,g.jsx)("div",{className:"flex items-center justify-center mb-12",children:(0,g.jsx)(h.Z,{className:"text-48",size:B.size,color:"action",children:"".concat(b,":").concat(e)})}),(0,g.jsx)(o.Z,{className:"text-sm text-center break-all",color:"text.secondary",children:"".concat(b,":").concat(e)})]},e)})):(0,g.jsx)("div",{className:"col-span-6 flex flex-auto items-center justify-center w-full h-full p-32 md:p-128",children:(0,g.jsx)(o.Z,{color:"text.secondary",variant:"h5",children:"No results!"})}))}),[I,U,B.size,b])})]})})}}}]);