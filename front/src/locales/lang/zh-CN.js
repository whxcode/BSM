import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'
import global from './zh-CN/global'

import menu from './zh-CN/menu'
import setting from './zh-CN/setting'
import user from './zh-CN/user'
import dashboard from './zh-CN/dashboard'
import form from './zh-CN/form'
import result from './zh-CN/result'
import account from './zh-CN/account'
import acustome from '@/locales/lang/zh-CN/acustome'

const components = {
  antLocale: antd,
  momentName: 'zh-cn',
  momentLocale: momentCN
}

export default {
  message: '-',
  ...acustome,
  ...components,
  ...global,
  ...menu,
  ...setting,
  ...user,
  ...dashboard,
  ...form,
  ...result,
  ...account
}
