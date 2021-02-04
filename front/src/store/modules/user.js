import Vue from 'vue'
import { login, getInfo, logout } from '@/api/login'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import { welcome } from '@/utils/util'
import { authButton } from '@/config/defaultSettings'

const user = {
  state: {
    token: '',
    name: '',
    welcome: '',
    avatar: '',
    roles: [],
    info: {},
    userInfo: {},
    realUser: {}
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, { name, welcome }) => {
      state.name = name
      state.welcome = welcome
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_INFO: (state, info) => {
      state.info = info
    },
    SET_USER_INFO: (state, info) => {
      state.userInfo = info
    },
    SET_REAL_USER(state,info) {
      state.realUser = info
    }
  },
  getters: {
    realUser(state) { return state.realUser }
  },
  actions: {
    // 登录
    Login ({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        login(userInfo).then(result => {
          Vue.ls.set(ACCESS_TOKEN, result.token + 'whx', 7 * 24 * 60 * 60 * 1000)
          Vue.ls.set('userInfo', result, 7 * 24 * 60 * 60 * 1000)
          commit('SET_TOKEN', result.token)
          commit('SET_USER_INFO', result)
          commit('SET_REAL_USER',result)
          resolve(result)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取用户信息
    GetInfo ({ commit,state }) {
      return new Promise((resolve, reject) => {
        getInfo({ roleId: Vue.ls.get('userInfo').roleId }).then(result => {
          if (result.role && result.role.permissions.length > 0) {
            const role = result.role
            role.permissions = result.role.permissions
            role.permissions.forEach(pre => {
              pre.actions = pre.actionsObject.map(item => item.operator)
              pre.actionEntitySet = authButton.filter(item => {
                return pre.actions.includes(item.action)
              })
            })
            role.permissions.map(per => {
              if (per.actionEntitySet != null && per.actionEntitySet.length > 0) {
                const action = per.actionEntitySet.map(action => { return action.action })
                per.actionList = action
              }
            })
            role.permissionList = role.permissions.map(permission => { return permission.permissionId })
            commit('SET_ROLES', result.role)
            commit('SET_INFO', result)
          } else {
            reject(new Error('getInfo: roles must be a non-null array !'))
          }

          commit('SET_NAME', { name: result.name, welcome: welcome() })
          commit('SET_AVATAR', result.avatar)

          resolve(result)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 登出
    Logout ({ commit, state }) {
      return new Promise((resolve) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          Vue.ls.remove(ACCESS_TOKEN)
          resolve()
        }).catch(() => {
          resolve()
        }).finally(() => {
        })
      })
    }

  }
}

export default user
