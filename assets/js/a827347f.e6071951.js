"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[448],{8084:(e,n,s)=>{s.r(n),s.d(n,{assets:()=>a,contentTitle:()=>o,default:()=>u,frontMatter:()=>t,metadata:()=>c,toc:()=>d});var r=s(2488),i=s(6428);const t={sidebar_position:5},o="api.sh",c={id:"api-reference/sh",title:"api.sh",description:"The api.sh module provides functions for executing shell commands.",source:"@site/docs/api-reference/sh.md",sourceDirName:"api-reference",slug:"/api-reference/sh",permalink:"/spito/docs/api-reference/sh",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/api-reference/sh.md",tags:[],version:"current",sidebarPosition:5,frontMatter:{sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"api.info",permalink:"/spito/docs/api-reference/info"}},a={},d=[{value:"api.sh.command",id:"apishcommand",level:2},{value:"Arguments:",id:"arguments",level:3},{value:"Returns:",id:"returns",level:3},{value:"Example usage:",id:"example-usage",level:3}];function l(e){const n={admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",ul:"ul",...(0,i.M)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.h1,{id:"apish",children:"api.sh"}),"\n",(0,r.jsxs)(n.p,{children:["The ",(0,r.jsx)(n.code,{children:"api.sh"})," module provides functions for executing shell commands."]}),"\n",(0,r.jsx)(n.admonition,{type:"warning",children:(0,r.jsx)(n.p,{children:"This module works only if the rule is unsafe."})}),"\n",(0,r.jsx)(n.h2,{id:"apishcommand",children:"api.sh.command"}),"\n",(0,r.jsx)(n.h3,{id:"arguments",children:"Arguments:"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"command"})," (string): The command to execute."]}),"\n"]}),"\n",(0,r.jsx)(n.h3,{id:"returns",children:"Returns:"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"output"})," (string): The output of the command."]}),"\n",(0,r.jsxs)(n.li,{children:[(0,r.jsx)(n.code,{children:"error"})," (string): The error message if the command fails."]}),"\n"]}),"\n",(0,r.jsx)(n.h3,{id:"example-usage",children:"Example usage:"}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-lua",children:'#![unsafe]\n\nfunction ls()\n  local output, err = api.sh.command("ls -l")\n  if err ~= nil then\n    api.info.Error("Error occured while executing the command: " .. err)\n    return false\n  end\n  return true\nend\n'})})]})}function u(e={}){const{wrapper:n}={...(0,i.M)(),...e.components};return n?(0,r.jsx)(n,{...e,children:(0,r.jsx)(l,{...e})}):l(e)}},6428:(e,n,s)=>{s.d(n,{I:()=>c,M:()=>o});var r=s(6651);const i={},t=r.createContext(i);function o(e){const n=r.useContext(t);return r.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function c(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:o(e.components),r.createElement(t.Provider,{value:n},e.children)}}}]);