"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[376],{220:(e,s,n)=>{n.r(s),n.d(s,{assets:()=>o,contentTitle:()=>a,default:()=>u,frontMatter:()=>t,metadata:()=>l,toc:()=>d});var i=n(2488),r=n(6428);const t={sidebar_position:2},a="api.sys",l={id:"api-reference/sys",title:"api.sys",description:"The api.sys module provides functions for working with the system.",source:"@site/docs/api-reference/sys.md",sourceDirName:"api-reference",slug:"/api-reference/sys",permalink:"/spito/docs/api-reference/sys",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/api-reference/sys.md",tags:[],version:"current",sidebarPosition:2,frontMatter:{sidebar_position:2},sidebar:"tutorialSidebar",previous:{title:"api.pkg",permalink:"/spito/docs/api-reference/pkg"},next:{title:"api.fs",permalink:"/spito/docs/api-reference/fs"}},o={},d=[{value:"api.sys.getDistro",id:"apisysgetdistro",level:2},{value:"Arguments:",id:"arguments",level:3},{value:"Returns:",id:"returns",level:3},{value:"Example usage:",id:"example-usage",level:3},{value:"api.sys.getDaeomon",id:"apisysgetdaeomon",level:2},{value:"Arguments:",id:"arguments-1",level:3},{value:"Returns:",id:"returns-1",level:3},{value:"Example usage:",id:"example-usage-1",level:3},{value:"api.sys.getInitSystem",id:"apisysgetinitsystem",level:2},{value:"Arguments:",id:"arguments-2",level:3},{value:"Returns:",id:"returns-2",level:3},{value:"Example usage:",id:"example-usage-2",level:3}];function c(e){const s={code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",ul:"ul",...(0,r.M)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(s.h1,{id:"apisys",children:"api.sys"}),"\n",(0,i.jsxs)(s.p,{children:["The ",(0,i.jsx)(s.code,{children:"api.sys"})," module provides functions for working with the system."]}),"\n",(0,i.jsx)(s.h2,{id:"apisysgetdistro",children:"api.sys.getDistro"}),"\n",(0,i.jsx)(s.h3,{id:"arguments",children:"Arguments:"}),"\n",(0,i.jsxs)(s.ul,{children:["\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"name"})," (string): The name of the package to get."]}),"\n"]}),"\n",(0,i.jsx)(s.h3,{id:"returns",children:"Returns:"}),"\n",(0,i.jsxs)(s.ul,{children:["\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"distro"})," (Distro): The distro info."]}),"\n"]}),"\n",(0,i.jsx)(s.h3,{id:"example-usage",children:"Example usage:"}),"\n",(0,i.jsx)(s.pre,{children:(0,i.jsx)(s.code,{className:"language-lua",children:"local distro = api.sys.getDistro()\n"})}),"\n",(0,i.jsx)(s.h2,{id:"apisysgetdaeomon",children:"api.sys.getDaeomon"}),"\n",(0,i.jsx)(s.h3,{id:"arguments-1",children:"Arguments:"}),"\n",(0,i.jsxs)(s.ul,{children:["\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"name"})," (string): The name of the package to get."]}),"\n"]}),"\n",(0,i.jsx)(s.h3,{id:"returns-1",children:"Returns:"}),"\n",(0,i.jsxs)(s.ul,{children:["\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"daemon"})," (Daemon): The daemon info."]}),"\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"error"})," (string): The error message if the daemon does not exist."]}),"\n"]}),"\n",(0,i.jsx)(s.h3,{id:"example-usage-1",children:"Example usage:"}),"\n",(0,i.jsx)(s.pre,{children:(0,i.jsx)(s.code,{className:"language-lua",children:'function networkManagerExists()\n  local daemon, err = api.sys.getDaemon("dbus")\n  if err ~= nil then\n    api.info.error("Error occured during obtaining daemon info!")\n    return false\n  end\n  return true\nend\n'})}),"\n",(0,i.jsx)(s.h2,{id:"apisysgetinitsystem",children:"api.sys.getInitSystem"}),"\n",(0,i.jsx)(s.h3,{id:"arguments-2",children:"Arguments:"}),"\n",(0,i.jsxs)(s.ul,{children:["\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"name"})," (string): The name of the package to get."]}),"\n"]}),"\n",(0,i.jsx)(s.h3,{id:"returns-2",children:"Returns:"}),"\n",(0,i.jsxs)(s.ul,{children:["\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"initSystem"})," (InitSystem): The init system info."]}),"\n",(0,i.jsxs)(s.li,{children:[(0,i.jsx)(s.code,{children:"error"})," (string): The error message if the init system does not exist."]}),"\n"]}),"\n",(0,i.jsx)(s.h3,{id:"example-usage-2",children:"Example usage:"}),"\n",(0,i.jsx)(s.pre,{children:(0,i.jsx)(s.code,{className:"language-lua",children:'function initSystemExists()\n  local initSystem, err = api.sys.getInitSystem()\n  if err ~= nil then\n    api.info.error("Error occured during obtaining init system info!")\n    return false\n  end\n  return true\nend\n'})})]})}function u(e={}){const{wrapper:s}={...(0,r.M)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(c,{...e})}):c(e)}},6428:(e,s,n)=>{n.d(s,{I:()=>l,M:()=>a});var i=n(6651);const r={},t=i.createContext(r);function a(e){const s=i.useContext(t);return i.useMemo((function(){return"function"==typeof e?e(s):{...s,...e}}),[s,e])}function l(e){let s;return s=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:a(e.components),i.createElement(t.Provider,{value:s},e.children)}}}]);