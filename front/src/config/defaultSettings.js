/**
 * 项目默认配置项
 * primaryColor - 默认主题色, 如果修改颜色不生效，请清理 localStorage
 * navTheme - sidebar theme ['dark', 'light'] 两种主题
 * colorWeak - 色盲模式
 * layout - 整体布局方式 ['sidemenu', 'topmenu'] 两种布局
 * fixedHeader - 固定 Header : boolean
 * fixSiderbar - 固定左侧菜单栏 ： boolean
 * contentWidth - 内容区布局： 流式 |  固定
 *
 * storageOptions: {} - Vue-ls 插件配置项 (localStorage/sessionStorage)
 *
 */

export default {
  navTheme: 'dark', // theme for nav menu
  primaryColor: '#F5222D', // primary color of ant design
  layout: 'sidemenu', // nav menu position: `sidemenu` or `topmenu`
  contentWidth: 'Fluid', // layout of content: `Fluid` or `Fixed`, only works when layout is topmenu
  fixedHeader: false, // sticky header
  fixSiderbar: false, // sticky siderbar
  colorWeak: false,
  menu: {
    locale: true
  },
  title: 'user.home.title',
  pwa: false,
  iconfontUrl: '',
  production: process.env.NODE_ENV === 'production' && process.env.VUE_APP_PREVIEW !== 'true'
}

export const authButton = [{
  'action': 'add',
  'describe': '新增',
  'defaultCheck': false
}, {
  'action': 'import',
  'describe': '导入',
  'defaultCheck': false
}, {
  'action': 'get',
  'describe': '详情',
  'defaultCheck': false
}, {
  'action': 'update',
  'describe': '修改',
  'defaultCheck': false
}, {
  'action': 'delete',
  'describe': '删除',
  'defaultCheck': false
}, {
  'action': 'export',
  'describe': '导出',
  'defaultCheck': false
}]

export const authMaps = authButton.reduce((pre,next) => {
  pre [next.action] = next.describe
  return pre
},{})
